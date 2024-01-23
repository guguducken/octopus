package issue

import (
	"maps"
	"strconv"

	"github.com/guguducken/octopus/pkg/common"
)

type Filter struct {
	filter map[string]string
}

func (f *Filter) SetFilter(key string, value string) {
	f.filter[key] = value
}

func (f *Filter) GetFilter() map[string]string {
	return maps.Clone(f.filter)
}

func (f *Filter) SetPageInfo(page int, perPage int) {
	f.filter["page"] = strconv.Itoa(page)
	f.filter["per_page"] = strconv.Itoa(perPage)
}

func NewFilter() *Filter {
	return &Filter{
		filter: make(map[string]string, 10),
	}
}

func (f *Filter) SetLabelsFilter(labels []common.Label) {
	labelFilter := ""
	for i := 0; i < len(labels); i++ {
		labelFilter += labels[i].Name + ","
	}
	if len(labelFilter) != 0 {
		labelFilter = labelFilter[:len(labelFilter)-1]
	}
	f.filter["labels"] = labelFilter
}

func (f *Filter) SetCreatorFilter(user *common.User) {
	f.filter["creator"] = user.Login
}

// SetMileStoneFilter only works for the issue
func (f *Filter) SetMileStoneFilter(milestone *common.Milestone) {
	milestoneFilter := milestone.Title
	if milestone.Number != 0 {
		milestoneFilter = strconv.Itoa(milestone.Number)
	}
	f.filter["milestone"] = milestoneFilter
}

func (f *Filter) SetStateFilter(state string) {
	f.filter["state"] = state
}
