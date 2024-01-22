package pulls

import (
	"encoding/json"
	"strconv"

	"github.com/guguducken/octopus/pkg/config"
	"github.com/guguducken/octopus/pkg/utils"
)

func (p *PullRequest) ListCommitForPR(cfg *config.Config) (commits []Commit, err error) {
	page := 1
	perPage := cfg.GetPerPage()
	commits = make([]Commit, 0, 10)
	for {
		var commitsPage []Commit
		commitsPage, err = p.ListCommitForPRByPage(cfg, page)
		if err != nil {
			break
		}
		commits = append(commits, commitsPage...)
		page++
		if len(commitsPage) != perPage {
			break
		}
	}
	return commits, err
}

func (p *PullRequest) ListCommitForPRByPage(cfg *config.Config, page int) (commits []Commit, err error) {
	commits = make([]Commit, 0, 10)

	url := utils.URL{
		RawURL: p.CommitsURL,
		Params: map[string]string{
			"page":     strconv.Itoa(page),
			"per_page": strconv.Itoa(cfg.GetPerPage()),
		},
	}
	reply, err := utils.GetWithRetryWithRateCheck(cfg, url)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(reply.Body, &commits)
	return commits, err
}
