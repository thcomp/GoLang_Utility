package utility

import (
	"reflect"
	"testing"
)

func Test_value(t *testing.T) {
	value := int8(8)
	refValue := reflect.ValueOf(&value)
	helper := NewInterfaceHelper(refValue)
	helper.Set(10)

	if !helper.IsNumber() {
		t.Fatalf("not match")
	} else if helper.IsBool() {
		t.Fatalf("not match")
	} else if helper.IsFloat() {
		t.Fatalf("not match")
	} else if !helper.IsInt() {
		t.Fatalf("not match")
	} else if helper.IsString() {
		t.Fatalf("not match")
	} else if helper.IsUint() {
		t.Fatalf("not match")
	}

	if valueFloat, assertionOK := helper.GetNumber(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if valueFloat != float64(value) {
			t.Fatalf("not match: %f vs %f", valueFloat, float64(value))
		}
	}

	if _, assertionOK := helper.GetBool(); assertionOK {
		t.Fatalf("not match")
	} else if _, assertionOK := helper.GetString(); assertionOK {
		t.Fatalf("not match")
	}

	if array, assertionOK := helper.GetNumberArray(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if len(array) != 1 {
			t.Fatalf("not match")
		} else if array[0] != float64(value) {
			t.Fatalf("not match")
		}
	}

	if _, assertionOK := helper.GetStringArray(); assertionOK {
		t.Fatalf("not match")
	}
}

func Test_field(t *testing.T) {
	structure := struct {
		IntField  int
		UintField uint
		StrField  string
		BoolField bool
	}{
		IntField:  -10,
		UintField: 10,
		StrField:  "100",
		BoolField: true,
	}
	refHelper := NewReflectHelper(&structure)
	{
		helper := NewInterfaceHelper(refHelper.GetValueByIndex(0))
		helper.Set(10)

		if structure.IntField != 10 {
			t.Fatalf("not matched: %d vs %d", structure.IntField, 10)
		}
	}

	{
		helper := NewInterfaceHelper(refHelper.GetValueByIndex(1))
		helper.Set(uint(100))

		if structure.UintField != 100 {
			t.Fatalf("not matched: %d vs %d", structure.UintField, 100)
		}
	}

	{
		helper := NewInterfaceHelper(refHelper.GetValueByIndex(2))
		helper.Set("1000")

		if structure.StrField != "1000" {
			t.Fatalf("not matched: %s vs 1000", structure.StrField)
		}
	}

	{
		helper := NewInterfaceHelper(refHelper.GetValueByIndex(3))
		helper.Set(false)

		if structure.BoolField != false {
			t.Fatalf("not matched: %v vs false", structure.BoolField)
		}
	}
}

func Test_ptr_int8(t *testing.T) {
	value := int8(8)
	helper := NewInterfaceHelper(&value)

	if !helper.IsNumber() {
		t.Fatalf("not match")
	} else if helper.IsBool() {
		t.Fatalf("not match")
	} else if helper.IsFloat() {
		t.Fatalf("not match")
	} else if !helper.IsInt() {
		t.Fatalf("not match")
	} else if helper.IsString() {
		t.Fatalf("not match")
	} else if helper.IsUint() {
		t.Fatalf("not match")
	}

	if valueFloat, assertionOK := helper.GetNumber(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if valueFloat != float64(value) {
			t.Fatalf("not match: %f vs %f", valueFloat, float64(value))
		}
	}

	if _, assertionOK := helper.GetBool(); assertionOK {
		t.Fatalf("not match")
	} else if _, assertionOK := helper.GetString(); assertionOK {
		t.Fatalf("not match")
	}

	if array, assertionOK := helper.GetNumberArray(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if len(array) != 1 {
			t.Fatalf("not match")
		} else if array[0] != float64(value) {
			t.Fatalf("not match")
		}
	}

	if _, assertionOK := helper.GetStringArray(); assertionOK {
		t.Fatalf("not match")
	}
}

