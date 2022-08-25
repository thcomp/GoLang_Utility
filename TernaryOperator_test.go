package utility

import (
	"reflect"
	"testing"
)

type Param struct {
	value1 interface{}
	value2 interface{}
}

var params = []Param{
	{
		value1: int(0),
		value2: int(1),
	},
	{
		value1: int8(0),
		value2: int8(1),
	},
	{
		value1: int16(0),
		value2: int16(1),
	},
	{
		value1: int32(0),
		value2: int32(1),
	},
	{
		value1: int64(0),
		value2: int64(1),
	},
	{
		value1: uint(0),
		value2: uint(1),
	},
	{
		value1: uint8(0),
		value2: uint8(1),
	},
	{
		value1: uint16(0),
		value2: uint16(1),
	},
	{
		value1: uint32(0),
		value2: uint32(1),
	},
	{
		value1: uint64(0),
		value2: uint64(1),
	},
	{
		value1: float32(0),
		value2: float32(1),
	},
	{
		value1: float64(0),
		value2: float64(1),
	},
	{
		value1: "0",
		value2: "1",
	},
	{
		value1: map[string]interface{}{"test": "0"},
		value2: map[string]interface{}{"test": "1"},
	},
}

func TestTernaryOperatorAll(t *testing.T) {
	TestTernaryOperator1(t)
	TestTernaryOperator2(t)
	TestTernaryOperator3(t)
}

