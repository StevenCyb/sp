package utils

import (
	"sp/test"
	"testing"
)

func TestGet(t *testing.T) {
	data := CreateInterfaceFromMap(map[string]interface{}{
		"github_url": "https://github.com/StevenCyb/sp",
		"uppercase":  map[string]interface{}{"a": "A", "b": "B", "x": "X"},
		"count":      []interface{}{1, 2, 9},
	})

	value, err := Get(data, "github_url")
	test.AssertEqual(t, err, nil)
	test.AssertEqual(t, *value, "https://github.com/StevenCyb/sp")

	value, err = Get(data, "uppercase.x")
	test.AssertEqual(t, err, nil)
	test.AssertEqual(t, *value, "X")

	value, err = Get(data, "count.[1]")
	test.AssertEqual(t, err, nil)
	test.AssertEqual(t, *value, 2)

	value, err = Get(data, "count")
	test.AssertEqual(t, err, nil)
	test.AssertEqual(t, *value, []interface{}{1, 2, 9})

	value, err = Get(data, "count.[*]")
	test.AssertEqual(t, err, nil)
	test.AssertEqual(t, *value, []interface{}{1, 2, 9})
}
