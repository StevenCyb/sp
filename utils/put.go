package utils

import (
	"fmt"
	"strings"
)

// Put value on a given query
func Put(data *interface{}, query string, value interface{}) error {
	queryKeys := strings.Split(query, ".")
	return putRecursive(data, queryKeys, nil, "", value)
}

// putRecursive search recursively put value on given query keys
func putRecursive(data *interface{}, queryKeys []string, previousData *interface{}, previousKey string, value interface{}) error {
	isLastKey := len(queryKeys) == 1

	if dataA, ok := (*data).([]interface{}); ok && queryKeys[0][0] == '[' {
		if isLastKey {
			newData := dataA[:]

			if queryKeys[0] == "[+]" {
				newData = append(newData, value)
			} else {
				before, after, index, err := GetIndexFromKey(queryKeys[0])
				if err != nil {
					return fmt.Errorf("index value of \"%s\", \"%v\" is not a number", queryKeys[0], index)
				}

				if index >= len(dataA) {
					return fmt.Errorf("index value \"%v\" is out of range", index)
				}

				if before {
					newData = append(newData[:index+1], newData[index:]...)
					newData[index] = value
				} else if after {
					newData = append(newData[:index+1], newData[index:]...)
					newData[index+1] = value
				} else {
					newData[index] = value
				}
			}

			return SetArrayInDataTo(data, queryKeys[0], previousData, previousKey, newData)
		}

		_, _, index, err := GetIndexFromKey(queryKeys[0])
		if queryKeys[0] == "[+]" || err == nil && index >= len(dataA) {
			var newValue interface{}
			if queryKeys[1][0] == '[' {
				newValue = []interface{}{}
			} else {
				newValue = map[string]interface{}{}
			}

			newData := dataA[:]
			newData = append(newData, newValue)

			err = SetArrayInDataTo(data, queryKeys[0], previousData, previousKey, newData)
			if err != nil {
				return fmt.Errorf("failed to add \"%s\" object to array", queryKeys[0])
			}

			return putRecursive(&newValue, queryKeys[1:], data, queryKeys[0], value)
		}

		if err != nil {
			return fmt.Errorf("index value of \"%s\", \"%v\" is not a number", queryKeys[0], index)
		}

		return putRecursive(&dataA[index], queryKeys[1:], data, queryKeys[0], value)
	} else if dataM, ok := (*data).(map[string]interface{}); ok && queryKeys[0][0] != '[' {
		if isLastKey {
			dataM[queryKeys[0]] = value
			return nil
		}

		for k, v := range dataM {
			if k == queryKeys[0] {
				return putRecursive(&v, queryKeys[1:], data, queryKeys[0], value)
			}
		}

		// If key not exist and the target is nested, create a new item
		if len(queryKeys) > 1 {
			var newValue interface{}
			if queryKeys[1][0] == '[' {
				newValue = []interface{}{}
			} else {
				newValue = map[string]interface{}{}
			}

			dataM[queryKeys[0]] = newValue
			return putRecursive(&newValue, queryKeys[1:], data, queryKeys[0], value)
		}
	}

	return fmt.Errorf("key \"%s\" not found", queryKeys[0])
}
