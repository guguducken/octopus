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

	if i.PullRequest != nil {
		return nil, ErrNotIssue
	}
	i.Repository = repo
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

func ListIssueForRepoByPage(cfg *config.Config, repo *repository.Repository, page int, filter common.Filter) (issues []Issue, err error) {
	issues, err = listIssueForRepoByPage(cfg, repo, page, filter)
	return FilterPullRequest(issues), err
}

func listIssueForRepo(cfg *config.Config, repo *repository.Repository, filter common.Filter) (issues []Issue, err error) {
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
	return FilterPullRequest(issues), err
}

func listIssueForRepoByPage(cfg *config.Config, repo *repository.Repository, page int, filter common.Filter) (issues []Issue, err error) {
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
	// repo content
	for i := 0; i < len(issues); i++ {
		issues[i].Repository = repo
	}
	// we will do not filter out pull request in there
	return issues, err
}

func ListIssueForRepoByLabels(cfg *config.Config, repo *repository.Repository, labels []common.Label) (issues []Issue, err error) {
	filter := NewFilter()
	filter.SetLabelsFilter(labels)
	return listIssueForRepo(cfg, repo, filter)
}

func ListIssueForRepoByMilestone(cfg *config.Config, repo *repository.Repository, milestone *common.Milestone) (issues []Issue, err error) {
	filter := NewFilter()
	filter.SetMileStoneFilter(milestone)
	return listIssueForRepo(cfg, repo, filter)
}

func ListIssueForRepo(cfg *config.Config, repo *repository.Repository) (issues []Issue, err error) {
	return listIssueForRepo(cfg, repo, nil)
}

func ListIssueForRepoByFilter(cfg *config.Config, repo *repository.Repository, filter common.Filter) (issues []Issue, err error) {
	return listIssueForRepo(cfg, repo, filter)
}

func (i Issue) ToJson() (string, error) {
	s, err := json.Marshal(i)
	return string(s), err
}

func FilterPullRequest(issues []Issue) []Issue {
	left, right := 0, len(issues)-1
	for left < right {
		if issues[left].PullRequest == nil {
			left++
			continue
		}
		if issues[right].PullRequest != nil {
			right--
			continue
		}
		issues[left], issues[right] = issues[right], issues[left]
	}
	return issues[:right]
}
