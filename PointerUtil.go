package utility

func ToIntPointer(value int) *int {
	return &value
}

func ToInt8Pointer(value int8) *int8 {
	return &value
}

func ToInt16Pointer(value int16) *int16 {
	return &value
}

func ToInt32Pointer(value int32) *int32 {
	return &value
}

func ToInt64Pointer(value int64) *int64 {
	return &value
}

func ToUintPointer(value uint) *uint {
	return &value
}

func ToUint8Pointer(value uint8) *uint8 {
	return &value
}

func ToUint16Pointer(value uint16) *uint16 {
	return &value
}

func ToUint32Pointer(value uint32) *uint32 {
	return &value
}

func ToUint64Pointer(value uint64) *uint64 {
	return &value
}

func ToFloat32Pointer(value float32) *float32 {
	return &value
}

func ToFloat64Pointer(value float64) *float64 {
	return &value
}

func ToStringPointer(value string) *string {
	return &value
}

func ToBytePointer(value byte) *byte {
	return &value
}

func ToBoolPointer(value bool) *bool {
	return &value
}

func ToInt(value *int, defaultValue int) (ret int) {
	if value != nil {
		ret = *value
	} else {
		ret = defaultValue
	}

	return ret
}

func ToInt8(value *int8, defaultValue int8) (ret int8) {
	if value != nil {
		ret = *value
	} else {
		ret = defaultValue
	}

	return ret
}

func ToInt16(value *int16, defaultValue int16) (ret int16) {
	if value != nil {
		ret = *value
	} else {
		ret = defaultValue
	}

	return ret
}

func ToInt32(value *int32, defaultValue int32) (ret int32) {
	if value != nil {
		ret = *value
	} else {
		ret = defaultValue
	}

	return ret
}

func ToInt64(value *int64, defaultValue int64) (ret int64) {
	if value != nil {
		ret = *value
	} else {
		ret = defaultValue
	}

	return ret
}

func ToUint(value *uint, defaultValue uint) (ret uint) {
	if value != nil {
		ret = *value
	} else {
		ret = defaultValue
	}

	return ret
}

func ToUint8(value *uint8, defaultValue uint8) (ret uint8) {
	if value != nil {
		ret = *value
	} else {
		ret = defaultValue
	}

	return ret
}

func ToUint16(value *uint16, defaultValue uint16) (ret uint16) {
	if value != nil {
		ret = *value
	} else {
		ret = defaultValue
	}

	return ret
}

func ToUint32(value *uint32, defaultValue uint32) (ret uint32) {
	if value != nil {
		ret = *value
	} else {
		ret = defaultValue
	}

	return ret
}

func ToUint64(value *uint64, defaultValue uint64) (ret uint64) {
	if value != nil {
		ret = *value
	} else {
		ret = defaultValue
	}

	return ret
}

func ToFloat32(value *float32, defaultValue float32) (ret float32) {
	if value != nil {
		ret = *value
	} else {
		ret = defaultValue
	}

	return ret
}

func ToFloat64(value *float64, defaultValue float64) (ret float64) {
	if value != nil {
		ret = *value
	} else {
		ret = defaultValue
	}

	return ret
}

func ToString(value *string, defaultValue string) (ret string) {
	if value != nil {
		ret = *value
	} else {
		ret = defaultValue
	}

	return ret
}

func ToByte(value *byte, defaultValue byte) (ret byte) {
	if value != nil {
		ret = *value
	} else {
		ret = defaultValue
	}

	return ret
}

func ToBool(value *bool, defaultValue bool) (ret bool) {
	if value != nil {
		ret = *value
	} else {
		ret = defaultValue
	}

	return ret
}
