package common

import (
	"fmt"
	"testing"
)

func TestFilter_AddMileStoneFilter(t *testing.T) {
	filter := NewFilter()
	filter.SetMileStoneFilter(&Milestone{
		Title: "test",
	})
	fmt.Printf("filter: %v\n", filter.GetFilter())
	filterMap := filter.GetFilter()
	filterMap["aaa"] = "bbbb"
	fmt.Printf("filter.GetFilter(): %v\n", filter.GetFilter())
	fmt.Printf("filterMap: %v\n", filterMap)
}
