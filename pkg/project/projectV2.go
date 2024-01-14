package project

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/guguducken/octopus/pkg/config"
	"github.com/guguducken/octopus/pkg/graphql"
	"github.com/guguducken/octopus/pkg/organization"
)

func GetProjectV2ForOrganization(cfg *config.Config, org *organization.Organization, id int) (projectv2 *ProjectV2, err error) {
	data := &ProjectReply{}

	query := QueryToJson(fmt.Sprintf(GetProjectForOrgByID, org.Login, id))
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
	data := &ProjectReply{}
	perPage := cfg.GetPerPage()

	query := QueryToJson(fmt.Sprintf(GetProjectV2FieldsByPage, p.Organization.Login, p.Number, perPage, cursor))
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
