package utils

import "reflect"

func Equal(dataA *interface{}, dataB *interface{}) bool {
	return reflect.DeepEqual(*dataA, *dataB)
}