func Test_ptr_uint8(t *testing.T) {
	value := uint8(8)
	helper := NewInterfaceHelper(&value)

	if !helper.IsNumber() {
		t.Fatalf("not match")
	} else if helper.IsBool() {
		t.Fatalf("not match")
	} else if helper.IsFloat() {
		t.Fatalf("not match")
	} else if helper.IsInt() {
		t.Fatalf("not match")
	} else if helper.IsString() {
		t.Fatalf("not match")
	} else if !helper.IsUint() {
		t.Fatalf("not match")
	}

	if valueFloat, assertionOK := helper.GetNumber(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if valueFloat != float64(value) {
			t.Fatalf("not match: %f vs %f", valueFloat, float64(value))
		}
	}

	if _, assertionOK := helper.GetBool(); assertionOK {
		t.Fatalf("not match")
	} else if _, assertionOK := helper.GetString(); assertionOK {
		t.Fatalf("not match")
	}

	if array, assertionOK := helper.GetNumberArray(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if len(array) != 1 {
			t.Fatalf("not match")
		} else if array[0] != float64(value) {
			t.Fatalf("not match")
		}
	}

	if _, assertionOK := helper.GetStringArray(); assertionOK {
		t.Fatalf("not match")
	}
}

func Test_ptr_string(t *testing.T) {
	value := "8"
	helper := NewInterfaceHelper(&value)

	if helper.IsNumber() {
		t.Fatalf("not match")
	} else if helper.IsBool() {
		t.Fatalf("not match")
	} else if helper.IsFloat() {
		t.Fatalf("not match")
	} else if helper.IsInt() {
		t.Fatalf("not match")
	} else if !helper.IsString() {
		t.Fatalf("not match")
	} else if helper.IsUint() {
		t.Fatalf("not match")
	}

	if _, assertionOK := helper.GetNumber(); assertionOK {
		t.Fatalf("not match")
	}

	if _, assertionOK := helper.GetBool(); assertionOK {
		t.Fatalf("not match")
	} else if _, assertionOK := helper.GetString(); !assertionOK {
		t.Fatalf("not match")
	}

	if _, assertionOK := helper.GetNumberArray(); assertionOK {
		t.Fatalf("not match")
	}

	if array, assertionOK := helper.GetStringArray(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if len(array) != 1 {
			t.Fatalf("not match")
		} else if array[0] != value {
			t.Fatalf("not match")
		}
	}
}

func Test_int8(t *testing.T) {
	value := int8(8)
	helper := NewInterfaceHelper(value)

	if !helper.IsNumber() {
		t.Fatalf("not match")
	} else if helper.IsBool() {
		t.Fatalf("not match")
	} else if helper.IsFloat() {
		t.Fatalf("not match")
	} else if !helper.IsInt() {
		t.Fatalf("not match")
	} else if helper.IsString() {
		t.Fatalf("not match")
	} else if helper.IsUint() {
		t.Fatalf("not match")
	}

	if valueFloat, assertionOK := helper.GetNumber(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if valueFloat != float64(value) {
			t.Fatalf("not match: %f vs %f", valueFloat, float64(value))
		}
	}

	if _, assertionOK := helper.GetBool(); assertionOK {
		t.Fatalf("not match")
	} else if _, assertionOK := helper.GetString(); assertionOK {
		t.Fatalf("not match")
	}

	if array, assertionOK := helper.GetNumberArray(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if len(array) != 1 {
			t.Fatalf("not match")
		} else if array[0] != float64(value) {
			t.Fatalf("not match")
		}
	}

	if _, assertionOK := helper.GetStringArray(); assertionOK {
		t.Fatalf("not match")
	}
}

func Test_uint8(t *testing.T) {
	value := uint8(8)
	helper := NewInterfaceHelper(value)

	if !helper.IsNumber() {
		t.Fatalf("not match")
	} else if helper.IsBool() {
		t.Fatalf("not match")
	} else if helper.IsFloat() {
		t.Fatalf("not match")
	} else if helper.IsInt() {
		t.Fatalf("not match")
	} else if helper.IsString() {
		t.Fatalf("not match")
	} else if !helper.IsUint() {
		t.Fatalf("not match")
	}

	if valueFloat, assertionOK := helper.GetNumber(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if valueFloat != float64(value) {
			t.Fatalf("not match: %f vs %f", valueFloat, float64(value))
		}
	}

	if _, assertionOK := helper.GetBool(); assertionOK {
		t.Fatalf("not match")
	} else if _, assertionOK := helper.GetString(); assertionOK {
		t.Fatalf("not match")
	}

	if array, assertionOK := helper.GetNumberArray(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if len(array) != 1 {
			t.Fatalf("not match")
		} else if array[0] != float64(value) {
			t.Fatalf("not match")
		}
	}

	if _, assertionOK := helper.GetStringArray(); assertionOK {
		t.Fatalf("not match")
	}
}

