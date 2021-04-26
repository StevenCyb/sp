package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// GetIndexFromKey return an int index out of query part if possible
func GetIndexFromKey(key string) (bool, bool, int, error) {
	before := false
	after := false

	key = strings.Replace(key, "[", "", -1)
	key = strings.Replace(key, "]", "", -1)

	if len(key) == 0 {
		return false, false, -1, fmt.Errorf("not valid")
	}

	if key[0] == ':' {
		after = true
	} else if key[len(key)-1] == ':' {
		before = true
	}

	key = strings.Replace(key, ":", "", -1)
	value, err := strconv.Atoi(key)

	return before, after, value, err
}

// SetArrayInDataTo set a new array value of data processing (on actual or previous pointer)
func SetArrayInDataTo(data *interface{}, key string, previousData *interface{}, previousKey string, newData interface{}) error {
	if previousData != nil {
		if previousDataM, ok := (*previousData).(map[string]interface{}); ok {
			previousDataM[previousKey] = newData
			return nil
		} else if previousDataA, ok := (*previousData).([]interface{}); ok {
			if key == "[*]" {
				*previousData = newData
				return nil
			}
			_, _, index, err := GetIndexFromKey(key)
			if err != nil {
				return fmt.Errorf("index value of \"%s\", \"%v\" is not a number", key, index)
			}

			if index >= len(previousDataA) {
				return fmt.Errorf("index value \"%v\" is out of range", index)
			}

			previousDataA[index] = newData
			return nil
		}
	} else {
		*data = newData
		return nil
	}

	return fmt.Errorf("key \"%s\" not found", key)
}

// StringCaster try to create int, float etc. from string if possible
func StringCaster(data string) interface{} {
	if c, err := strconv.Atoi(data); err == nil {
		return c
	} else if c, err := strconv.ParseBool(data); err == nil {
		return c
	} else if c, err := strconv.ParseFloat(data, 64); err == nil {
		return c
	} else if c, err := strconv.ParseInt(data, 10, 64); err == nil {
		return c
	}

	return data
}

// DataTypeIsJSON checks if given data type match json
func DataTypeIsJSON(dataType string) bool {
	return dataType == "j" || dataType == "json" || dataType == "extension"
}

// DataTypeIsToml checks if given data type match toml
func DataTypeIsToml(dataType string) bool {
	return dataType == "t" || dataType == "toml"
}

// DataTypeIsYaml checks if given data type match yaml
func DataTypeIsYaml(dataType string) bool {
	return dataType == "y" || dataType == "yml" || dataType == "yaml"
}

// Min value of two int
func Min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

// Max value of two int
func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

type Dummy struct {
	internal interface{}
}

func CreateEmptyMap() *interface{} {
	data := Dummy{
		internal: make(map[string]interface{}),
	}

	return &data.internal
}

func CreateInterfaceFromMap(m map[string]interface{}) *interface{} {
	data := Dummy{
		internal: m,
	}

	return &data.internal
}

func CreateEmptyArray() *interface{} {
	data := Dummy{
		internal: []interface{}{},
	}

	return &data.internal
}

func CreateInterfaceFromArray(a []interface{}) *interface{} {
	data := Dummy{
		internal: a,
	}

	return &data.internal
}
