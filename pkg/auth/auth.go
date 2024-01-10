package auth

import (
	"encoding/json"
	"octopus/pkg/config"
	"octopus/pkg/user"
	"octopus/pkg/utils"
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
