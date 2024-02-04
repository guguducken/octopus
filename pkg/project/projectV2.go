package project

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/guguducken/octopus/pkg/config"
	"github.com/guguducken/octopus/pkg/graphql"
	"github.com/guguducken/octopus/pkg/issue"
	"github.com/guguducken/octopus/pkg/organization"
)

const (
	DefaultStartCursor = ""
)

func GetProjectV2ForOrganization(cfg *config.Config, org *organization.Organization, id int) (projectv2 *ProjectV2, err error) {
	data := &ProjectReply[CommonFieldValue]{}

	query := QueryToJson(fmt.Sprintf(QueryProjectForOrgByID, org.Login, id))
	reply, err := graphql.NewGitHubGraphQL().Exec(cfg, query)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(reply, data); err != nil {
		return nil, err
	}
	if len(data.Errors) != 0 {
		return nil, errors.Join(ErrGraphQLResult, errors.New(data.Errors.ToJson()))
	}
	data.Data.Organization.ProjectV2.Organization.Login = data.Data.Organization.Login
	data.Data.Organization.ProjectV2.Organization.ID = data.Data.Organization.ID
	return data.Data.Organization.ProjectV2, err
}

func (p *ProjectV2) ListFields(cfg *config.Config) (fields []Field, err error) {
	cursor := ""
	fields = make([]Field, 0, 10)
	for {
		var fieldsReply *FieldsReply
		fieldsReply, err = p.ListFieldsByCursor(cfg, cursor)
		if err != nil {
			break
		}
		fields = append(fields, fieldsReply.Nodes...)
		if fieldsReply.PageInfo.HasNextPage {
			cursor = fieldsReply.PageInfo.EndCursor
			continue
		}
		break
	}
	return fields, err
}

func (p *ProjectV2) ListFieldsByCursor(cfg *config.Config, cursor string) (fieldsReply *FieldsReply, err error) {
	data := &ProjectReply[CommonFieldValue]{}
	perPage := cfg.GetPerPage()

	query := QueryToJson(fmt.Sprintf(QueryProjectV2FieldsByPage, p.Organization.Login, p.Number, perPage, cursor))
	reply, err := graphql.NewGitHubGraphQL().Exec(cfg, query)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(reply, data); err != nil {
		return nil, err
	}
	if len(data.Errors) != 0 {
		return nil, errors.Join(ErrGraphQLResult, errors.New(data.Errors.ToJson()))
	}
	return data.Data.Organization.ProjectV2.Fields, err
}

func ListProjectV2ForIssue(cfg *config.Config, issue *issue.Issue) (projectsV2 []ProjectV2, err error) {
	projectsV2 = make([]ProjectV2, 0, 10)
	cursor := ""
	for {
		var projectsReply *IssueRelatedProjectReply
		projectsReply, err = ListProjectV2ForIssueByCursor(cfg, issue, cursor)
		if err != nil {
			break
		}
		projectsV2 = append(projectsV2, projectsReply.Nodes...)
		if projectsReply.PageInfo.HasNextPage {
			cursor = projectsReply.PageInfo.EndCursor
		} else {
			break
		}
	}
	return projectsV2, err
}

func ListProjectV2ForIssueByCursor(cfg *config.Config, issue *issue.Issue, cursor string) (projectsV2Reply *IssueRelatedProjectReply, err error) {
	data := &ProjectReply[CommonFieldValue]{}
	perPage := cfg.GetPerPage()

	query := QueryToJson(fmt.Sprintf(QueryIssueRelatedProjectsV2,
		issue.Repository.Owner.Login,
		issue.Repository.Name,
		issue.Number,
		perPage,
		cursor,
	))
	reply, err := graphql.NewGitHubGraphQL().Exec(cfg, query)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(reply, data); err != nil {
		return nil, err
	}
	if len(data.Errors) != 0 {
		return nil, errors.Join(ErrGraphQLResult, errors.New(data.Errors.ToJson()))
	}
	return data.Data.Repository.Issue.ProjectsV2, err
}

func listFieldValueForIssue[T FieldValuer](cfg *config.Config, issue *issue.Issue,
	field *Field, includeArchived bool, cursor string) ([]T, error) {
	allFieldValues := GenFieldValuers[T]()
	var err error
	var filterType T

	for {
		var fieldValues *FieldValues[T]
		fieldValues, err = ListFieldValueForIssueByCursor(cfg, issue, field, includeArchived, filterType, cursor, "")
		if err != nil {
			break
		}
		allFieldValues.AddSlice(fieldValues.GetFieldValues())
		pageInfo := fieldValues.GetPageInfo()
		if pageInfo.HasNextPage {
			cursor = pageInfo.EndCursor
			continue
		}
		break
	}
	return allFieldValues.GetFieldValues(), err
}

func ListFieldValueForIssueByCursor[T FieldValuer](cfg *config.Config, issue *issue.Issue,
	field *Field, includeArchived bool, fieldValue T,
	cursor string, subCursor string) (*FieldValues[T], error) {

	data := &ProjectReply[T]{}
	items := GenFieldValuers[T]()

	perPage := cfg.GetPerPage()

	for {
		query, err := fieldValue.GenQuery(cfg, issue, field, includeArchived, perPage, cursor, subCursor)
		if err != nil {
			return items, err
		}
		reply, err := graphql.NewGitHubGraphQL().Exec(cfg, query)
		if err != nil {
			return items, err
		}
		if err = json.Unmarshal(reply, data); err != nil {
			return items, err
		}
		if len(data.Errors) != 0 {
			return items, errors.Join(ErrGraphQLResult, errors.New(data.Errors.ToJson()))
		}

		nodes := data.Data.Repository.Issue.ProjectItems.Nodes
		for j := 0; j < len(nodes); j++ {
			if !nodes[j].FieldValueByName.IsNil() {
				items.Add(nodes[j].FieldValueByName)
			}
		}

		pageInfo := data.Data.Repository.Issue.ProjectItems.PageInfo
		items.SetPageInfo(pageInfo)

		if len(nodes) != 0 {
			subPageInfo := nodes[0].FieldValueByName.GetSubPageInfo()
			if subPageInfo != nil && subPageInfo.HasNextPage {
				subCursor = subPageInfo.EndCursor
				continue
			}
		}
		break
	}
	return items, nil
}

func ListTextFieldValueForIssue(cfg *config.Config, issue *issue.Issue,
	field *Field, includeArchived bool) ([]TextFieldValue, error) {
	return listFieldValueForIssue[TextFieldValue](cfg, issue, field, includeArchived, DefaultStartCursor)
}

func ListSingleSelectFieldValueForIssue(cfg *config.Config, issue *issue.Issue,
	field *Field, includeArchived bool) ([]SingleSelectFieldValue, error) {
	return listFieldValueForIssue[SingleSelectFieldValue](cfg, issue, field, includeArchived, DefaultStartCursor)
}

func ListNumberFieldValueForIssue(cfg *config.Config, issue *issue.Issue,
	field *Field, includeArchived bool) ([]NumberFieldValue, error) {
	return listFieldValueForIssue[NumberFieldValue](cfg, issue, field, includeArchived, DefaultStartCursor)
}

func ListLabelFieldValueForIssue(cfg *config.Config, issue *issue.Issue,
	field *Field, includeArchived bool) ([]LabelFieldValue, error) {
	return listFieldValueForIssue[LabelFieldValue](cfg, issue, field, includeArchived, DefaultStartCursor)
}