func Test_int16(t *testing.T) {
	value := int16(8)
	helper := NewInterfaceHelper(value)

	if !helper.IsNumber() {
		t.Fatalf("not match")
	} else if helper.IsBool() {
		t.Fatalf("not match")
	} else if helper.IsFloat() {
		t.Fatalf("not match")
	} else if !helper.IsInt() {
		t.Fatalf("not match")
	} else if helper.IsString() {
		t.Fatalf("not match")
	} else if helper.IsUint() {
		t.Fatalf("not match")
	}

	if valueFloat, assertionOK := helper.GetNumber(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if valueFloat != float64(value) {
			t.Fatalf("not match: %f vs %f", valueFloat, float64(value))
		}
	}

	if _, assertionOK := helper.GetBool(); assertionOK {
		t.Fatalf("not match")
	} else if _, assertionOK := helper.GetString(); assertionOK {
		t.Fatalf("not match")
	}

	if array, assertionOK := helper.GetNumberArray(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if len(array) != 1 {
			t.Fatalf("not match")
		} else if array[0] != float64(value) {
			t.Fatalf("not match")
		}
	}

	if _, assertionOK := helper.GetStringArray(); assertionOK {
		t.Fatalf("not match")
	}
}

func Test_uint16(t *testing.T) {
	value := uint16(8)
	helper := NewInterfaceHelper(value)

	if !helper.IsNumber() {
		t.Fatalf("not match")
	} else if helper.IsBool() {
		t.Fatalf("not match")
	} else if helper.IsFloat() {
		t.Fatalf("not match")
	} else if helper.IsInt() {
		t.Fatalf("not match")
	} else if helper.IsString() {
		t.Fatalf("not match")
	} else if !helper.IsUint() {
		t.Fatalf("not match")
	}

	if valueFloat, assertionOK := helper.GetNumber(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if valueFloat != float64(value) {
			t.Fatalf("not match: %f vs %f", valueFloat, float64(value))
		}
	}

	if _, assertionOK := helper.GetBool(); assertionOK {
		t.Fatalf("not match")
	} else if _, assertionOK := helper.GetString(); assertionOK {
		t.Fatalf("not match")
	}

	if array, assertionOK := helper.GetNumberArray(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if len(array) != 1 {
			t.Fatalf("not match")
		} else if array[0] != float64(value) {
			t.Fatalf("not match")
		}
	}

	if _, assertionOK := helper.GetStringArray(); assertionOK {
		t.Fatalf("not match")
	}
}

func Test_int32(t *testing.T) {
	value := int32(8)
	helper := NewInterfaceHelper(value)

	if !helper.IsNumber() {
		t.Fatalf("not match")
	} else if helper.IsBool() {
		t.Fatalf("not match")
	} else if helper.IsFloat() {
		t.Fatalf("not match")
	} else if !helper.IsInt() {
		t.Fatalf("not match")
	} else if helper.IsString() {
		t.Fatalf("not match")
	} else if helper.IsUint() {
		t.Fatalf("not match")
	}

	if valueFloat, assertionOK := helper.GetNumber(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if valueFloat != float64(value) {
			t.Fatalf("not match: %f vs %f", valueFloat, float64(value))
		}
	}

	if _, assertionOK := helper.GetBool(); assertionOK {
		t.Fatalf("not match")
	} else if _, assertionOK := helper.GetString(); assertionOK {
		t.Fatalf("not match")
	}

	if array, assertionOK := helper.GetNumberArray(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if len(array) != 1 {
			t.Fatalf("not match")
		} else if array[0] != float64(value) {
			t.Fatalf("not match")
		}
	}

	if _, assertionOK := helper.GetStringArray(); assertionOK {
		t.Fatalf("not match")
	}
}

