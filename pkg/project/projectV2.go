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

func (p *ProjectV2) ListFieldsForProjectByCursor(cfg *config.Config, cursor string) (fieldsReply *FieldsReply, err error) {
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

func listFieldValueForIssue[T FieldValue](cfg *config.Config, issue *issue.Issue,
	field *Field, includeArchived bool, cursor string) {
	switch field.DataType {
	case "TITLE":
	}
}

func ListFieldValueForIssueByCursor[T FieldValue](cfg *config.Config, issue *issue.Issue,
	field *Field, includeArchived bool, fieldValue T,
	cursor string, subCursor string) (FieldValues[T], error) {

	data := &ProjectReply[T]{}
	items := make(FieldValues[T], 0, 10)

	perPage := cfg.GetPerPage()

	for {
		query, err := fieldValue.GenQuery(cfg, issue, field, includeArchived, perPage, cursor, subCursor)
		if err != nil {
			return nil, err
		}
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

		nodes := data.Data.Repository.Issue.ProjectItems.Nodes
		for j := 0; j < len(nodes); j++ {
			if !nodes[j].FieldValueByName.IsNil() {
				items = append(items, nodes[j].FieldValueByName)
			}
		}

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

func ListFieldValueForIssueByCursorTest(cfg *config.Config, issue *issue.Issue,
	field *Field, includeArchived bool, fieldValue FieldValue,
	cursor string, subCursor string) ([]FieldValue, error) {

	data := &ProjectReply{}
	items := make([]FieldValue, 0, 10)

	perPage := cfg.GetPerPage()

	for {
		query, err := fieldValue.GenQuery(cfg, issue, field, includeArchived, perPage, cursor, subCursor)
		if err != nil {
			return nil, err
		}
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

		nodes := data.Data.Repository.Issue.ProjectItems.Nodes
		for j := 0; j < len(nodes); j++ {
			if !nodes[j].FieldValueByName.IsNil() {
				items = append(items, nodes[j].FieldValueByName)
			}
		}

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
