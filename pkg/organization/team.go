package organization

import (
	"encoding/json"
	"fmt"
	"github.com/guguducken/octopus/pkg/common"
	"github.com/guguducken/octopus/pkg/config"
	"github.com/guguducken/octopus/pkg/utils"
	"strconv"
)

type Members []common.User
type Teams []Team

func (o Organization) ListTeams(cfg *config.Config) (teams Teams, err error) {
	teams = make(Teams, 0, 10)
	page := 1
	perPage := cfg.GetPerPage()

	for {
		var teamsPage []Team
		teamsPage, err = o.ListTeamsByPage(cfg, page)
		if err != nil {
			break
		}
		teams = append(teams, teamsPage...)
		if len(teamsPage) == perPage {
			page++
			continue
		}
		break
	}

	return teams, err

}

func (o Organization) ListTeamsByPage(cfg *config.Config, page int) (teams Teams, err error) {
	teams = make(Teams, 0, 10)

	url := utils.URL{
		Endpoint: cfg.GetGithubRestAPI(),
		Path:     fmt.Sprintf("orgs/%s/teams", o.Login),
		Params: map[string]string{
			"per_page": strconv.Itoa(cfg.GetPerPage()),
			"page":     strconv.Itoa(page),
		},
	}
	reply, err := utils.GetWithRetryWithRateCheck(cfg, url)
	if err != nil {
		return teams, err
	}
	err = json.Unmarshal(reply.Body, &teams)
	return teams, err
}

func (o Organization) ListTeamMembers(cfg *config.Config, team Team, filter *common.Filter) (members Members, err error) {
	members = make([]common.User, 0, 20)
	page := 1
	perPage := cfg.GetPerPage()

	for {
		var memberPage []common.User
		memberPage, err = o.ListTeamMembersByPage(cfg, team, page, filter)
		if err != nil {
			break
		}
		members = append(members, memberPage...)
		if len(memberPage) == perPage {
			page++
			continue
		}
		break
	}

	return members, err
}

func (o Organization) ListTeamMembersByPage(cfg *config.Config, team Team, page int, filter *common.Filter) (members Members, err error) {
	members = make([]common.User, 0, 10)
	params := map[string]string{
		"per_page": strconv.Itoa(cfg.GetPerPage()),
		"page":     strconv.Itoa(page),
	}

	if filter != nil {
		fi := filter.GetFilter()
		for key, value := range fi {
			if _, ok := params[key]; !ok {
				params[key] = value
			}
		}
	}

	url := utils.URL{
		Endpoint: cfg.GetGithubRestAPI(),
		Path:     fmt.Sprintf("orgs/%s/teams/%s/members", o.Login, team.Slug),
		Params:   params,
	}
	reply, err := utils.GetWithRetryWithRateCheck(cfg, url)
	if err != nil {
		return members, err
	}
	err = json.Unmarshal(reply.Body, &members)
	return members, err
}

func (m Members) FindUser(u common.User) *common.User {
	for _, member := range m {
		if member.Login == u.Login {
			return &member
		}
	}
	return nil
}

func (o Organization) GetTeamsByUser(cfg *config.Config, u common.User) (teams Teams, err error) {
	teams = make(Teams, 0, 5)

	orgTeams, err := o.ListTeams(cfg)
	if err != nil {
		return nil, err
	}
	for ind, t := range orgTeams {
		var members Members
		members, err = o.ListTeamMembers(cfg, t, nil)
		if err != nil {
			break
		}
		if members.FindUser(u) != nil {
			teams = append(teams, orgTeams[ind])
			break
		}
	}
	return teams, err
}
