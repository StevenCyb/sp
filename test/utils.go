package test

import (
	"fmt"
	"testing"
)

func AssertEqual(t *testing.T, actualValue interface{}, expectedValue interface{}) {
	t.Helper()

	if fmt.Sprintf("%T", expectedValue) == "map[string]interface {}" {
		actualValueMap := actualValue.(map[string]interface{})
		expectedValueMap := expectedValue.(map[string]interface{})
		for key, actualValue := range actualValueMap {
			if expectedValue, ok := expectedValueMap[key]; !ok {
				AssertEqual(t, actualValue, expectedValue)
			}
		}
	} else if fmt.Sprintf("%T", expectedValue) == "[]interface {}" {
		if fmt.Sprintf("%+v", actualValue) != fmt.Sprintf("%+v", expectedValue) {
			t.Errorf("Expected %+v, got %+v", expectedValue, actualValue)
		}

	} else if expectedValue != actualValue {
		t.Errorf("Expected %+v, got %+v", expectedValue, actualValue)
	}
}

func AssertNotEqual(t *testing.T, actualValue interface{}, expectedValue interface{}) {
	t.Helper()

	if fmt.Sprintf("%T", expectedValue) == "map[string]interface {}" {
		actualValueMap := actualValue.(map[string]interface{})
		expectedValueMap := expectedValue.(map[string]interface{})
		for key, actualValue := range actualValueMap {
			if expectedValue, ok := expectedValueMap[key]; !ok {
				AssertEqual(t, actualValue, expectedValue)
			}
		}
		return
	}

	if expectedValue == actualValue {
		t.Errorf("Expected %+v to be not equal %+v", expectedValue, actualValue)
	}
}
