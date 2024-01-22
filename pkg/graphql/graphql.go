package graphql

import (
	"github.com/guguducken/octopus/pkg/config"
	"github.com/guguducken/octopus/pkg/utils"
)

type GraphQL interface {
	Exec(cfg *config.Config, query string) ([]byte, error)
}

type GitHubGraphQL struct{}

func NewGitHubGraphQL() *GitHubGraphQL {
	return &GitHubGraphQL{}
}

func (gg *GitHubGraphQL) Exec(cfg *config.Config, query string) ([]byte, error) {
	url := utils.URL{
		Endpoint: cfg.GetGithubGraphQLAPI(),
	}
	reply, err := utils.PostWithRetryWithRateCheck(cfg, url, query)
	if err != nil {
		return nil, err
	}
	return reply.Body, err
}