func Test_uint32(t *testing.T) {
	value := uint32(8)
	helper := NewInterfaceHelper(value)

	if !helper.IsNumber() {
		t.Fatalf("not match")
	} else if helper.IsBool() {
		t.Fatalf("not match")
	} else if helper.IsFloat() {
		t.Fatalf("not match")
	} else if helper.IsInt() {
		t.Fatalf("not match")
	} else if helper.IsString() {
		t.Fatalf("not match")
	} else if !helper.IsUint() {
		t.Fatalf("not match")
	}

	if valueFloat, assertionOK := helper.GetNumber(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if valueFloat != float64(value) {
			t.Fatalf("not match: %f vs %f", valueFloat, float64(value))
		}
	}

	if _, assertionOK := helper.GetBool(); assertionOK {
		t.Fatalf("not match")
	} else if _, assertionOK := helper.GetString(); assertionOK {
		t.Fatalf("not match")
	}

	if array, assertionOK := helper.GetNumberArray(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if len(array) != 1 {
			t.Fatalf("not match")
		} else if array[0] != float64(value) {
			t.Fatalf("not match")
		}
	}

	if _, assertionOK := helper.GetStringArray(); assertionOK {
		t.Fatalf("not match")
	}
}

func Test_int64(t *testing.T) {
	value := int64(8)
	helper := NewInterfaceHelper(value)

	if !helper.IsNumber() {
		t.Fatalf("not match")
	} else if helper.IsBool() {
		t.Fatalf("not match")
	} else if helper.IsFloat() {
		t.Fatalf("not match")
	} else if !helper.IsInt() {
		t.Fatalf("not match")
	} else if helper.IsString() {
		t.Fatalf("not match")
	} else if helper.IsUint() {
		t.Fatalf("not match")
	}

	if valueFloat, assertionOK := helper.GetNumber(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if valueFloat != float64(value) {
			t.Fatalf("not match: %f vs %f", valueFloat, float64(value))
		}
	}

	if _, assertionOK := helper.GetBool(); assertionOK {
		t.Fatalf("not match")
	} else if _, assertionOK := helper.GetString(); assertionOK {
		t.Fatalf("not match")
	}

	if array, assertionOK := helper.GetNumberArray(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if len(array) != 1 {
			t.Fatalf("not match")
		} else if array[0] != float64(value) {
			t.Fatalf("not match")
		}
	}

	if _, assertionOK := helper.GetStringArray(); assertionOK {
		t.Fatalf("not match")
	}
}

func Test_uint64(t *testing.T) {
	value := uint64(8)
	helper := NewInterfaceHelper(value)

	if !helper.IsNumber() {
		t.Fatalf("not match")
	} else if helper.IsBool() {
		t.Fatalf("not match")
	} else if helper.IsFloat() {
		t.Fatalf("not match")
	} else if helper.IsInt() {
		t.Fatalf("not match")
	} else if helper.IsString() {
		t.Fatalf("not match")
	} else if !helper.IsUint() {
		t.Fatalf("not match")
	}

	if valueFloat, assertionOK := helper.GetNumber(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if valueFloat != float64(value) {
			t.Fatalf("not match: %f vs %f", valueFloat, float64(value))
		}
	}

	if _, assertionOK := helper.GetBool(); assertionOK {
		t.Fatalf("not match")
	} else if _, assertionOK := helper.GetString(); assertionOK {
		t.Fatalf("not match")
	}

	if array, assertionOK := helper.GetNumberArray(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if len(array) != 1 {
			t.Fatalf("not match")
		} else if array[0] != float64(value) {
			t.Fatalf("not match")
		}
	}

	if _, assertionOK := helper.GetStringArray(); assertionOK {
		t.Fatalf("not match")
	}
}

