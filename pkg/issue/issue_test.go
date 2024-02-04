package issue

import (
	"fmt"
	"os"
	"sort"
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

func TestListIssueForRepoByFilter(t *testing.T) {
	cfg := config.New(os.Getenv("GITHUB_TOKEN"))
	repo, err := repository.GetRepository(cfg, "matrixorigin", "matrixone")
	if err != nil {
		panic(err)
	}
	filter := NewFilter()
	filter.SetLabelsFilter([]string{
		"kind/bug",
	})
	issues, err := ListIssueForRepoByFilter(cfg, repo, filter)
	if err != nil {
		panic(err)
	}
	fmt.Printf("len(issues): %v\n", len(issues))
	sort.Slice(issues, func(i, j int) bool {
		return issues[i].Number < issues[j].Number
	})
	for _, i := range issues {
		fmt.Printf("i.Number: %v\n", i.Number)
	}
}

func TestListIssueForRepo(t *testing.T) {
	cfg := config.New(os.Getenv("GITHUB_TOKEN"))
	repo, err := repository.GetRepository(cfg, "matrixorigin", "matrixone")
	if err != nil {
		panic(err)
	}
	filter := NewFilter()
	filter.SetLabelsFilter([]string{
		"kind/bug",
	})
	issues, err := ListIssueForRepo(cfg, repo, filter)
	if err != nil {
		panic(err)
	}
	fmt.Printf("len(issues): %v\n", len(issues))
	issues = issues.RemovePullRequest()
	fmt.Printf("len(issues): %v\n", len(issues))
	for _, issue := range issues {
		fmt.Printf("issue.Number: %v\n", issue.Number)
	}
}
