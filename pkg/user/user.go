package user

import (
	"encoding/json"
	"fmt"

	"github.com/guguducken/octopus/pkg/common"
	"github.com/guguducken/octopus/pkg/config"
	"github.com/guguducken/octopus/pkg/utils"
)

func GetSelf(cfg *config.Config) (user *common.User, err error) {
	user = &common.User{}
	url := utils.URL{
		Endpoint: cfg.ApiConfig.GitHubRestAPI,
		Path:     "user",
	}
	reply, err := utils.Get(cfg, url)
	if err != nil {
		return user, err
	}
	fmt.Printf("string(reply.Body): %v\n", string(reply.Body))
	if err := json.Unmarshal(reply.Body, user); err != nil {
		return user, err
	}
	return user, nil
}