func Test_int(t *testing.T) {
	value := int(8)
	helper := NewInterfaceHelper(value)

	if !helper.IsNumber() {
		t.Fatalf("not match")
	} else if helper.IsBool() {
		t.Fatalf("not match")
	} else if helper.IsFloat() {
		t.Fatalf("not match")
	} else if !helper.IsInt() {
		t.Fatalf("not match")
	} else if helper.IsString() {
		t.Fatalf("not match")
	} else if helper.IsUint() {
		t.Fatalf("not match")
	}

	if valueFloat, assertionOK := helper.GetNumber(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if valueFloat != float64(value) {
			t.Fatalf("not match: %f vs %f", valueFloat, float64(value))
		}
	}

	if _, assertionOK := helper.GetBool(); assertionOK {
		t.Fatalf("not match")
	} else if _, assertionOK := helper.GetString(); assertionOK {
		t.Fatalf("not match")
	}

	if array, assertionOK := helper.GetNumberArray(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if len(array) != 1 {
			t.Fatalf("not match")
		} else if array[0] != float64(value) {
			t.Fatalf("not match")
		}
	}

	if _, assertionOK := helper.GetStringArray(); assertionOK {
		t.Fatalf("not match")
	}
}

func Test_uint(t *testing.T) {
	value := uint(8)
	helper := NewInterfaceHelper(value)

	if !helper.IsNumber() {
		t.Fatalf("not match")
	} else if helper.IsBool() {
		t.Fatalf("not match")
	} else if helper.IsFloat() {
		t.Fatalf("not match")
	} else if helper.IsInt() {
		t.Fatalf("not match")
	} else if helper.IsString() {
		t.Fatalf("not match")
	} else if !helper.IsUint() {
		t.Fatalf("not match")
	}

	if valueFloat, assertionOK := helper.GetNumber(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if valueFloat != float64(value) {
			t.Fatalf("not match: %f vs %f", valueFloat, float64(value))
		}
	}

	if _, assertionOK := helper.GetBool(); assertionOK {
		t.Fatalf("not match")
	} else if _, assertionOK := helper.GetString(); assertionOK {
		t.Fatalf("not match")
	}

	if array, assertionOK := helper.GetNumberArray(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if len(array) != 1 {
			t.Fatalf("not match")
		} else if array[0] != float64(value) {
			t.Fatalf("not match")
		}
	}

	if _, assertionOK := helper.GetStringArray(); assertionOK {
		t.Fatalf("not match")
	}
}

func Test_float32(t *testing.T) {
	value := float32(8.1)
	helper := NewInterfaceHelper(value)

	if !helper.IsNumber() {
		t.Fatalf("not match")
	} else if helper.IsBool() {
		t.Fatalf("not match")
	} else if !helper.IsFloat() {
		t.Fatalf("not match")
	} else if helper.IsInt() {
		t.Fatalf("not match")
	} else if helper.IsString() {
		t.Fatalf("not match")
	} else if helper.IsUint() {
		t.Fatalf("not match")
	}

	if valueFloat, assertionOK := helper.GetNumber(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if valueFloat != float64(value) {
			t.Fatalf("not match: %f vs %f", valueFloat, float64(value))
		}
	}

	if _, assertionOK := helper.GetBool(); assertionOK {
		t.Fatalf("not match")
	} else if _, assertionOK := helper.GetString(); assertionOK {
		t.Fatalf("not match")
	}

	if array, assertionOK := helper.GetNumberArray(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if len(array) != 1 {
			t.Fatalf("not match")
		} else if array[0] != float64(value) {
			t.Fatalf("not match")
		}
	}

	if _, assertionOK := helper.GetStringArray(); assertionOK {
		t.Fatalf("not match")
	}
}

func Test_float64(t *testing.T) {
	value := float64(8.1)
	helper := NewInterfaceHelper(value)

	if !helper.IsNumber() {
		t.Fatalf("not match")
	} else if helper.IsBool() {
		t.Fatalf("not match")
	} else if !helper.IsFloat() {
		t.Fatalf("not match")
	} else if helper.IsInt() {
		t.Fatalf("not match")
	} else if helper.IsString() {
		t.Fatalf("not match")
	} else if helper.IsUint() {
		t.Fatalf("not match")
	}

	if valueFloat, assertionOK := helper.GetNumber(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if valueFloat != float64(value) {
			t.Fatalf("not match: %f vs %f", valueFloat, float64(value))
		}
	}

	if _, assertionOK := helper.GetBool(); assertionOK {
		t.Fatalf("not match")
	} else if _, assertionOK := helper.GetString(); assertionOK {
		t.Fatalf("not match")
	}

	if array, assertionOK := helper.GetNumberArray(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if len(array) != 1 {
			t.Fatalf("not match")
		} else if array[0] != float64(value) {
			t.Fatalf("not match")
		}
	}

	if _, assertionOK := helper.GetStringArray(); assertionOK {
		t.Fatalf("not match")
	}
}

