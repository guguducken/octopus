package issue

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/guguducken/octopus/pkg/common"
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

	reply, err := utils.GetWithRetryWithRateCheck(cfg, url)
	if err = json.Unmarshal(reply.Body, i); err != nil {
		return i, err
	}
	return i, err
}

func (i Issue) GetTimeLine(cfg *config.Config) (events Events, err error) {
	events = make(Events, 0, 10)
	page := 1
	for {
		var eventPerPage Events
		eventPerPage, err = i.GetTimeLineByPage(cfg, page)
		if err != nil {
			return events, err
		}
		events = append(events, eventPerPage...)
		page++
		if len(eventPerPage) != cfg.GetPerPage() {
			break
		}
	}

	return events, err
}

func (i Issue) GetTimeLineByPage(cfg *config.Config, page int) (events Events, err error) {
	events = make(Events, 0, 10)
	url := utils.URL{
		RawURL: i.TimelineURL,
		Params: map[string]string{
			"page":     strconv.Itoa(page),
			"per_page": strconv.Itoa(cfg.GetPerPage()),
		},
	}
	reply, err := utils.GetWithRetryWithRateCheck(cfg, url)
	if err = json.Unmarshal(reply.Body, &events); err != nil {
		return nil, err
	}

	return events, err
}

func ListIssueForRepoByPage(cfg *config.Config, repo *repository.Repository, page int, filter *common.Filter) (issues []Issue, err error) {
	return listIssueForRepoByPage(cfg, repo, page, filter)
}

func listIssueForRepo(cfg *config.Config, repo *repository.Repository, filter *common.Filter) (issues []Issue, err error) {
	issues = make([]Issue, 0, 20)

	page := 1
	for {
		var issuesPerPage []Issue
		issuesPerPage, err = listIssueForRepoByPage(cfg, repo, page, filter)
		if err != nil {
			return issues, err
		}
		issues = append(issues, issuesPerPage...)
		page++
		if len(issuesPerPage) != cfg.GetPerPage() {
			break
		}
	}
	return issues, err
}

func listIssueForRepoByPage(cfg *config.Config, repo *repository.Repository, page int, filter *common.Filter) (issues []Issue, err error) {
	issues = make([]Issue, 0, 20)

	// prepare url params
	params := map[string]string{
		"page":     strconv.Itoa(page),
		"per_page": strconv.Itoa(cfg.GetPerPage()),
	}
	if filter != nil {
		filterMap := filter.GetFilter()
		for key, value := range filterMap {
			params[key] = value
		}
	}
	url := utils.URL{
		Endpoint: cfg.GetGithubRestAPI(),
		Path:     fmt.Sprintf("repos/%s/issues", repo.FullName),
		Params:   params,
	}

	reply, err := utils.GetWithRetryWithRateCheck(cfg, url)
	if err = json.Unmarshal(reply.Body, &issues); err != nil {
		return issues, err
	}
	return issues, err
}

func ListIssueForRepoByLabels(cfg *config.Config, repo *repository.Repository, labels []common.Label) (issues []Issue, err error) {
	filter := common.NewFilter().SetLabelsFilter(labels)
	return listIssueForRepo(cfg, repo, filter)
}

func ListIssueForRepoByMilestone(cfg *config.Config, repo *repository.Repository, milestone *common.Milestone) (issues []Issue, err error) {
	filter := common.NewFilter().SetMileStoneFilter(milestone)
	return listIssueForRepo(cfg, repo, filter)
}

func ListIssueForRepo(cfg *config.Config, repo *repository.Repository) (issues []Issue, err error) {
	return listIssueForRepo(cfg, repo, nil)
}

func ListIssueForRepoByFilter(cfg *config.Config, repo *repository.Repository, filter *common.Filter) (issues []Issue, err error) {
	return listIssueForRepo(cfg, repo, filter)
}

func (i Issue) ToJson() (string, error) {
	s, err := json.Marshal(i)
	return string(s), err
}
