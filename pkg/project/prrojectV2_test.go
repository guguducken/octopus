package project

import (
	"fmt"
	"github.com/guguducken/octopus/pkg/issue"
	"github.com/guguducken/octopus/pkg/repository"
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

func TestListProjectV2ForIssueByCursor(t *testing.T) {
	cfg := config.New(os.Getenv("GITHUB_TOKEN"))
	repo, err := repository.GetRepository(cfg, "gugus-test", "test")
	if err != nil {
		panic(err)
	}
	is, err := issue.GetIssueForRepo(cfg, repo, 3)
	if err != nil {
		panic(err)
	}
	projects, err := ListProjectV2ForIssueByCursor(cfg, is, "")
	if err != nil {
		panic(err)
	}
	fmt.Printf("projects.PageInfo.HasNextPage: %v\n", projects.PageInfo.HasNextPage)
	fmt.Printf("projects.Nodes: %v\n", projects.Nodes)
}

func TestListProjectV2ForIssue(t *testing.T) {
	cfg := config.New(os.Getenv("GITHUB_TOKEN")).SetPerPage(1)
	repo, err := repository.GetRepository(cfg, "matrixorigin", "matrixone")
	if err != nil {
		panic(err)
	}
	is, err := issue.GetIssueForRepo(cfg, repo, 9675)
	if err != nil {
		panic(err)
	}
	projects, err := ListProjectV2ForIssue(cfg, is)
	if err != nil {
		panic(err)
	}
	fmt.Printf("len(projects): %v\n", len(projects))
	for i := 0; i < len(projects); i++ {
		fmt.Printf("projects[i].Number: %v\n", projects[i].Number)
		fmt.Printf("projects[i].Title: %v\n", projects[i].Title)
		fmt.Printf("projects[i].Closed: %v\n", projects[i].Closed)
		fmt.Println("============================================")
	}
}

func TestListFieldValueForIssueByCursor(t *testing.T) {
	cfg := config.New(os.Getenv("GITHUB_TOKEN"))
	org := &organization.Organization{
		Login: "gugus-test",
	}
	p, err := GetProjectV2ForOrganization(cfg, org, 1)
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}
	// get fields for a project
	fieldsReply, err := p.ListFieldsForProjectByCursor(cfg, "")
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return
	}

	repo, err := repository.GetRepository(cfg, "gugus-test", "test")
	if err != nil {
		panic(err)
	}
	is, err := issue.GetIssueForRepo(cfg, repo, 3)
	if err != nil {
		panic(err)
	}
	testValue := SingleSelectFieldValue{}
	var field *Field
	for i := 0; i < len(fieldsReply.Nodes); i++ {
		if fieldsReply.Nodes[i].DataType == "SINGLE_SELECT" {
			field = &fieldsReply.Nodes[i]
		}
	}
	fieldValues, err := ListFieldValueForIssueByCursor(cfg, is, field, false, testValue, "", "")
	if err != nil {
		return
	}
	for _, value := range fieldValues {
		fmt.Printf("value.Text: %v\n", value.Name)
		fmt.Printf("value.FilterProject(p): %v\n", value.FilterProject(p))
	}
}
