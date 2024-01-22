package common

const (
	OpenState = iota
)

type Filter interface {
	SetFilter(key string, value string)
	GetFilter() map[string]string
	SetPageInfo(page int, perPage int)
}

// SetStateFilter is works for the both of issue and pull_request
// func (f *Filter) SetStateFilter(state string) *Filter {
// 	if _, ok := f.filter["state"]; !ok {
// 		f.filter["state"] = state
// 	}
// 	return f
// }
