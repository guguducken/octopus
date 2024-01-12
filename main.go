package main

import (
	"fmt"
	"os"

	"github.com/guguducken/octopus/pkg/auth"
	"github.com/guguducken/octopus/pkg/config"
)

func main() {
	cfg := config.New(os.Getenv("GITHUB_TOKEN"))
	authedUser, err := auth.GetAuthenticatedUser(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello, %s!\n", authedUser.Login)
	fmt.Printf("authedUser.Bio: %v\n", authedUser.Bio)
}
