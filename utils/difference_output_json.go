package utils

import (
	"fmt"
	"sp/model"
)

// JSONDifferencePrinter create a printable difference JSON string
func JSONDifferencePrinter(differenceEntry *model.DifferenceEntry, result *string, indent int) {
	isArrayChild := differenceEntry.Parent != nil && differenceEntry.Parent.DifferenceEntryType == model.DET_ARRAY
	prefix := createPrefix(differenceEntry.DifferenceEntryStatus, indent)

	switch differenceEntry.DifferenceEntryType {
	case model.DET_ROOT:
		JSONDifferencePrinter(differenceEntry.ChildEntries[0], result, indent+1)

	case model.DET_OBJECT:
		if isArrayChild || differenceEntry.Key == "" {
			*result += fmt.Sprintf("%s{\n", prefix)
		} else {
			*result += fmt.Sprintf("%s\"%s\": {\n", prefix, differenceEntry.Key)
		}

		for _, child := range differenceEntry.ChildEntries {
			JSONDifferencePrinter(child, result, indent+1)
		}

		if isArrayChild && !differenceEntry.ArrayHelpMarker {
			*result += fmt.Sprintf("%s},\n", prefix)
		} else {
			*result += fmt.Sprintf("%s}\n", prefix)
		}

	case model.DET_ARRAY:
		if isArrayChild && !differenceEntry.ArrayHelpMarker {
			*result += fmt.Sprintf("%s[\n", prefix)
		} else {
			*result += fmt.Sprintf("%s\"%s\": [\n", prefix, differenceEntry.Key)
		}

		var lastAddChild *model.DifferenceEntry
		var lastRemoveChild *model.DifferenceEntry
		lastIsNeutral := false
		for _, child := range differenceEntry.ChildEntries {
			lastIsNeutral = child.DifferenceEntryStatus == model.DES_NEUTRAL
			if child.DifferenceEntryStatus == model.DES_ADD {
				lastAddChild = child
			}
			if child.DifferenceEntryStatus == model.DES_REMOVE {
				lastRemoveChild = child
			}
		}
		if !lastIsNeutral {
			if lastAddChild != nil {
				lastAddChild.ArrayHelpMarker = true
			}
			if lastRemoveChild != nil {
				lastRemoveChild.ArrayHelpMarker = true
			}
		}

		for _, child := range differenceEntry.ChildEntries {
			JSONDifferencePrinter(child, result, indent+1)
		}
		if isArrayChild && !differenceEntry.ArrayHelpMarker {
			*result += fmt.Sprintf("%s],\n", prefix)
		} else {
			*result += fmt.Sprintf("%s]\n", prefix)
		}

	case model.DET_KEY_VALUE:
		*result += prefix

		if differenceEntry.Key != "" {
			*result += fmt.Sprintf("\"%s\":", differenceEntry.Key)
		}

		if fmt.Sprintf("%T", differenceEntry.Value) == "string" {
			*result += fmt.Sprintf("\"%v\"", differenceEntry.Value)
		} else {
			*result += fmt.Sprintf("%v", differenceEntry.Value)
		}

		if isArrayChild && !differenceEntry.ArrayHelpMarker {
			*result += ",\n"
		} else {
			*result += "\n"
		}
	}
}
