package issue

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/guguducken/octopus/pkg/config"
	"github.com/guguducken/octopus/pkg/repository"
	"github.com/guguducken/octopus/pkg/utils"
)

func GetIssueForRepo(cfg *config.Config, repo *repository.Repository, number int) (i *Issue, err error) {
	i = &Issue{}
	url := utils.URL{
		Endpoint: cfg.GetGithubRestAPI(),
		Path:     fmt.Sprintf("repos/%s/issues/%d", repo.FullName, number),
	}
	reply, err := utils.GetWithRetry(cfg, url)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(reply.Body, i); err != nil {
		return nil, err
	}
	return i, nil
}

func (i Issue) GetTimeLine(cfg *config.Config) (events Events, err error) {
	events = make(Events, 0, 10)
	url := utils.URL{
		RawURL: i.TimelineURL,
	}
	reply, err := utils.GetWithRetry(cfg, url)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(reply.Body, &events); err != nil {
		return nil, err
	}
	return events, err
}

func (i Issue) GetTimeLineByPage(cfg *config.Config, page int) (events Events, err error) {
	events = make(Events, 0, 10)
	url := utils.URL{
		RawURL: i.TimelineURL,
		Params: map[string]string{
			"page":     strconv.Itoa(page),
			"per_page": cfg.GetPerPage(),
		},
	}
	reply, err := utils.GetWithRetry(cfg, url)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(reply.Body, &events); err != nil {
		return nil, err
	}
	return events, err
}

func (i Issue) toJson() (string, error) {
	s, err := json.Marshal(i)
	return string(s), err
}
