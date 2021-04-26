package utils

import (
	"sp/test"
	"testing"
)

func TestGetIndexFromKey(t *testing.T) {
	b, a, v, e := GetIndexFromKey("[1]")
	tmp := map[string]interface{}{"b": b, "a": a, "v": v, "e": e}
	test.AssertEqual(t, tmp, map[string]interface{}{
		"b": false, "a": false, "v": 1, "e": nil,
	})

	b, a, v, e = GetIndexFromKey("[:1]")
	tmp = map[string]interface{}{"b": b, "a": a, "v": v, "e": e}
	test.AssertEqual(t, tmp, map[string]interface{}{
		"b": false, "a": true, "v": 1, "e": nil,
	})

	b, a, v, e = GetIndexFromKey("[1:]")
	tmp = map[string]interface{}{"b": b, "a": a, "v": v, "e": e}
	test.AssertEqual(t, tmp, map[string]interface{}{
		"b": true, "a": false, "v": 1, "e": nil,
	})

	b, a, v, e = GetIndexFromKey("[A]")
	tmp = map[string]interface{}{"b": b, "a": a, "v": v}
	test.AssertEqual(t, tmp, map[string]interface{}{
		"b": false, "a": false, "v": 0,
	})
	test.AssertNotEqual(t, e, nil)
}

func TestDataTypeIsJSON(t *testing.T) {
	test.AssertEqual(t, DataTypeIsJSON("j"), true)
	test.AssertEqual(t, DataTypeIsJSON("json"), true)
	test.AssertEqual(t, DataTypeIsJSON("extension"), true)
	test.AssertEqual(t, DataTypeIsJSON("a"), false)
}

func TestDataTypeIsToml(t *testing.T) {
	test.AssertEqual(t, DataTypeIsToml("t"), true)
	test.AssertEqual(t, DataTypeIsToml("toml"), true)
	test.AssertEqual(t, DataTypeIsToml("a"), false)
}

func TestDataTypeIsYaml(t *testing.T) {
	test.AssertEqual(t, DataTypeIsYaml("y"), true)
	test.AssertEqual(t, DataTypeIsYaml("yml"), true)
	test.AssertEqual(t, DataTypeIsYaml("yaml"), true)
	test.AssertEqual(t, DataTypeIsYaml("a"), false)
}

func TestMin(t *testing.T) {
	test.AssertEqual(t, Min(-5, 1), -5)
	test.AssertEqual(t, Min(3, 7), 3)
}

func TestMax(t *testing.T) {
	test.AssertEqual(t, Max(-5, 1), 1)
	test.AssertEqual(t, Max(3, 7), 7)
}