func Test_bool(t *testing.T) {
	value := true
	helper := NewInterfaceHelper(value)

	if helper.IsNumber() {
		t.Fatalf("not match")
	} else if !helper.IsBool() {
		t.Fatalf("not match")
	} else if helper.IsFloat() {
		t.Fatalf("not match")
	} else if helper.IsInt() {
		t.Fatalf("not match")
	} else if helper.IsString() {
		t.Fatalf("not match")
	} else if helper.IsUint() {
		t.Fatalf("not match")
	}

	if _, assertionOK := helper.GetNumber(); assertionOK {
		t.Fatalf("not match")
	}

	if valueBool, assertionOK := helper.GetBool(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if valueBool != value {
			t.Fatalf("not match")
		}
	}

	if _, assertionOK := helper.GetString(); assertionOK {
		t.Fatalf("not match")
	}

	if array, assertionOK := helper.GetBoolArray(); !assertionOK {
		t.Fatalf("not match")
	} else {
		if len(array) != 1 {
			t.Fatalf("not match")
		} else if array[0] != value {
			t.Fatalf("not match")
		}
	}

	if _, assertionOK := helper.GetNumberArray(); assertionOK {
		t.Fatalf("not match")
	}

	if _, assertionOK := helper.GetStringArray(); assertionOK {
		t.Fatalf("not match")
	}
}

func Test_SetString(t *testing.T) {
	value := "true"
	helper := NewInterfaceHelper(&value)

	{
		fromValue := "test"
		if !helper.Set(fromValue) {
			t.Fatalf("cannot set value: %v", fromValue)
		} else {
			if newValue, getRet := helper.GetString(); getRet == false {
				t.Fatalf("cannot get value")
			} else if newValue != fromValue {
				t.Fatalf("not match value: %v vs %v", newValue, fromValue)
			}
		}
	}

	{
		fromValue := true

		if helper.Set(fromValue) {
			t.Fatalf("can set value: %v", fromValue)
		}
	}

	{
		fromValue := int(-10)

		if helper.Set(fromValue) {
			t.Fatalf("can set value: %v", fromValue)
		}
	}

	{
		fromValue := int8(-10)

		if helper.Set(fromValue) {
			t.Fatalf("can set value: %v", fromValue)
		}
	}

	{
		fromValue := int16(-10)

		if helper.Set(fromValue) {
			t.Fatalf("can set value: %v", fromValue)
		}
	}

	{
		fromValue := int32(-10)

		if helper.Set(fromValue) {
			t.Fatalf("can set value: %v", fromValue)
		}
	}

	{
		fromValue := int64(-10)

		if helper.Set(fromValue) {
			t.Fatalf("can set value: %v", fromValue)
		}
	}

	{
		fromValue := uint(10)

		if helper.Set(fromValue) {
			t.Fatalf("can set value: %v", fromValue)
		}
	}

	{
		fromValue := uint8(10)

		if helper.Set(fromValue) {
			t.Fatalf("can set value: %v", fromValue)
		}
	}

	{
		fromValue := uint16(10)

		if helper.Set(fromValue) {
			t.Fatalf("can set value: %v", fromValue)
		}
	}

	{
		fromValue := uint32(10)

		if helper.Set(fromValue) {
			t.Fatalf("can set value: %v", fromValue)
		}
	}

	{
		fromValue := uint64(10)

		if helper.Set(fromValue) {
			t.Fatalf("can set value: %v", fromValue)
		}
	}
}

