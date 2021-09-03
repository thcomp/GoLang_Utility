package utility

import (
	"reflect"
	"time"
)

type InterfaceHelper struct {
	value      reflect.Value
	kind       reflect.Kind
	ptrToValue reflect.Value
	ptrToKind  reflect.Kind
}

func NewInterfaceHelper(valueInterface interface{}) *InterfaceHelper {
	ret := InterfaceHelper{}

	if value, assertionOK := valueInterface.(reflect.Value); assertionOK {
		ret.value = value
		ret.kind = ret.value.Kind()
	} else {
		ret.value = reflect.ValueOf(valueInterface)
		ret.kind = ret.value.Kind()
	}

	if ret.kind == reflect.Ptr {
		ret.ptrToValue = reflect.Indirect(ret.value)
		ret.ptrToKind = ret.ptrToValue.Kind()
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
	case reflect.Ptr:
		switch helper.ptrToKind {
		case reflect.String:
			matched = true
			ret = helper.ptrToValue.String()
			break
		}
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
	case reflect.Ptr:
		switch helper.ptrToKind {
		case reflect.Bool:
			matched = true
			ret = helper.ptrToValue.Bool()
			break
		}
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
	case reflect.Ptr:
		switch helper.ptrToKind {
		case reflect.Float32, reflect.Float64:
			matched = true
			ret = helper.ptrToValue.Float()
			break
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			matched = true
			ret = float64(helper.ptrToValue.Int())
			break
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			matched = true
			ret = float64(helper.ptrToValue.Uint())
			break
		}
		break
	}

	return ret, matched
}

func (helper *InterfaceHelper) GetTime() (time.Time, bool) {
	ret := time.Time{}
	matched := false

	switch helper.kind {
	case reflect.Struct:
		if tempRet, assertionOK := helper.value.Interface().(time.Time); assertionOK {
			ret = tempRet
			matched = true
		}
		break
	case reflect.Ptr:
		switch helper.ptrToKind {
		case reflect.Struct:
			if tempRet, assertionOK := helper.ptrToValue.Interface().(time.Time); assertionOK {
				ret = tempRet
				matched = true
			}
			break
		}
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
			tempInterface := helper.value.Index(i).Interface()
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
			tempInterface := helper.value.Index(i).Interface()
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
			tempInterface := helper.value.Index(i).Interface()
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

func (helper *InterfaceHelper) GetTimeArray() ([]time.Time, bool) {
	ret := []time.Time(nil)
	matched := false

	switch helper.kind {
	case reflect.Array, reflect.Slice:
		matched = true
		for i := 0; i < helper.value.Len(); i++ {
			tempInterface := helper.value.Index(i1).Interface()
			tempHelper := NewInterfaceHelper(tempInterface)
			if tempTime, matched := tempHelper.GetTime(); matched {
				if ret == nil {
					ret = []time.Time{}
				}

				ret = append(ret, tempTime)
			}
		}
		break
	default:
		if helper.IsTime() {
			matched = true
			tempValue, _ := helper.GetTime()
			ret = []time.Time{tempValue}
		}
		break
	}

	return ret, matched
}

func (helper *InterfaceHelper) GetKind() reflect.Kind {
	ret := helper.kind

	if ret == reflect.Ptr {
		ret = helper.ptrToKind
	}

	return ret
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
	case reflect.Ptr:
		switch helper.ptrToKind {
		case reflect.Float32, reflect.Float64:
			ret = true
			break
		}
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
	case reflect.Ptr:
		switch helper.ptrToKind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			ret = true
			break
		}
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
	case reflect.Ptr:
		switch helper.ptrToKind {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			ret = true
			break
		}
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
	case reflect.Ptr:
		switch helper.ptrToKind {
		case reflect.String:
			ret = true
			break
		}
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
	case reflect.Ptr:
		switch helper.ptrToKind {
		case reflect.Bool:
			ret = true
			break
		}
		break
	}

	return ret
}

func (helper *InterfaceHelper) IsTime() bool {
	ret := false

	switch helper.kind {
	case reflect.Struct:
		_, ret = helper.value.Interface().(time.Time)
		break
	case reflect.Ptr:
		switch helper.ptrToKind {
		case reflect.Struct:
			_, ret = helper.ptrToValue.Interface().(time.Time)
			break
		}
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
	case reflect.Ptr:
		switch helper.ptrToKind {
		case reflect.Array, reflect.Slice:
			ret = true
			break
		}
		break
	}

	return ret
}

func (helper *InterfaceHelper) Set(from interface{}) bool {
	ret := true

	if helper.kind == reflect.Ptr {
		fromInfHelper := NewInterfaceHelper(from)

		if helper.IsBool() && fromInfHelper.IsBool() {
			value, _ := fromInfHelper.GetBool()
			helper.ptrToValue.SetBool(value)
		} else if helper.IsString() && fromInfHelper.IsString() {
			value, _ := fromInfHelper.GetString()
			helper.ptrToValue.SetString(value)
		} else if helper.IsInt() && fromInfHelper.IsInt() {
			value, _ := fromInfHelper.GetNumber()
			helper.ptrToValue.SetInt(int64(value))
		} else if helper.IsUint() && fromInfHelper.IsUint() {
			value, _ := fromInfHelper.GetNumber()
			helper.ptrToValue.SetUint(uint64(value))
		} else if helper.IsFloat() && fromInfHelper.IsFloat() {
			value, _ := fromInfHelper.GetNumber()
			helper.ptrToValue.SetFloat(value)
		} else if helper.IsTime() && fromInfHelper.IsTime() {
			value, _ := fromInfHelper.GetTime()
			helper.ptrToValue.Set(reflect.ValueOf(value))
		} else if helper.IsArrayOrSlice() && fromInfHelper.IsArrayOrSlice() {
			ret = false
		} else {
			ret = false
		}
	} else if helper.value.CanSet() {
		fromInfHelper := NewInterfaceHelper(from)

		if helper.IsBool() && fromInfHelper.IsBool() {
			value, _ := fromInfHelper.GetBool()
			helper.value.SetBool(value)
		} else if helper.IsString() && fromInfHelper.IsString() {
			value, _ := fromInfHelper.GetString()
			helper.value.SetString(value)
		} else if helper.IsInt() && fromInfHelper.IsInt() {
			value, _ := fromInfHelper.GetNumber()
			helper.value.SetInt(int64(value))
		} else if helper.IsUint() && fromInfHelper.IsUint() {
			value, _ := fromInfHelper.GetNumber()
			helper.value.SetUint(uint64(value))
		} else if helper.IsFloat() && fromInfHelper.IsFloat() {
			value, _ := fromInfHelper.GetNumber()
			helper.value.SetFloat(value)
		} else if helper.IsTime() && fromInfHelper.IsTime() {
			value, _ := fromInfHelper.GetNumber()
			helper.value.Set(reflect.ValueOf(value))
		} else if helper.IsArrayOrSlice() && fromInfHelper.IsArrayOrSlice() {
			ret = false
		} else {
			ret = false
		}
	} else {
		ret = false
	}

	return ret
}
