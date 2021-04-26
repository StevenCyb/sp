package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/BurntSushi/toml"
	"sigs.k8s.io/yaml"
)

// NewDataFromFile read data from file and create an interface pointer
func NewDataFromFile(filePath string, dataType string) (*interface{}, error) {
	rawData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	if dataType == "extension" {
		dataType = strings.ToLower(path.Ext(filePath))
		if len(dataType) > 2 {
			dataType = dataType[1:]
		}
	}

	return NewDataFromBytes(rawData, dataType)
}

// NewDataFromString create an interface pointer from string data
func NewDataFromString(data string, dataType string) (*interface{}, error) {
	return NewDataFromBytes([]byte(data), dataType)
}

// NewDataFromBytes create an interface pointer from byte data
func NewDataFromBytes(rawData []byte, dataType string) (*interface{}, error) {
	var err error
	var data interface{}

	if DataTypeIsJSON(dataType) {
		if !ByteIsValidJson(rawData) {
			return nil, fmt.Errorf("data has invalid JSON structure")
		}
		err = json.Unmarshal(rawData, &data)

	} else if DataTypeIsToml(dataType) {
		if !ByteIsValidToml(rawData) {
			return nil, fmt.Errorf("data has invalid Toml structure")
		}
		err = toml.Unmarshal(rawData, &data)

	} else if DataTypeIsYaml(dataType) {
		if !ByteIsValidYaml(rawData) {
			return nil, fmt.Errorf("data has invalid Yaml structure")
		}
		err = yaml.Unmarshal(rawData, &data)

	} else {
		return nil, fmt.Errorf("data not match any known extension, use %s", dataType)
	}

	if err != nil {
		return nil, err
	}
	return &data, nil
}
