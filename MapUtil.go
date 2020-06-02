package utility

import (
	"fmt"
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

var NotFoundErr = fmt.Errorf("not found the value")

func GetInterface(mapObject interface{}, key interface{}) (interface{}, error) {
	ret := interface{}(nil)
	retErr := error(nil)

	keyValue := reflect.ValueOf(key)
	mapObjValue := reflect.ValueOf(mapObject)
	keyValues := mapObjValue.MapKeys()

	if len(keyValues) > 0 {
		if keyValue.Type() == keyValues[0].Type() {
			tempRet := mapObjValue.MapIndex(keyValue)
			if tempRet.IsValid() {
				ret = tempRet.Interface()
			} else {
				retErr = NotFoundErr
			}
		} else {
			retErr = fmt.Errorf("not match key type: %v vs %v", keyValue.Type(), keyValues[0].Type())
		}
	} else {
		retErr = NotFoundErr
	}

	return ret, retErr
}

func GetNumber(mapObject interface{}, key interface{}) (float64, error) {
	ret := float64(0)
	tempRet := interface{}(nil)
	retErr := error(nil)

	if tempRet, retErr = GetInterface(mapObject, key); retErr == nil {
		if tempRet != nil {
			if tempValue, assertionOK := tempRet.(float64); assertionOK {
				ret = tempValue
			} else if tempValue, assertionOK := tempRet.(float32); assertionOK {
				ret = float64(tempValue)
			} else if tempValue, assertionOK := tempRet.(int64); assertionOK {
				ret = float64(tempValue)
			} else if tempValue, assertionOK := tempRet.(int32); assertionOK {
				ret = float64(tempValue)
			} else if tempValue, assertionOK := tempRet.(int16); assertionOK {
				ret = float64(tempValue)
			} else if tempValue, assertionOK := tempRet.(int8); assertionOK {
				ret = float64(tempValue)
			} else if tempValue, assertionOK := tempRet.(int); assertionOK {
				ret = float64(tempValue)
			} else if tempValue, assertionOK := tempRet.(uint64); assertionOK {
				ret = float64(tempValue)
			} else if tempValue, assertionOK := tempRet.(uint32); assertionOK {
				ret = float64(tempValue)
			} else if tempValue, assertionOK := tempRet.(uint16); assertionOK {
				ret = float64(tempValue)
			} else if tempValue, assertionOK := tempRet.(uint8); assertionOK {
				ret = float64(tempValue)
			} else if tempValue, assertionOK := tempRet.(uint); assertionOK {
				ret = float64(tempValue)
			} else {
				retErr = fmt.Errorf("value is not Number type: %v", reflect.TypeOf(tempRet))
			}
		}
	}

	return ret, retErr
}

func GetString(mapObject interface{}, key interface{}) (string, error) {
	ret := ``
	tempRet := interface{}(nil)
	retErr := error(nil)

	if tempRet, retErr = GetInterface(mapObject, key); retErr == nil {
		if tempValue, assertionOK := tempRet.(string); assertionOK {
			ret = tempValue
		} else {
			retErr = fmt.Errorf("value is not string type: %v", reflect.TypeOf(tempRet))
		}
	}

	return ret, retErr
}

func IsNotFound(err error) bool {
	return err == NotFoundErr
}