func TestTernaryOperator1(t *testing.T) {
	for _, param := range params {
		conditions := []bool{true, false}

		switch reflect.TypeOf(param.value1).Kind() {
		case reflect.Int:
			value1, _ := param.value1.(int)
			value2, _ := param.value2.(int)
			targetFunction := TernaryOpInt
			expects := []int{value1, value2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Int8:
			value1, _ := param.value1.(int8)
			value2, _ := param.value2.(int8)
			targetFunction := TernaryOpInt8
			expects := []int8{value1, value2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Int16:
			value1, _ := param.value1.(int16)
			value2, _ := param.value2.(int16)
			targetFunction := TernaryOpInt16
			expects := []int16{value1, value2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Int32:
			value1, _ := param.value1.(int32)
			value2, _ := param.value2.(int32)
			targetFunction := TernaryOpInt32
			expects := []int32{value1, value2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Int64:
			value1, _ := param.value1.(int64)
			value2, _ := param.value2.(int64)
			targetFunction := TernaryOpInt64
			expects := []int64{value1, value2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Uint:
			value1, _ := param.value1.(uint)
			value2, _ := param.value2.(uint)
			targetFunction := TernaryOpUint
			expects := []uint{value1, value2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Uint8:
			value1, _ := param.value1.(uint8)
			value2, _ := param.value2.(uint8)
			targetFunction := TernaryOpUint8
			expects := []uint8{value1, value2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Uint16:
			value1, _ := param.value1.(uint16)
			value2, _ := param.value2.(uint16)
			targetFunction := TernaryOpUint16
			expects := []uint16{value1, value2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Uint32:
			value1, _ := param.value1.(uint32)
			value2, _ := param.value2.(uint32)
			targetFunction := TernaryOpUint32
			expects := []uint32{value1, value2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Uint64:
			value1, _ := param.value1.(uint64)
			value2, _ := param.value2.(uint64)
			targetFunction := TernaryOpUint64
			expects := []uint64{value1, value2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Float32:
			value1, _ := param.value1.(float32)
			value2, _ := param.value2.(float32)
			targetFunction := TernaryOpFloat32
			expects := []float32{value1, value2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Float64:
			value1, _ := param.value1.(float64)
			value2, _ := param.value2.(float64)
			targetFunction := TernaryOpFloat64
			expects := []float64{value1, value2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.String:
			value1, _ := param.value1.(string)
			value2, _ := param.value2.(string)
			targetFunction := TernaryOpString
			expects := []string{value1, value2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Interface:
			value1, _ := param.value1.(interface{})
			value2, _ := param.value2.(interface{})
			targetFunction := TernaryOpInterface
			expects := []interface{}{value1, value2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		}
	}
}

func TestTernaryOperator2(t *testing.T) {
	for _, param := range params {
		conditions := []bool{true, false}

		switch reflect.TypeOf(param.value1).Kind() {
		case reflect.Int:
			realValue1, _ := param.value1.(int)
			realValue2, _ := param.value2.(int)
			value1 := func() int { return realValue1 }
			value2 := func() int { return realValue2 }
			targetFunction := TernaryOpIntFunc
			expects := []int{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Int8:
			realValue1, _ := param.value1.(int8)
			realValue2, _ := param.value2.(int8)
			value1 := func() int8 { return realValue1 }
			value2 := func() int8 { return realValue2 }
			targetFunction := TernaryOpInt8Func
			expects := []int8{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Int16:
			realValue1, _ := param.value1.(int16)
			realValue2, _ := param.value2.(int16)
			value1 := func() int16 { return realValue1 }
			value2 := func() int16 { return realValue2 }
			targetFunction := TernaryOpInt16Func
			expects := []int16{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Int32:
			realValue1, _ := param.value1.(int32)
			realValue2, _ := param.value2.(int32)
			value1 := func() int32 { return realValue1 }
			value2 := func() int32 { return realValue2 }
			targetFunction := TernaryOpInt32Func
			expects := []int32{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Int64:
			realValue1, _ := param.value1.(int64)
			realValue2, _ := param.value2.(int64)
			value1 := func() int64 { return realValue1 }
			value2 := func() int64 { return realValue2 }
			targetFunction := TernaryOpInt64Func
			expects := []int64{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Uint:
			realValue1, _ := param.value1.(uint)
			realValue2, _ := param.value2.(uint)
			value1 := func() uint { return realValue1 }
			value2 := func() uint { return realValue2 }
			targetFunction := TernaryOpUintFunc
			expects := []uint{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Uint8:
			realValue1, _ := param.value1.(uint8)
			realValue2, _ := param.value2.(uint8)
			value1 := func() uint8 { return realValue1 }
			value2 := func() uint8 { return realValue2 }
			targetFunction := TernaryOpUint8Func
			expects := []uint8{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Uint16:
			realValue1, _ := param.value1.(uint16)
			realValue2, _ := param.value2.(uint16)
			value1 := func() uint16 { return realValue1 }
			value2 := func() uint16 { return realValue2 }
			targetFunction := TernaryOpUint16Func
			expects := []uint16{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Uint32:
			realValue1, _ := param.value1.(uint32)
			realValue2, _ := param.value2.(uint32)
			value1 := func() uint32 { return realValue1 }
			value2 := func() uint32 { return realValue2 }
			targetFunction := TernaryOpUint32Func
			expects := []uint32{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Uint64:
			realValue1, _ := param.value1.(uint64)
			realValue2, _ := param.value2.(uint64)
			value1 := func() uint64 { return realValue1 }
			value2 := func() uint64 { return realValue2 }
			targetFunction := TernaryOpUint64Func
			expects := []uint64{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Float32:
			realValue1, _ := param.value1.(float32)
			realValue2, _ := param.value2.(float32)
			value1 := func() float32 { return realValue1 }
			value2 := func() float32 { return realValue2 }
			targetFunction := TernaryOpFloat32Func
			expects := []float32{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Float64:
			realValue1, _ := param.value1.(float64)
			realValue2, _ := param.value2.(float64)
			value1 := func() float64 { return realValue1 }
			value2 := func() float64 { return realValue2 }
			targetFunction := TernaryOpFloat64Func
			expects := []float64{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.String:
			realValue1, _ := param.value1.(string)
			realValue2, _ := param.value2.(string)
			value1 := func() string { return realValue1 }
			value2 := func() string { return realValue2 }
			targetFunction := TernaryOpStringFunc
			expects := []string{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Interface:
			realValue1, _ := param.value1.(interface{})
			realValue2, _ := param.value2.(interface{})
			value1 := func() interface{} { return realValue1 }
			value2 := func() interface{} { return realValue2 }
			targetFunction := TernaryOpInterfaceFunc
			expects := []interface{}{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		}
	}
}

func TestTernaryOperator3(t *testing.T) {
	for _, param := range params {
		conditions := []bool{true, false}

		switch reflect.TypeOf(param.value1).Kind() {
		case reflect.Int:
			realValue1, _ := param.value1.(int)
			realValue2, _ := param.value2.(int)
			value1 := func(...interface{}) int { return realValue1 }
			value2 := func(...interface{}) int { return realValue2 }
			targetFunction := TernaryOpIntFuncWithParams
			expects := []int{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Int8:
			realValue1, _ := param.value1.(int8)
			realValue2, _ := param.value2.(int8)
			value1 := func(...interface{}) int8 { return realValue1 }
			value2 := func(...interface{}) int8 { return realValue2 }
			targetFunction := TernaryOpInt8FuncWithParams
			expects := []int8{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Int16:
			realValue1, _ := param.value1.(int16)
			realValue2, _ := param.value2.(int16)
			value1 := func(...interface{}) int16 { return realValue1 }
			value2 := func(...interface{}) int16 { return realValue2 }
			targetFunction := TernaryOpInt16FuncWithParams
			expects := []int16{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Int32:
			realValue1, _ := param.value1.(int32)
			realValue2, _ := param.value2.(int32)
			value1 := func(...interface{}) int32 { return realValue1 }
			value2 := func(...interface{}) int32 { return realValue2 }
			targetFunction := TernaryOpInt32FuncWithParams
			expects := []int32{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Int64:
			realValue1, _ := param.value1.(int64)
			realValue2, _ := param.value2.(int64)
			value1 := func(...interface{}) int64 { return realValue1 }
			value2 := func(...interface{}) int64 { return realValue2 }
			targetFunction := TernaryOpInt64FuncWithParams
			expects := []int64{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Uint:
			realValue1, _ := param.value1.(uint)
			realValue2, _ := param.value2.(uint)
			value1 := func(...interface{}) uint { return realValue1 }
			value2 := func(...interface{}) uint { return realValue2 }
			targetFunction := TernaryOpUintFuncWithParams
			expects := []uint{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Uint8:
			realValue1, _ := param.value1.(uint8)
			realValue2, _ := param.value2.(uint8)
			value1 := func(...interface{}) uint8 { return realValue1 }
			value2 := func(...interface{}) uint8 { return realValue2 }
			targetFunction := TernaryOpUint8FuncWithParams
			expects := []uint8{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Uint16:
			realValue1, _ := param.value1.(uint16)
			realValue2, _ := param.value2.(uint16)
			value1 := func(...interface{}) uint16 { return realValue1 }
			value2 := func(...interface{}) uint16 { return realValue2 }
			targetFunction := TernaryOpUint16FuncWithParams
			expects := []uint16{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Uint32:
			realValue1, _ := param.value1.(uint32)
			realValue2, _ := param.value2.(uint32)
			value1 := func(...interface{}) uint32 { return realValue1 }
			value2 := func(...interface{}) uint32 { return realValue2 }
			targetFunction := TernaryOpUint32FuncWithParams
			expects := []uint32{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Uint64:
			realValue1, _ := param.value1.(uint64)
			realValue2, _ := param.value2.(uint64)
			value1 := func(...interface{}) uint64 { return realValue1 }
			value2 := func(...interface{}) uint64 { return realValue2 }
			targetFunction := TernaryOpUint64FuncWithParams
			expects := []uint64{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Float32:
			realValue1, _ := param.value1.(float32)
			realValue2, _ := param.value2.(float32)
			value1 := func(...interface{}) float32 { return realValue1 }
			value2 := func(...interface{}) float32 { return realValue2 }
			targetFunction := TernaryOpFloat32FuncWithParams
			expects := []float32{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Float64:
			realValue1, _ := param.value1.(float64)
			realValue2, _ := param.value2.(float64)
			value1 := func(...interface{}) float64 { return realValue1 }
			value2 := func(...interface{}) float64 { return realValue2 }
			targetFunction := TernaryOpFloat64FuncWithParams
			expects := []float64{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.String:
			realValue1, _ := param.value1.(string)
			realValue2, _ := param.value2.(string)
			value1 := func(...interface{}) string { return realValue1 }
			value2 := func(...interface{}) string { return realValue2 }
			targetFunction := TernaryOpStringFuncWithParams
			expects := []string{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		case reflect.Interface:
			realValue1, _ := param.value1.(interface{})
			realValue2, _ := param.value2.(interface{})
			value1 := func(...interface{}) interface{} { return realValue1 }
			value2 := func(...interface{}) interface{} { return realValue2 }
			targetFunction := TernaryOpInterfaceFuncWithParams
			expects := []interface{}{realValue1, realValue2}

			for index, condition := range conditions {
				expect := expects[index]

				if value := targetFunction(condition, value1, value2); value != expect {
					t.Fatalf("not match: type: %v, condition: %t, value: %v, expect: %v", reflect.TypeOf(param.value1).String(), condition, value, expect)
				}
			}
			break
		}
	}
}
