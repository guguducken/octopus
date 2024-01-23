package pulls

import (
	"encoding/json"
	"fmt"

	"github.com/guguducken/octopus/pkg/common"
	"github.com/guguducken/octopus/pkg/config"
	"github.com/guguducken/octopus/pkg/repository"
	"github.com/guguducken/octopus/pkg/utils"
)

func GetPullForRepo(cfg *config.Config, repo *repository.Repository, pullNumber int) (pull *PullRequest, err error) {
	pull = &PullRequest{}
	url := utils.URL{
		Endpoint: cfg.GetGithubRestAPI(),
		Path:     fmt.Sprintf("repos/%s/pulls/%d", repo.FullName, pullNumber),
	}

	reply, err := utils.GetWithRetryWithRateCheck(cfg, url)
	if err = json.Unmarshal(reply.Body, pull); err != nil {
		return pull, err
	}
	return pull, err
}

func ListPullsForRepoByFilter(cfg *config.Config, repo *repository.Repository, filter common.Filter) (pulls []PullRequest, err error) {
	pulls = make([]PullRequest, 0, 10)

	url := utils.URL{
		Endpoint: cfg.GetGithubRestAPI(),
		Path:     fmt.Sprintf("repos/%s/pulls", repo.FullName),
		Params:   filter.GetFilter(),
	}

	reply, err := utils.GetWithRetryWithRateCheck(cfg, url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(reply.Body, &pulls)
	return pulls, err
}

func ListPullsForRepo(cfg *config.Config, repo *repository.Repository, filter common.Filter) (pulls []PullRequest, err error) {
	pulls = make([]PullRequest, 0, 10)
	page := 1
	for {
		var pullsPerPage []PullRequest
		pullsPerPage, err = ListPullsForRepoByPage(cfg, repo, page, filter)
		if err != nil {
			break
		}
		pulls = append(pulls, pullsPerPage...)
		page++
		if len(pullsPerPage) != cfg.GetPerPage() {
			break
		}
	}
	return pulls, err
}

func ListPullsForRepoByPage(cfg *config.Config, repo *repository.Repository, page int, filter common.Filter) (pulls []PullRequest, err error) {
	if filter == nil {
		filter = NewFilter()
	}
	filter.SetPageInfo(page, cfg.GetPerPage())
	return ListPullsForRepoByFilter(cfg, repo, filter)
}

func (p *PullRequest) ToJson() (string, error) {
	s, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(s), err
}
