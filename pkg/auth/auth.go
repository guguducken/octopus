package auth

import (
	"encoding/json"

	"github.com/guguducken/octopus/pkg/config"
	"github.com/guguducken/octopus/pkg/user"
	"github.com/guguducken/octopus/pkg/utils"
)

func GetAuthenticatedUser(cfg *config.Config) (user user.User, err error) {
	url := utils.URL{
		Endpoint: cfg.ApiConfig.GitHubRestAPI,
		Path:     "user",
	}
	reply, err := utils.Get(cfg, url)
	if err != nil {
		return user, err
	}
	if err := json.Unmarshal(reply.Body, &user); err != nil {
		return user, err
	}
	return user, nil

}
