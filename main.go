package main

import (
	"fmt"
	"octopus/pkg/auth"
	"octopus/pkg/config"
	"os"
)

func main() {
	cfg := config.New(os.Getenv("GITHUB_TOKEN"))
	authedUser, err := auth.GetAuthenticatedUser(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("authedUser: %v\n", authedUser)
}
