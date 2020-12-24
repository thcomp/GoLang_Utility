package utility

import (
	"testing"
)

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