func Test_SetBool(t *testing.T) {
	value := true
	helper := NewInterfaceHelper(&value)

	{
		fromValue := "false"

		if helper.Set(fromValue) {
			t.Fatalf("can set value: %v", fromValue)
		}
	}

	{
		fromValue := false

		if !helper.Set(fromValue) {
			t.Fatalf("cannot set value: %v", fromValue)
		} else {
			if newValue, getRet := helper.GetBool(); getRet == false {
				t.Fatalf("cannot get value")
			} else if newValue != fromValue {
				t.Fatalf("not match value: %v vs %v", newValue, fromValue)
			}
		}
	}

	{
		fromValue := int(-10)

		if helper.Set(fromValue) {
			t.Fatalf("can set value: %v", fromValue)
		}
	}

	{
		fromValue := int8(-10)

		if helper.Set(fromValue) {
			t.Fatalf("can set value: %v", fromValue)
		}
	}

	{
		fromValue := int16(-10)

		if helper.Set(fromValue) {
			t.Fatalf("can set value: %v", fromValue)
		}
	}

	{
		fromValue := int32(-10)

		if helper.Set(fromValue) {
			t.Fatalf("can set value: %v", fromValue)
		}
	}

	{
		fromValue := int64(-10)

		if helper.Set(fromValue) {
			t.Fatalf("can set value: %v", fromValue)
		}
	}

	{
		fromValue := uint(10)

		if helper.Set(fromValue) {
			t.Fatalf("can set value: %v", fromValue)
		}
	}

	{
		fromValue := uint8(10)

		if helper.Set(fromValue) {
			t.Fatalf("can set value: %v", fromValue)
		}
	}

	{
		fromValue := uint16(10)

		if helper.Set(fromValue) {
			t.Fatalf("can set value: %v", fromValue)
		}
	}

	{
		fromValue := uint32(10)

		if helper.Set(fromValue) {
			t.Fatalf("can set value: %v", fromValue)
		}
	}

	{
		fromValue := uint64(10)

		if helper.Set(fromValue) {
			t.Fatalf("can set value: %v", fromValue)
		}
	}
}

func Test_SetInt(t *testing.T) {
	value := int(100)
	helper := NewInterfaceHelper(&value)

	{
		fromValue := "false"

		if helper.Set(fromValue) {
			t.Fatalf("can set value: %v", fromValue)
		}
	}

	{
		fromValue := false

		if helper.Set(fromValue) {
			t.Fatalf("can set value: %v", fromValue)
		}
	}

	{
		fromValue := int(-10)

		if !helper.Set(fromValue) {
			t.Fatalf("cannot set value: %v", fromValue)
		} else {
			if newValue, getRet := helper.GetNumber(); getRet == false {
				t.Fatalf("cannot get value")
			} else if int(newValue) != fromValue {
				t.Fatalf("not match value: %v vs %v", newValue, fromValue)
			}
		}
	}

	{
		fromValue := int8(-10)

		if !helper.Set(fromValue) {
			t.Fatalf("cannot set value: %v", fromValue)
		} else {
			if newValue, getRet := helper.GetNumber(); getRet == false {
				t.Fatalf("cannot get value")
			} else if int8(newValue) != fromValue {
				t.Fatalf("not match value: %v vs %v", newValue, fromValue)
			}
		}
	}

	{
		fromValue := int16(-10)

		if !helper.Set(fromValue) {
			t.Fatalf("cannot set value: %v", fromValue)
		} else {
			if newValue, getRet := helper.GetNumber(); getRet == false {
				t.Fatalf("cannot get value")
			} else if int16(newValue) != fromValue {
				t.Fatalf("not match value: %v vs %v", newValue, fromValue)
			}
		}
	}

	{
		fromValue := int32(-10)

		if !helper.Set(fromValue) {
			t.Fatalf("cannot set value: %v", fromValue)
		} else {
			if newValue, getRet := helper.GetNumber(); getRet == false {
				t.Fatalf("cannot get value")
			} else if int32(newValue) != fromValue {
				t.Fatalf("not match value: %v vs %v", newValue, fromValue)
			}
		}
	}

	{
		fromValue := int64(-10)

		if !helper.Set(fromValue) {
			t.Fatalf("cannot set value: %v", fromValue)
		} else {
			if newValue, getRet := helper.GetNumber(); getRet == false {
				t.Fatalf("cannot get value")
			} else if int64(newValue) != fromValue {
				t.Fatalf("not match value: %v vs %v", newValue, fromValue)
			}
		}
	}

	{
		fromValue := uint(10)

		if helper.Set(fromValue) {
			t.Fatalf("cannot set value: %v", fromValue)
		}
	}

	{
		fromValue := uint8(10)

		if helper.Set(fromValue) {
			t.Fatalf("cannot set value: %v", fromValue)
		}
	}

	{
		fromValue := uint16(10)

		if helper.Set(fromValue) {
			t.Fatalf("cannot set value: %v", fromValue)
		}
	}

	{
		fromValue := uint32(10)

		if helper.Set(fromValue) {
			t.Fatalf("cannot set value: %v", fromValue)
		}
	}

	{
		fromValue := uint64(10)

		if helper.Set(fromValue) {
			t.Fatalf("cannot set value: %v", fromValue)
		}
	}
}

