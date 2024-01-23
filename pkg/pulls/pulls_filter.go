package pulls

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

func (f *Filter) SetStateFilter(state string) {
	f.filter["state"] = state
}
