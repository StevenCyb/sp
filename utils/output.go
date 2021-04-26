package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/BurntSushi/toml"
	"sigs.k8s.io/yaml"
)

func DataToFile(data *interface{}, dataType string, filePath string, pretty bool) error {
	if dataType == "extension" {
		dataType = strings.ToLower(path.Ext(filePath))
		if len(dataType) > 2 {
			dataType = dataType[1:]
		}
	}

	rawData, err := createRawOutputData(data, dataType, pretty)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, rawData, os.ModePerm)
}

func DataToStd(data *interface{}, dataType string, pretty bool) error {
	rawData, err := createRawOutputData(data, dataType, pretty)
	if err != nil {
		return err
	}

	fmt.Println(string(rawData))

	return nil
}

func createRawOutputData(data *interface{}, dataType string, pretty bool) ([]byte, error) {
	var err error
	rawData := []byte{}

	if DataTypeIsJSON(dataType) {
		if pretty {
			rawData, err = json.MarshalIndent(*data, "", "  ")
		} else {
			rawData, err = json.Marshal(*data)
		}
	} else if DataTypeIsToml(dataType) {
		buf := new(bytes.Buffer)
		err = toml.NewEncoder(buf).Encode(*data)
		if err == nil {
			rawData = buf.Bytes()
		}
	} else if DataTypeIsYaml(dataType) {
		rawData, err = yaml.Marshal(*data)
	}

	if err == nil {
		return rawData, nil
	}

	return nil, fmt.Errorf("unsupported output data type %s", dataType)
}
