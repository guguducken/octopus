package issue

import (
	"fmt"
	"os"
	"testing"

	"github.com/guguducken/octopus/pkg/config"
	"github.com/guguducken/octopus/pkg/repository"
)

func TestGetIssueForRepo(t *testing.T) {
	cfg := config.New(os.Getenv("GITHUB_TOKEN"))
	repo, err := repository.GetRepository(cfg, "guguducken", "octopus")
	if err != nil {
		panic(err)
	}
	issue, err := GetIssueForRepo(cfg, repo, 1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("issue.Number: %v\n", issue.Number)
	fmt.Printf("issue.Title: %v\n", issue.Title)
	fmt.Printf("issue.CreatedAt: %v\n", issue.CreatedAt)
	fmt.Printf("issue.UpdatedAt: %v\n", issue.UpdatedAt)
}

func TestGetTimeLine(t *testing.T) {
	cfg := config.New(os.Getenv("GITHUB_TOKEN"))
	repo, err := repository.GetRepository(cfg, "guguducken", "octopus")
	if err != nil {
		panic(err)
	}
	issue, err := GetIssueForRepo(cfg, repo, 1)
	if err != nil {
		panic(err)
	}
	events, err := issue.GetTimeLine(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("len(events): %v\n", len(events))
	fmt.Printf("events: %v\n", events)
}
