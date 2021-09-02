package utility

import (
	"fmt"
	"reflect"
	"strconv"
)

func RemoveSliceItem(slice interface{}, position int) (interface{}, bool) {
	retSlice := interface{}(nil)
	ret := false

	if slice != nil {
		sliceValue := reflect.ValueOf(slice)
		if sliceValue.Kind() == reflect.Slice {
			originalSliceLen := sliceValue.Len()
			if position >= 0 && originalSliceLen > position {
				tempSlice := reflect.MakeSlice(reflect.TypeOf(slice), 0, 0)
				if position > 0 {
					tempSlice = reflect.AppendSlice(tempSlice, sliceValue.Slice(0, position))
				}
				if position < originalSliceLen-1 {
					tempSlice = reflect.AppendSlice(tempSlice, sliceValue.Slice(position+1, originalSliceLen))
				}
				retSlice = tempSlice.Interface()
				ret = true
			}
		}
	}

	return retSlice, ret
}

func ToIntSlice(fromInterface interface{}) (ret []int, retErr error) {
	ret = []int{}

	if fromInterface == nil {
		// no work
	} else {
		infHelper := NewInterfaceHelper(fromInterface)

		if infHelper.IsArrayOrSlice() {
			if float64Slice, matched := infHelper.GetNumberArray(); matched {
				for _, float64Value := range float64Slice {
					ret = append(ret, int(float64Value))
				}
			} else if stringSlice, matched := infHelper.GetStringArray(); matched {
				for _, stringValue := range stringSlice {
					if float64Value, parseErr := strconv.ParseFloat(stringValue, 64); parseErr == nil {
						ret = append(ret, int(float64Value))
					} else if intValue, parseErr := strconv.ParseInt(stringValue, 10, 64); parseErr == nil {
						ret = append(ret, int(intValue))
					} else if intHexValue, parseErr := strconv.ParseInt(stringValue, 16, 64); parseErr == nil {
						ret = append(ret, int(intHexValue))
					} else {
						LogfE("not support string: %s", stringValue)
					}
				}
			} else {
				retErr = fmt.Errorf("not support type: %v", fromInterface)
			}
		} else if infHelper.IsNumber() {
			value, _ := infHelper.GetNumber()
			ret = append(ret, int(value))
		} else if infHelper.IsString() {
			stringValue, _ := infHelper.GetString()
			if float64Value, parseErr := strconv.ParseFloat(stringValue, 64); parseErr == nil {
				ret = append(ret, int(float64Value))
			} else if intValue, parseErr := strconv.ParseInt(stringValue, 10, 64); parseErr == nil {
				ret = append(ret, int(intValue))
			} else if intHexValue, parseErr := strconv.ParseInt(stringValue, 16, 64); parseErr == nil {
				ret = append(ret, int(intHexValue))
			} else {
				retErr = fmt.Errorf("not support string: %s", stringValue)
			}
		} else {
			retErr = fmt.Errorf("not support type: %v", fromInterface)
		}
	}

	return
}
