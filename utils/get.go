package utils

import (
	"fmt"
	"strings"
)

// Get return an interface on a given query
func Get(data *interface{}, query string) (*interface{}, error) {
	queryKeys := strings.Split(query, ".")
	return getRecursive(data, queryKeys)
}

// getRecursive search recursively an interface on given query keys
func getRecursive(data *interface{}, queryKeys []string) (*interface{}, error) {
	if len(queryKeys) == 0 {
		return data, nil
	}

	if dataA, ok := (*data).([]interface{}); ok && queryKeys[0][0] == '[' {
		if queryKeys[0] == "[*]" {
			return data, nil
		}

		_, _, index, err := GetIndexFromKey(queryKeys[0])
		if err != nil {
			return nil, fmt.Errorf("index value \"%v\" is not a number", index)
		}

		if index >= len(dataA) {
			return nil, fmt.Errorf("index value \"%v\" is out of range", index)
		}

		return getRecursive(&dataA[index], queryKeys[1:])
	} else if dataM, ok := (*data).(map[string]interface{}); ok && queryKeys[0][0] != '[' {
		for k, v := range dataM {
			if k == queryKeys[0] {
				return getRecursive(&v, queryKeys[1:])
			}
		}
	}

	return nil, fmt.Errorf("key \"%s\" not found", queryKeys[0])
}
