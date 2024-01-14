package common

import (
	"maps"
	"strconv"
)

type Filter struct {
	filter map[string]string
}

func NewFilter() *Filter {
	return &Filter{
		filter: make(map[string]string, 10),
	}

}

// SetStateFilter is works for the both of issue and pull_request
func (f *Filter) SetStateFilter(state string) *Filter {
	if _, ok := f.filter["state"]; !ok {
		f.filter["state"] = state
	}
	return f
}

// SetMileStoneFilter only works for the issue
func (f *Filter) SetMileStoneFilter(milestone *Milestone) *Filter {
	if _, ok := f.filter["milestone"]; !ok {
		milestoneFilter := milestone.Title
		if milestone.Number != 0 {
			milestoneFilter = strconv.Itoa(milestone.Number)
		}
		f.filter["milestone"] = milestoneFilter
	}

	return f
}

// SetLabelsFilter only works for the issue
func (f *Filter) SetLabelsFilter(labels []Label) *Filter {
	if _, ok := f.filter["labels"]; !ok {
		labelFilter := ""
		for i := 0; i < len(labels); i++ {
			labelFilter += labels[i].Name + ","
		}
		if len(labelFilter) != 0 {
			labelFilter = labelFilter[:len(labelFilter)-1]
		}
		f.filter["labels"] = labelFilter
	}
	return f
}

// SetCreatorFilter only works for the issue
func (f *Filter) SetCreatorFilter(user *User) *Filter {
	if _, ok := f.filter["creator"]; !ok {
		f.filter["creator"] = user.Login
	}
	return f
}

func (f *Filter) GetFilter() map[string]string {
	return maps.Clone(f.filter)
}
