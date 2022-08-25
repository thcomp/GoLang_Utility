package utility

func TernaryOpString(conv bool, v1, v2 string) string {
	if conv {
		return v1
	} else {
		return v2
	}
}

func TernaryOpStringFunc(conv bool, f1, f2 func() string) string {
	if conv {
		return f1()
	} else {
		return f2()
	}
}

func TernaryOpStringFuncWithParams(conv bool, f1, f2 func(...interface{}) string, params ...interface{}) string {
	if conv {
		return f1(params...)
	} else {
		return f2(params...)
	}
}

func TernaryOpInt(conv bool, v1, v2 int) int {
	if conv {
		return v1
	} else {
		return v2
	}
}

func TernaryOpIntFunc(conv bool, f1, f2 func() int) int {
	if conv {
		return f1()
	} else {
		return f2()
	}
}

func TernaryOpIntFuncWithParams(conv bool, f1, f2 func(...interface{}) int, params ...interface{}) int {
	if conv {
		return f1(params...)
	} else {
		return f2(params...)
	}
}

func TernaryOpInt8(conv bool, v1, v2 int8) int8 {
	if conv {
		return v1
	} else {
		return v2
	}
}

func TernaryOpInt8Func(conv bool, f1, f2 func() int8) int8 {
	if conv {
		return f1()
	} else {
		return f2()
	}
}

func TernaryOpInt8FuncWithParams(conv bool, f1, f2 func(...interface{}) int8, params ...interface{}) int8 {
	if conv {
		return f1(params...)
	} else {
		return f2(params...)
	}
}

func TernaryOpInt16(conv bool, v1, v2 int16) int16 {
	if conv {
		return v1
	} else {
		return v2
	}
}

func TernaryOpInt16Func(conv bool, f1, f2 func() int16) int16 {
	if conv {
		return f1()
	} else {
		return f2()
	}
}

func TernaryOpInt16FuncWithParams(conv bool, f1, f2 func(...interface{}) int16, params ...interface{}) int16 {
	if conv {
		return f1(params...)
	} else {
		return f2(params...)
	}
}

func TernaryOpInt32(conv bool, v1, v2 int32) int32 {
	if conv {
		return v1
	} else {
		return v2
	}
}

func TernaryOpInt32Func(conv bool, f1, f2 func() int32) int32 {
	if conv {
		return f1()
	} else {
		return f2()
	}
}

func TernaryOpInt32FuncWithParams(conv bool, f1, f2 func(...interface{}) int32, params ...interface{}) int32 {
	if conv {
		return f1(params...)
	} else {
		return f2(params...)
	}
}

func TernaryOpInt64(conv bool, v1, v2 int64) int64 {
	if conv {
		return v1
	} else {
		return v2
	}
}

func TernaryOpInt64Func(conv bool, f1, f2 func() int64) int64 {
	if conv {
		return f1()
	} else {
		return f2()
	}
}

func TernaryOpInt64FuncWithParams(conv bool, f1, f2 func(...interface{}) int64, params ...interface{}) int64 {
	if conv {
		return f1(params...)
	} else {
		return f2(params...)
	}
}

func TernaryOpUint(conv bool, v1, v2 uint) uint {
	if conv {
		return v1
	} else {
		return v2
	}
}

func TernaryOpUintFunc(conv bool, f1, f2 func() uint) uint {
	if conv {
		return f1()
	} else {
		return f2()
	}
}

func TernaryOpUintFuncWithParams(conv bool, f1, f2 func(...interface{}) uint, params ...interface{}) uint {
	if conv {
		return f1(params...)
	} else {
		return f2(params...)
	}
}

func TernaryOpUint8(conv bool, v1, v2 uint8) uint8 {
	if conv {
		return v1
	} else {
		return v2
	}
}

func TernaryOpUint8Func(conv bool, f1, f2 func() uint8) uint8 {
	if conv {
		return f1()
	} else {
		return f2()
	}
}

func TernaryOpUint8FuncWithParams(conv bool, f1, f2 func(...interface{}) uint8, params ...interface{}) uint8 {
	if conv {
		return f1(params...)
	} else {
		return f2(params...)
	}
}

func TernaryOpUint16(conv bool, v1, v2 uint16) uint16 {
	if conv {
		return v1
	} else {
		return v2
	}
}

func TernaryOpUint16Func(conv bool, f1, f2 func() uint16) uint16 {
	if conv {
		return f1()
	} else {
		return f2()
	}
}

func TernaryOpUint16FuncWithParams(conv bool, f1, f2 func(...interface{}) uint16, params ...interface{}) uint16 {
	if conv {
		return f1(params...)
	} else {
		return f2(params...)
	}
}

func TernaryOpUint32(conv bool, v1, v2 uint32) uint32 {
	if conv {
		return v1
	} else {
		return v2
	}
}

func TernaryOpUint32Func(conv bool, f1, f2 func() uint32) uint32 {
	if conv {
		return f1()
	} else {
		return f2()
	}
}

func TernaryOpUint32FuncWithParams(conv bool, f1, f2 func(...interface{}) uint32, params ...interface{}) uint32 {
	if conv {
		return f1(params...)
	} else {
		return f2(params...)
	}
}

func TernaryOpUint64(conv bool, v1, v2 uint64) uint64 {
	if conv {
		return v1
	} else {
		return v2
	}
}

func TernaryOpUint64Func(conv bool, f1, f2 func() uint64) uint64 {
	if conv {
		return f1()
	} else {
		return f2()
	}
}

func TernaryOpUint64FuncWithParams(conv bool, f1, f2 func(...interface{}) uint64, params ...interface{}) uint64 {
	if conv {
		return f1(params...)
	} else {
		return f2(params...)
	}
}

func TernaryOpFloat32(conv bool, v1, v2 float32) float32 {
	if conv {
		return v1
	} else {
		return v2
	}
}

func TernaryOpFloat32Func(conv bool, f1, f2 func() float32) float32 {
	if conv {
		return f1()
	} else {
		return f2()
	}
}

func TernaryOpFloat32FuncWithParams(conv bool, f1, f2 func(...interface{}) float32, params ...interface{}) float32 {
	if conv {
		return f1(params...)
	} else {
		return f2(params...)
	}
}

func TernaryOpFloat64(conv bool, v1, v2 float64) float64 {
	if conv {
		return v1
	} else {
		return v2
	}
}

func TernaryOpFloat64Func(conv bool, f1, f2 func() float64) float64 {
	if conv {
		return f1()
	} else {
		return f2()
	}
}

func TernaryOpFloat64FuncWithParams(conv bool, f1, f2 func(...interface{}) float64, params ...interface{}) float64 {
	if conv {
		return f1(params...)
	} else {
		return f2(params...)
	}
}

func TernaryOpInterface(conv bool, v1, v2 interface{}) interface{} {
	if conv {
		return v1
	} else {
		return v2
	}
}

func TernaryOpInterfaceFunc(conv bool, f1, f2 func() interface{}) interface{} {
	if conv {
		return f1()
	} else {
		return f2()
	}
}

func TernaryOpInterfaceFuncWithParams(conv bool, f1, f2 func(...interface{}) interface{}, params ...interface{}) interface{} {
	if conv {
		return f1(params...)
	} else {
		return f2(params...)
	}
}
