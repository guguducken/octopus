package user

import (
	"encoding/json"

	"github.com/guguducken/octopus/pkg/common"
	"github.com/guguducken/octopus/pkg/config"
	"github.com/guguducken/octopus/pkg/utils"
)

func GetSelf(cfg *config.Config) (user *common.User, err error) {
	user = &common.User{}
	url := utils.URL{
		Endpoint: cfg.GetGithubRestAPI(),
		Path:     "user",
	}
	reply, err := utils.GetWithRetry(cfg, url)
	if err != nil {
		return user, err
	}
	if err := json.Unmarshal(reply.Body, user); err != nil {
		return user, err
	}
	return user, nil
}
