package main

import (
	"fmt"
	"os"
	"time"

	"github.com/guguducken/octopus/pkg/auth"
	"github.com/guguducken/octopus/pkg/config"
)

func main() {
	cfg := config.New(os.Getenv("GITHUB_TOKEN"))
	cfg.SetRetryTimes(10).SetTimeOut(30 * time.Second).SetRetryDuration(5 * time.Second)
	authedUser, err := auth.GetAuthenticatedUser(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello, %s!\n", authedUser.Login)
}
