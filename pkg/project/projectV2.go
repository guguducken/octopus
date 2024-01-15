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

func ListFieldValueForIssueByCursor[T FieldValue](cfg *config.Config, issue *issue.Issue,
	filed *Field, includeArchived bool, fieldValue T,
	cursor string, subCursor string) (items FieldValues[T], err error) {

	data := &ProjectReply[T]{}
	perPage := cfg.GetPerPage()

	query, err := fieldValue.GenQuery(cfg, issue, filed, includeArchived, perPage, cursor, subCursor)
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

	items = make(FieldValues[T], 0, 10)
	for j := 0; j < len(data.Data.Repository.Issue.ProjectItems.Nodes); j++ {
		items = append(items, data.Data.Repository.Issue.ProjectItems.Nodes[j].FieldValueByName)
	}
	return items, err
}
