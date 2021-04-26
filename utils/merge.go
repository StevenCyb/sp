package utils

// Merge two structured data sets
func Merge(dataA *interface{}, dataB *interface{}, arrayAppend bool) {
	mergeRecursive(dataA, dataB, arrayAppend, nil, "")
}

// mergeRecursive perform a deep merge
func mergeRecursive(dataA *interface{}, dataB *interface{}, arrayAppend bool, previousData *interface{}, previousKey string) {
	aArray, aIsArray := (*dataA).([]interface{})
	bArray, bIsArray := (*dataB).([]interface{})
	aMap, aIsMap := (*dataA).(map[string]interface{})
	bMap, bIsMap := (*dataB).(map[string]interface{})

	if aIsArray && bIsArray {
		if arrayAppend {
			SetArrayInDataTo(dataB, "[+]", previousData, previousKey, append(aArray, bArray...))
		} else {
			SetArrayInDataTo(dataB, "[+]", previousData, previousKey, aArray)
		}
		return
	} else if aIsMap && bIsMap {
		for aKey, aValue := range aMap {
			if _, exists := bMap[aKey]; exists {
				tmp := bMap[aKey]
				mergeRecursive(&aValue, &tmp, arrayAppend, dataB, aKey)
			} else {
				bMap[aKey] = aValue
			}
		}
	} else {
		if previousData != nil {
			pMap, pIsMap := (*previousData).(map[string]interface{})

			if pIsMap {
				pMap[previousKey] = *dataA
				return
			}
		}

		*dataB = *dataA
	}
}
