package utils

import (
	"fmt"
	"sp/model"
	"strings"
)

// TomlDifferencePrinter create a printable difference Toml string
func TomlDifferencePrinter(differenceEntry *model.DifferenceEntry, result *string, indent int) {
	isArrayChild := differenceEntry.Parent != nil && differenceEntry.Parent.DifferenceEntryType == model.DET_ARRAY
	prefix := createPrefix(differenceEntry.DifferenceEntryStatus, 0)

	switch differenceEntry.DifferenceEntryType {
	case model.DET_ROOT:
		TomlDifferencePrinter(differenceEntry.ChildEntries[0], result, indent+1)

	case model.DET_OBJECT:
		objectKeySlice := []string{}
		current := differenceEntry

		for current != nil {
			if current.Key != "" {
				objectKeySlice = append([]string{current.Key}, objectKeySlice...)
			}
			current = current.Parent
		}

		hasKeyValueChild := false
		for _, child := range differenceEntry.ChildEntries {
			if child.DifferenceEntryType == model.DET_KEY_VALUE {
				hasKeyValueChild = true
				break
			}
		}

		if len(objectKeySlice) > 0 && !isArrayChild && hasKeyValueChild {
			if *result != "" && !strings.HasSuffix(*result, "\n\n") {
				*result += "\n"
			}

			*result += fmt.Sprintf("%s[%s]\n", prefix, strings.Join(objectKeySlice[:], "."))
		}

		separateChildNodes := []*model.DifferenceEntry{}
		for _, child := range differenceEntry.ChildEntries {
			if child.DifferenceEntryType != model.DET_OBJECT &&
				!(child.DifferenceEntryType == model.DET_ARRAY &&
					(len(child.ChildEntries) > 0 ||
						child.ChildEntries[0].DifferenceEntryType != model.DET_KEY_VALUE)) {
				TomlDifferencePrinter(child, result, indent+1)
			} else {
				separateChildNodes = append(separateChildNodes, child)
			}
		}

		if *result != "" && !strings.HasSuffix(*result, "\n\n") {
			*result += "\n"
		}

		for _, child := range separateChildNodes {
			TomlDifferencePrinter(child, result, indent+1)
		}

	case model.DET_ARRAY:
		if *result != "" && !strings.HasSuffix(*result, "\n\n") {
			*result += "\n"
		}

		if len(differenceEntry.ChildEntries) == 0 {
			*result += fmt.Sprintf("%s%s = []\n", prefix, differenceEntry.Key)
		} else if differenceEntry.ChildEntries[0].DifferenceEntryType == model.DET_KEY_VALUE {
			*result += fmt.Sprintf("%s%s = [\n", prefix, differenceEntry.Key)

			for _, child := range differenceEntry.ChildEntries {
				TomlDifferencePrinter(child, result, indent+1)
			}

			*result += fmt.Sprintf("%s]\n", prefix)
		} else {
			objectKeySlice := []string{}
			current := differenceEntry

			for current != nil {
				if current.Key != "" {
					objectKeySlice = append([]string{current.Key}, objectKeySlice...)
				}
				current = current.Parent
			}

			for _, child := range differenceEntry.ChildEntries {
				if *result != "" && !strings.HasSuffix(*result, "\n\n") {
					*result += "\n"
				}

				*result += fmt.Sprintf("%s[[%s]]\n", prefix, strings.Join(objectKeySlice[:], "."))
				TomlDifferencePrinter(child, result, indent+1)

				if *result != "" && !strings.HasSuffix(*result, "\n\n") {
					*result += "\n"
				}
			}
		}

		if *result != "" && !strings.HasSuffix(*result, "\n\n") {
			*result += "\n"
		}

	case model.DET_KEY_VALUE:
		*result += prefix

		if isArrayChild {
			*result += "  "
		}

		if differenceEntry.Key != "" {
			*result += fmt.Sprintf("%s = ", differenceEntry.Key)
		}

		if fmt.Sprintf("%T", differenceEntry.Value) == "string" {
			*result += fmt.Sprintf("\"%v\"", differenceEntry.Value)
		} else {
			*result += fmt.Sprintf("%v", differenceEntry.Value)
		}

		if isArrayChild {
			*result += ","
		}

		*result += "\n"
	}
}
