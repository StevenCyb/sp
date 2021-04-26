package utils

import (
	"fmt"
	"sp/model"
)

// YamlDifferencePrinter create a printable difference Yaml string
func YamlDifferencePrinter(differenceEntry *model.DifferenceEntry, result *string, indent int) {
	isArrayChild := differenceEntry.Parent != nil && differenceEntry.Parent.DifferenceEntryType == model.DET_ARRAY
	prefix := createPrefix(differenceEntry.DifferenceEntryStatus, indent)

	switch differenceEntry.DifferenceEntryType {
	case model.DET_ROOT:
		YamlDifferencePrinter(differenceEntry.ChildEntries[0], result, indent+1)

	case model.DET_OBJECT:
		if isArrayChild {
			*result += fmt.Sprintf("%s-\n", prefix)
		}
		if differenceEntry.Key != "" {
			*result += fmt.Sprintf("%s%s:\n", prefix, differenceEntry.Key)
		}

		for _, child := range differenceEntry.ChildEntries {
			YamlDifferencePrinter(child, result, indent+1)
		}

	case model.DET_ARRAY:
		if differenceEntry.Key != "" {
			*result += fmt.Sprintf("%s%s:\n", prefix, differenceEntry.Key)
		}

		for _, child := range differenceEntry.ChildEntries {
			YamlDifferencePrinter(child, result, indent+1)
		}

	case model.DET_KEY_VALUE:
		*result += prefix

		if isArrayChild {
			*result += "- "
		}

		if differenceEntry.Key != "" {
			*result += fmt.Sprintf("%s: ", differenceEntry.Key)
		}

		if fmt.Sprintf("%T", differenceEntry.Value) == "string" {
			*result += fmt.Sprintf("\"%v\"", differenceEntry.Value)
		} else {
			*result += fmt.Sprintf("%v", differenceEntry.Value)
		}

		*result += "\n"
	}
}
