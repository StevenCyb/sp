package utils

import (
	"sp/test"
	"testing"
)

func TestMerge(t *testing.T) {
	data1 := CreateInterfaceFromMap(map[string]interface{}{
		"github_url": "https://github.com/StevenCyb/sp",
		"uppercase":  map[string]string{"a": "A", "b": "B", "x": "X"},
		"count":      []int{1, 2, 9},
	})
	data2 := CreateInterfaceFromMap(map[string]interface{}{
		"github_url":  "https://github.com/StevenCyb/sp",
		"github_user": "StevenCyb",
		"uppercase":   map[string]string{"a": "A", "b": "B", "c": "C"},
		"count":       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	})
	dataResult := CreateInterfaceFromMap(map[string]interface{}{
		"github_url":  "https://github.com/StevenCyb/sp",
		"github_user": "StevenCyb",
		"uppercase":   map[string]string{"a": "A", "b": "B", "c": "C", "x": "X"},
		"count":       []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	})
	Merge(data1, data2, false)
	test.AssertEqual(t, *data2, *dataResult)

	data1 = CreateInterfaceFromArray([]interface{}{1, 2, 3})
	data2 = CreateInterfaceFromArray([]interface{}{4, 5, 6})
	dataResult = CreateInterfaceFromArray([]interface{}{1, 2, 3})
	Merge(data1, data2, false)
	test.AssertEqual(t, *data2, *dataResult)

	data1 = CreateInterfaceFromArray([]interface{}{1, 2, 3})
	data2 = CreateInterfaceFromArray([]interface{}{4, 5, 6})
	dataResult = CreateInterfaceFromArray([]interface{}{1, 2, 3, 4, 5, 6})
	Merge(data1, data2, true)
	test.AssertEqual(t, *data2, *dataResult)
}
