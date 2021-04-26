package utils

import (
	"encoding/json"

	"github.com/BurntSushi/toml"
	"sigs.k8s.io/yaml"
)

func StringIsValidStructure(data string, dataType string) bool {
	if DataTypeIsJSON(dataType) {
		return StringIsValidJson(data)
	} else if DataTypeIsToml(dataType) {
		return StringIsValidToml(data)
	} else if DataTypeIsYaml(dataType) {
		return StringIsValidYaml(data)
	}

	return false
}
func ByteIsValidStructure(data []byte, dataType string) bool {
	if DataTypeIsJSON(dataType) {
		return ByteIsValidJson(data)
	} else if DataTypeIsToml(dataType) {
		return ByteIsValidToml(data)
	} else if DataTypeIsYaml(dataType) {
		return ByteIsValidYaml(data)
	}

	return false
}

// ByteIsValidJson checks if byte array contains valid JSON data
func ByteIsValidJson(data []byte) bool {
	return StringIsValidJson(string(data))
}

// StringIsValidJson checks if string contains valid JSON data
func StringIsValidJson(data string) bool {
	var tmp interface{}
	return json.Unmarshal([]byte(data), &tmp) == nil
}

// ByteIsValidToml checks if byte array contains valid toml data
func ByteIsValidToml(data []byte) bool {
	return StringIsValidToml(string(data))
}

// StringIsValidToml checks if string contains valid toml data
func StringIsValidToml(data string) bool {
	var tmp interface{}
	return toml.Unmarshal([]byte(data), &tmp) == nil
}

// ByteIsValidYaml checks if byte array contains valid yaml data
func ByteIsValidYaml(data []byte) bool {
	return StringIsValidYaml(string(data))
}

// StringIsValidYaml checks if string contains valid yaml data
func StringIsValidYaml(data string) bool {
	var tmp interface{}
	return yaml.Unmarshal([]byte(data), &tmp) == nil
}
