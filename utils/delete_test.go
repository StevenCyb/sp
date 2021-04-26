package utils

import (
	"sp/test"
	"testing"
)

func TestDelete(t *testing.T) {
	data := CreateInterfaceFromMap(map[string]interface{}{
		"github_url": "https://github.com/StevenCyb/sp",
		"uppercase":  map[string]interface{}{"a": "A", "b": "B", "x": "X"},
		"count":      []interface{}{1, 2, 9},
	})
	dataResult := CreateInterfaceFromMap(map[string]interface{}{
		"uppercase": map[string]interface{}{"a": "A", "b": "B", "x": "X"},
		"count":     []interface{}{1, 2, 9},
	})
	err := Delete(data, "github_url")
	test.AssertEqual(t, err, nil)
	test.AssertEqual(t, *data, *dataResult)

	data = CreateInterfaceFromMap(map[string]interface{}{
		"github_url": "https://github.com/StevenCyb/sp",
		"uppercase":  map[string]interface{}{"a": "A", "b": "B", "x": "X"},
		"count":      []interface{}{1, 2, 9},
	})
	dataResult = CreateInterfaceFromMap(map[string]interface{}{
		"github_url": "https://github.com/StevenCyb/sp",
		"uppercase":  map[string]interface{}{"a": "A", "b": "B"},
		"count":      []interface{}{1, 2, 9},
	})
	err = Delete(data, "uppercase.x")
	test.AssertEqual(t, err, nil)
	test.AssertEqual(t, *data, *dataResult)

	data = CreateInterfaceFromMap(map[string]interface{}{
		"github_url": "https://github.com/StevenCyb/sp",
		"uppercase":  map[string]interface{}{"a": "A", "b": "B", "x": "X"},
		"count":      []interface{}{1, 2, 9},
	})
	dataResult = CreateInterfaceFromMap(map[string]interface{}{
		"github_url": "https://github.com/StevenCyb/sp",
		"uppercase":  map[string]interface{}{"a": "A", "b": "B", "x": "X"},
		"count":      []interface{}{},
	})
	err = Delete(data, "count.[*]")
	test.AssertEqual(t, err, nil)
	test.AssertEqual(t, *data, *dataResult)

	data = CreateInterfaceFromMap(map[string]interface{}{
		"github_url": "https://github.com/StevenCyb/sp",
		"uppercase":  map[string]interface{}{"a": "A", "b": "B", "x": "X"},
		"count":      []interface{}{1, 2, 9},
	})
	dataResult = CreateInterfaceFromMap(map[string]interface{}{
		"github_url": "https://github.com/StevenCyb/sp",
		"uppercase":  map[string]interface{}{"a": "A", "b": "B", "x": "X"},
		"count":      []interface{}{1, 2},
	})
	err = Delete(data, "count.[2]")
	test.AssertEqual(t, err, nil)
	test.AssertEqual(t, *data, *dataResult)
}
