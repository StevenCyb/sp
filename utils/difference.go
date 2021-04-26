package utils

import (
	"sp/model"
)

// This file is a horrible mess but does its job
// for now. Refactore it or use a better approach.

// Difference create a result string with the difference between two strcured data
func Difference(dataA *interface{}, dataB *interface{}, dataType string) (string, error) {
	differenceEntryRoot := &model.DifferenceEntry{
		ChildEntries:        []*model.DifferenceEntry{},
		DifferenceEntryType: model.DET_ROOT,
	}

	differenceRecursive(dataA, dataB, "", differenceEntryRoot)

	return CreateDifferenceOutput(differenceEntryRoot, dataType)
}

// differenceRecursive run a deep difference
func differenceRecursive(dataA *interface{}, dataB *interface{}, previousKey string, differenceEntry *model.DifferenceEntry) {
	aIsArray := false
	aArray := []interface{}{}
	aIsMap := false
	aMap := map[string]interface{}{}

	bIsArray := false
	bArray := []interface{}{}
	bIsMap := false
	bMap := map[string]interface{}{}

	if dataA != nil {
		aArray, aIsArray = (*dataA).([]interface{})
		aMap, aIsMap = (*dataA).(map[string]interface{})
	}

	if dataB != nil {
		bArray, bIsArray = (*dataB).([]interface{})
		bMap, bIsMap = (*dataB).(map[string]interface{})
	}

	if dataA != nil && dataB != nil && (aIsArray != bIsArray || aIsMap != bIsMap) {
		// Totally different strcutres
		differenceRecursive(dataA, nil, "", differenceEntry)
		differenceRecursive(nil, dataB, "", differenceEntry)

	} else if dataB == nil && aIsArray {
		// Structured data A has an array that are not in B.
		// So we mark the whole array as model.REMOVEd.
		nextDifferenceEntry := differenceEntry.AddChild(previousKey, "", model.DES_REMOVE, model.DET_ARRAY)

		for _, aValue := range aArray {
			differenceRecursive(&aValue, nil, "", nextDifferenceEntry)
		}

	} else if dataA == nil && bIsArray {
		// Structured data B has an array that are not in A.
		// So we mark the whole array as added.
		nextDifferenceEntry := differenceEntry.AddChild(previousKey, "", model.DES_ADD, model.DET_ARRAY)

		for _, bValue := range bArray {
			differenceRecursive(nil, &bValue, "", nextDifferenceEntry)
		}

	} else if aIsArray && bIsArray {
		// Both structured data has an array.
		// So we check all entries with the following rules:
		// - The items of A that are above the index of B are marked as model.REMOVEd (done in recursive deep difference by setting B as nil)
		// - The items of B that are above the index of A are marked as added (done in recursive deep difference by setting A as nil)
		// - Indices that are available on A and B are passed for deep difference check
		nextDifferenceEntry := differenceEntry.AddChild(previousKey, "", model.DES_NEUTRAL, model.DET_ARRAY)

		aLen := len(aArray)
		bLen := len(bArray)
		max := Max(aLen, bLen)

		for i := 0; i < max; i++ {
			if i >= aLen {
				differenceRecursive(nil, &bArray[i], "", nextDifferenceEntry)
			} else if i >= bLen {
				differenceRecursive(&aArray[i], nil, "", nextDifferenceEntry)
			} else {
				differenceRecursive(&aArray[i], &bArray[i], "", nextDifferenceEntry)
			}
		}

	} else if aIsMap && bIsMap {
		// A and B have a map with on the same layer or with the same name.
		// So we check for all items in A, if there are present in B.
		// If there are present, we pass it to a deep difference otherwise we
		// mark them as model.REMOVEd (done in recursive deep difference by setting B as nil).
		// Afterwards we search for keys that are in B but not in A.
		// All this items are marked as added.
		nextDifferenceEntry := differenceEntry.AddChild(previousKey, "", model.DES_NEUTRAL, model.DET_OBJECT)

		for aKey, aValue := range aMap {
			if _, exists := bMap[aKey]; exists {
				tmp := bMap[aKey]
				differenceRecursive(&aValue, &tmp, aKey, nextDifferenceEntry)
			} else {
				differenceRecursive(&aValue, nil, aKey, nextDifferenceEntry)
			}
		}
		for bKey, bValue := range bMap {
			if _, exists := aMap[bKey]; !exists {
				differenceRecursive(nil, &bValue, bKey, nextDifferenceEntry)
			}
		}

	} else if dataB == nil && aIsMap {
		// IF A has a map that not exist on B, we mark the whole map as model.REMOVEd.
		nextDifferenceEntry := differenceEntry.AddChild(previousKey, "", model.DES_REMOVE, model.DET_OBJECT)

		for aKey, aValue := range aMap {
			differenceRecursive(&aValue, nil, aKey, nextDifferenceEntry)
		}

	} else if dataA == nil && bIsMap {
		// IF B has a map that not exist on A, we mark the whole map as added.
		nextDifferenceEntry := differenceEntry.AddChild(previousKey, "", model.DES_ADD, model.DET_OBJECT)

		for bKey, bValue := range bMap {
			differenceRecursive(nil, &bValue, bKey, nextDifferenceEntry)
		}

	} else {
		// If A and B are not a map or array...
		// ... B not exists so we mark A as model.REMOVEd.
		// ... A not exists so we mark B as added.
		// ... A and B exist but are not equal mark A as model.REMOVEd and B as added.
		// ... A and B exist and are equal mark them as neutral.

		if dataB == nil {
			differenceEntry.AddChild(previousKey, *dataA, model.DES_REMOVE, model.DET_KEY_VALUE)
		} else if dataA == nil {
			differenceEntry.AddChild(previousKey, *dataB, model.DES_ADD, model.DET_KEY_VALUE)
		} else if *dataA != *dataB {
			differenceEntry.AddChild(previousKey, *dataA, model.DES_REMOVE, model.DET_KEY_VALUE)
			differenceEntry.AddChild(previousKey, *dataB, model.DES_ADD, model.DET_KEY_VALUE)
		} else {
			differenceEntry.AddChild(previousKey, *dataB, model.DES_NEUTRAL, model.DET_KEY_VALUE)
		}
	}
}