func Test_SetUint(t *testing.T) {
	value := uint(100)
	helper := NewInterfaceHelper(&value)

	{
		fromValue := "false"

		if helper.Set(fromValue) {
			t.Fatalf("can set value: %v", fromValue)
		}
	}

	{
		fromValue := false

		if helper.Set(fromValue) {
			t.Fatalf("can set value: %v", fromValue)
		}
	}

	{
		fromValue := uint(10)

		if !helper.Set(fromValue) {
			t.Fatalf("cannot set value: %v", fromValue)
		} else {
			if newValue, getRet := helper.GetNumber(); getRet == false {
				t.Fatalf("cannot get value")
			} else if uint(newValue) != fromValue {
				t.Fatalf("not match value: %v vs %v", newValue, fromValue)
			}
		}
	}

	{
		fromValue := uint8(10)

		if !helper.Set(fromValue) {
			t.Fatalf("cannot set value: %v", fromValue)
		} else {
			if newValue, getRet := helper.GetNumber(); getRet == false {
				t.Fatalf("cannot get value")
			} else if uint8(newValue) != fromValue {
				t.Fatalf("not match value: %v vs %v", newValue, fromValue)
			}
		}
	}

	{
		fromValue := uint16(10)

		if !helper.Set(fromValue) {
			t.Fatalf("cannot set value: %v", fromValue)
		} else {
			if newValue, getRet := helper.GetNumber(); getRet == false {
				t.Fatalf("cannot get value")
			} else if uint16(newValue) != fromValue {
				t.Fatalf("not match value: %v vs %v", newValue, fromValue)
			}
		}
	}

	{
		fromValue := uint32(10)

		if !helper.Set(fromValue) {
			t.Fatalf("cannot set value: %v", fromValue)
		} else {
			if newValue, getRet := helper.GetNumber(); getRet == false {
				t.Fatalf("cannot get value")
			} else if uint32(newValue) != fromValue {
				t.Fatalf("not match value: %v vs %v", newValue, fromValue)
			}
		}
	}

	{
		fromValue := uint64(10)

		if !helper.Set(fromValue) {
			t.Fatalf("cannot set value: %v", fromValue)
		} else {
			if newValue, getRet := helper.GetNumber(); getRet == false {
				t.Fatalf("cannot get value")
			} else if uint64(newValue) != fromValue {
				t.Fatalf("not match value: %v vs %v", newValue, fromValue)
			}
		}
	}

	{
		fromValue := int(-10)

		if helper.Set(fromValue) {
			t.Fatalf("cannot set value: %v", fromValue)
		}
	}

	{
		fromValue := int8(-10)

		if helper.Set(fromValue) {
			t.Fatalf("cannot set value: %v", fromValue)
		}
	}

	{
		fromValue := int16(-10)

		if helper.Set(fromValue) {
			t.Fatalf("cannot set value: %v", fromValue)
		}
	}

	{
		fromValue := int32(-10)

		if helper.Set(fromValue) {
			t.Fatalf("cannot set value: %v", fromValue)
		}
	}

	{
		fromValue := int64(-10)

		if helper.Set(fromValue) {
			t.Fatalf("cannot set value: %v", fromValue)
		}
	}
}
