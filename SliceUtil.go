package utility

import "reflect"

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
