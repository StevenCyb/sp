package utils

import (
	"fmt"
	"strings"
)

// Delete an interface on a given query
func Delete(data *interface{}, query string) error {
	queryKeys := strings.Split(query, ".")
	return deleteRecursive(data, queryKeys, nil, "")
}

// deleteRecursive search recursively and delete an interface on given query keys
func deleteRecursive(data *interface{}, queryKeys []string, previousData *interface{}, previousKey string) error {
	isLastKey := len(queryKeys) == 1

	if dataA, ok := (*data).([]interface{}); ok && queryKeys[0][0] == '[' {
		if queryKeys[0] == "[*]" {
			return SetArrayInDataTo(data, queryKeys[0], previousData, previousKey, []interface{}{})
		}

		_, _, index, err := GetIndexFromKey(queryKeys[0])
		if err != nil {
			return fmt.Errorf("index value of \"%s\", \"%v\" is not a number", queryKeys[0], index)
		}

		if index >= len(dataA) {
			return fmt.Errorf("index value \"%v\" is out of range", index)
		}

		if isLastKey {
			newData := append(dataA[:index], dataA[index+1:]...)
			return SetArrayInDataTo(data, queryKeys[0], previousData, previousKey, newData)
		}

		return deleteRecursive(&dataA[index], queryKeys[1:], data, queryKeys[0])
	} else if dataM, ok := (*data).(map[string]interface{}); ok && queryKeys[0][0] != '[' {
		for k, v := range dataM {
			if k == queryKeys[0] {
				if isLastKey {
					delete(dataM, queryKeys[0])
					return nil
				}
				return deleteRecursive(&v, queryKeys[1:], data, queryKeys[0])
			}
		}
	}

	return fmt.Errorf("key \"%s\" not found", queryKeys[0])
}
