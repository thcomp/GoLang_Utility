package utility

import (
	"reflect"
)

type InterfaceHelper struct {
	value reflect.Value
	kind  reflect.Kind
}

func NewInterfaceHelper(valueInterface interface{}) *InterfaceHelper {
	ret := InterfaceHelper{
		value: reflect.ValueOf(valueInterface),
		kind:  reflect.ValueOf(valueInterface).Kind(),
	}

	return &ret
}

func (helper *InterfaceHelper) GetString() (string, bool) {
	ret := ""
	matched := false

	switch helper.kind {
	case reflect.String:
		matched = true
		ret = helper.value.String()
		break
	}

	return ret, matched
}

func (helper *InterfaceHelper) GetBool() (bool, bool) {
	ret := false
	matched := false

	switch helper.kind {
	case reflect.Bool:
		matched = true
		ret = helper.value.Bool()
		break
	}

	return ret, matched
}

func (helper *InterfaceHelper) GetNumber() (float64, bool) {
	ret := float64(0)
	matched := false

	switch helper.kind {
	case reflect.Float32, reflect.Float64:
		matched = true
		ret = helper.value.Float()
		break
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		matched = true
		ret = float64(helper.value.Int())
		break
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		matched = true
		ret = float64(helper.value.Uint())
		break
	}

	return ret, matched
}

func (helper *InterfaceHelper) GetNumberArray() ([]float64, bool) {
	ret := []float64(nil)
	matched := false

	switch helper.kind {
	case reflect.Array, reflect.Slice:
		matched = true
		for i := 0; i < helper.value.Len(); i++ {
			tempInterface := helper.value.Slice(i, i+1).Interface()
			tempHelper := NewInterfaceHelper(tempInterface)
			if tempFloat, matched := tempHelper.GetNumber(); matched {
				if ret == nil {
					ret = []float64{}
				}

				ret = append(ret, tempFloat)
			}
		}
		break
	default:
		if helper.IsNumber() {
			matched = true
			tempValue, _ := helper.GetNumber()
			ret = []float64{tempValue}
		}
		break
	}

	return ret, matched
}

func (helper *InterfaceHelper) GetStringArray() ([]string, bool) {
	ret := []string(nil)
	matched := false

	switch helper.kind {
	case reflect.Array, reflect.Slice:
		matched = true
		for i := 0; i < helper.value.Len(); i++ {
			tempInterface := helper.value.Slice(i, i+1).Interface()
			tempHelper := NewInterfaceHelper(tempInterface)
			if tempString, matched := tempHelper.GetString(); matched {
				if ret == nil {
					ret = []string{}
				}

				ret = append(ret, tempString)
			}
		}
		break
	default:
		if helper.IsString() {
			matched = true
			tempValue, _ := helper.GetString()
			ret = []string{tempValue}
		}
		break
	}

	return ret, matched
}

func (helper *InterfaceHelper) GetBoolArray() ([]bool, bool) {
	ret := []bool(nil)
	matched := false

	switch helper.kind {
	case reflect.Array, reflect.Slice:
		matched = true
		for i := 0; i < helper.value.Len(); i++ {
			tempInterface := helper.value.Slice(i, i+1).Interface()
			tempHelper := NewInterfaceHelper(tempInterface)
			if tempBool, matched := tempHelper.GetBool(); matched {
				if ret == nil {
					ret = []bool{}
				}

				ret = append(ret, tempBool)
			}
		}
		break
	default:
		if helper.IsBool() {
			matched = true
			tempValue, _ := helper.GetBool()
			ret = []bool{tempValue}
		}
		break
	}

	return ret, matched
}

func (helper *InterfaceHelper) GetKind() reflect.Kind {
	return helper.kind
}

func (helper *InterfaceHelper) IsNumber() bool {
	ret := false

	if helper.IsInt() || helper.IsUint() || helper.IsFloat() {
		ret = true
	}

	return ret
}

func (helper *InterfaceHelper) IsFloat() bool {
	ret := false

	switch helper.kind {
	case reflect.Float32, reflect.Float64:
		ret = true
		break
	}

	return ret
}

func (helper *InterfaceHelper) IsInt() bool {
	ret := false

	switch helper.kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ret = true
		break
	}

	return ret
}

func (helper *InterfaceHelper) IsUint() bool {
	ret := false

	switch helper.kind {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		ret = true
		break
	}

	return ret
}

func (helper *InterfaceHelper) IsString() bool {
	ret := false

	switch helper.kind {
	case reflect.String:
		ret = true
		break
	}

	return ret
}

func (helper *InterfaceHelper) IsBool() bool {
	ret := false

	switch helper.kind {
	case reflect.Bool:
		ret = true
		break
	}

	return ret
}

func (helper *InterfaceHelper) IsArrayOrSlice() bool {
	ret := false

	switch helper.kind {
	case reflect.Array, reflect.Slice:
		ret = true
		break
	}

	return ret
}
