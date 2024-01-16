package organization

import (
	"fmt"
	"github.com/guguducken/octopus/pkg/config"
	"github.com/guguducken/octopus/pkg/user"
	"os"
	"testing"
)

func TestListTeams(t *testing.T) {
	cfg := config.New(os.Getenv("GITHUB_TOKEN"))
	org := Organization{
		Login: "matrixorigin",
	}
	teams, err := org.ListTeams(cfg)
	if err != nil {
		panic(err)
	}
	for _, team := range teams {
		fmt.Printf("team.Name: %v\n", team.Name)
		fmt.Printf("team.URL: %v\n", team.URL)
	}
}

func TestListTeamMembers(t *testing.T) {
	cfg := config.New(os.Getenv("GITHUB_TOKEN"))
	org := Organization{
		Login: "matrixorigin",
	}
	selfUser, err := user.GetSelf(cfg)
	if err != nil {
		panic(err)
	}
	teams, err := org.ListTeams(cfg)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(teams); i++ {
		members, err := org.ListTeamMembers(cfg, teams[i], nil)
		if err != nil {
			panic(err)
		}
		if members.FindUser(*selfUser) != nil {
			fmt.Printf("teams[i].Name: %v\n", teams[i].Name)
		}
	}
}

func TestGetTeamByUser(t *testing.T) {
	cfg := config.New(os.Getenv("GITHUB_TOKEN"))
	org := Organization{
		Login: "matrixorigin",
	}
	selfUser, err := user.GetSelf(cfg)
	if err != nil {
		panic(err)
	}
	teams, err := org.GetTeamsByUser(cfg, *selfUser)
	if err != nil {
		panic(err)
	}
	fmt.Printf("teams.Name: %v\n", teams)
}
