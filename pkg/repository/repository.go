package repository

import (
	"encoding/json"
	"fmt"

	"github.com/guguducken/octopus/pkg/config"
	"github.com/guguducken/octopus/pkg/utils"
)

func GetRepository(cfg *config.Config, repoOwner, repoName string) (repo *Repository, err error) {
	repo = &Repository{}
	url := utils.URL{
		Endpoint: cfg.ApiConfig.GitHubRestAPI,
		Path:     fmt.Sprintf("repos/%s/%s", repoOwner, repoName),
	}
	reply, err := utils.Get(cfg, url)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(reply.Body, repo); err != nil {
		return nil, err
	}
	return repo, nil
}

func (r Repository) ToJson() (string, error) {
	s, err := json.Marshal(r)
	return string(s), err
}
