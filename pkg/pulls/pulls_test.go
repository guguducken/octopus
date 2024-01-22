package pulls

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/guguducken/octopus/pkg/config"
	"github.com/guguducken/octopus/pkg/repository"
)

func TestGetPullForRepo(t *testing.T) {
	cfg := config.New(os.Getenv("GITHUB_TOKEN"))
	repo, err := repository.GetRepository(cfg, "matrixorigin", "matrixone")
	if err != nil {
		panic(err)
	}
	pull, err := GetPullForRepo(cfg, repo, 14175)
	fmt.Printf("pull.Number: %v\n", pull.Number)
	fmt.Printf("pull.Title: %v\n", pull.Title)
	pullStr, err := pull.ToJson()
	if err != nil {
		panic(err)
	}
	fmt.Printf("pullStr: %v\n", pullStr)
}

func TestListPullsForRepo(t *testing.T) {
	cfg := config.New(os.Getenv("GITHUB_TOKEN"))
	cfg.SetRetryDuration(5 * time.Second)
	repo, err := repository.GetRepository(cfg, "matrixorigin", "matrixone")
	if err != nil {
		panic(err)
	}
	pulls, err := ListPullsForRepo(cfg, repo)
	if err != nil {
		panic(err)
	}
	fmt.Printf("len(pulls): %v\n", len(pulls))
	for _, pull := range pulls {
		fmt.Printf("pull.Number: %v\n", pull.Number)
	}
}
