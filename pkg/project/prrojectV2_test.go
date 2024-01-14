package project

import (
	"fmt"
	"os"
	"testing"

	"github.com/guguducken/octopus/pkg/config"
	"github.com/guguducken/octopus/pkg/organization"
)

func TestGetProjectV2ForOrganization(t *testing.T) {
	cfg := config.New(os.Getenv("GITHUB_TOKEN"))
	org := &organization.Organization{
		Login: "gugus-test",
	}
	p, err := GetProjectV2ForOrganization(cfg, org, 1)
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}
	fmt.Printf("p.Title: %v\n", p.Organization.Login)
	fmt.Printf("p.Organization.ID: %v\n", p.Organization.ID)
}

func TestListFieldsForProjectByCursor(t *testing.T) {
	cfg := config.New(os.Getenv("GITHUB_TOKEN"))
	org := &organization.Organization{
		Login: "gugus-test",
	}
	p, err := GetProjectV2ForOrganization(cfg, org, 1)
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}
	fieldsReply, err := p.ListFieldsForProjectByCursor(cfg, "")
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}
	fmt.Printf("fieldsReply.PageInfo.EndCursor: %v\n", fieldsReply.PageInfo.EndCursor)
	fmt.Printf("fieldsReply.PageInfo.HasNextPage: %v\n", fieldsReply.PageInfo.HasNextPage)
	for _, node := range fieldsReply.Nodes {
		fmt.Printf("node.Name: %v\n", node.Name)
		fmt.Printf("node.DataType: %v\n", node.DataType)
		fmt.Println("=========================================")
	}
}
