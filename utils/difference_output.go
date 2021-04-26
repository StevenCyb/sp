package utils

import (
	"fmt"
	"sp/model"
	"strings"
)

// CreateDifferenceOutput use the root DifferenceEntry object to create a formated result string
func CreateDifferenceOutput(differenceEntry *model.DifferenceEntry, dataType string) (string, error) {
	result := ""

	if DataTypeIsJSON(dataType) {
		JSONDifferencePrinter(differenceEntry, &result, 0)
	} else if DataTypeIsToml(dataType) {
		TomlDifferencePrinter(differenceEntry, &result, 0)
	} else if DataTypeIsYaml(dataType) {
		YamlDifferencePrinter(differenceEntry, &result, 0)
	}

	if len(result) > 0 {
		return result, nil
	}

	return "", fmt.Errorf("unsupported output format \"%s\"", dataType)
}

// createPrefix create a prefix by given difference type and indent
func createPrefix(des model.DifferenceEntryStatus, indent int) string {
	prefix := strings.Repeat("  ", indent+1)

	if des == model.DES_NEUTRAL {
		prefix = "   " + prefix
	} else if des == model.DES_ADD {
		prefix = "+++" + prefix
	} else if des == model.DES_REMOVE {
		prefix = "---" + prefix
	}

	return prefix
}
