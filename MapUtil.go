package utility

import (
	"reflect"
)

func MapKeys(mapObj interface{}) interface{} {
	var ret interface{} = nil
	mapObjValue := reflect.ValueOf(mapObj)
	keyValues := mapObjValue.MapKeys()

	if len(keyValues) > 0 {
		keyType := keyValues[0].Type()
		sliceType := reflect.SliceOf(keyType)
		sliceValue := reflect.MakeSlice(sliceType, 0, 0)

		for _, key := range keyValues {
			sliceValue = reflect.Append(sliceValue, key)
		}

		ret = sliceValue.Interface()
	}

	return ret
}
