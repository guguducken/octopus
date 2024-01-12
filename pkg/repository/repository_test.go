package repository

import (
	"fmt"
	"os"
	"testing"

	"github.com/guguducken/octopus/pkg/config"
)

func TestGetRepository(t *testing.T) {
	cfg := config.New(os.Getenv("GITHUB_TOKEN"))
	repo, err := GetRepository(cfg, "guguducken", "octopus")
	if err != nil {
		panic(err)
	}
	fmt.Printf("repo.Name: %v\n", repo.Name)
	fmt.Printf("repo.FullName: %v\n", repo.FullName)
	fmt.Println("=====================================")
	fmt.Println(repo.ToJson())

}
