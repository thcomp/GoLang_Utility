package utility

import (
	//"strconv"
	"errors"
	"reflect"
	"strings"
	"testing"
	//"time"
)

type SmallData struct {
	Value0 int32
}

type Data struct {
	Value0 int32
	Value1 float32
	Value2 float64
	Value3 string
}

type LargeData struct {
	Value0  int32
	Value1  float32
	Value2  float64
	Value3  string
	Value4  bool
	Value5  int8
	Value6  int16
	Value7  int64
	Value8  uint8
	Value9  uint16
	Value10 uint32
	Value11 uint64
	Value12 SmallData
}

type DataIf interface {
	FuncA(valueI int, valueS string) (int, error)
}

func (this *SmallData) FuncA(valueI int, valueS string) (int, error) {
	return valueI, errors.New(valueS)
}

func TestReflectHelper(t *testing.T) {
	var smallData SmallData
	var data Data
	var largeData LargeData

	var refHelperForSmall = NewReflectHelper(&smallData)
	var refHelper = NewReflectHelper(&data)
	var refHelperForLarge = NewReflectHelper(&largeData)

	if refHelperForSmall.NumField() != 1 {
		t.Error(`NumField is failed`)
	}
	if refHelper.NumField() != 4 {
		t.Error(`NumField is failed`)
	}
	if refHelperForLarge.NumField() != 13 {
		t.Error(`NumField is failed`)
	}

	for i := 0; i < refHelperForSmall.NumField(); i++ {
		if i == 0 {
			if refHelperForSmall.IsBool(i) {
				t.Error(`IsBool is failed`)
			}
			if refHelperForSmall.IsFloat32(i) {
				t.Error(`IsFloat32 is failed`)
			}
			if refHelperForSmall.IsFloat64(i) {
				t.Error(`IsFloat64 is failed`)
			}
			if refHelperForSmall.IsInt8(i) {
				t.Error(`IsInt8 is failed`)
			}
			if refHelperForSmall.IsInt16(i) {
				t.Error(`IsInt16 is failed`)
			}
			if !refHelperForSmall.IsInt32(i) {
				t.Error(`IsInt32 is failed`)
			}
			if refHelperForSmall.IsInt64(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelperForSmall.IsString(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelperForSmall.IsUint8(i) {
				t.Error(`IsUint8 is failed`)
			}
			if refHelperForSmall.IsUint16(i) {
				t.Error(`IsUint16 is failed`)
			}
			if refHelperForSmall.IsUint32(i) {
				t.Error(`IsUint32 is failed`)
			}
			if refHelperForSmall.IsUint64(i) {
				t.Error(`IsUint64 is failed`)
			}
		}
	}

	for i := 0; i < refHelper.NumField(); i++ {
		switch i {
		case 0:
			if refHelper.IsBool(i) {
				t.Error(`IsBool is failed`)
			}
			if refHelper.IsFloat32(i) {
				t.Error(`IsFloat32 is failed`)
			}
			if refHelper.IsFloat64(i) {
				t.Error(`IsFloat64 is failed`)
			}
			if refHelper.IsInt8(i) {
				t.Error(`IsInt8 is failed`)
			}
			if refHelper.IsInt16(i) {
				t.Error(`IsInt16 is failed`)
			}
			if !refHelper.IsInt32(i) {
				t.Error(`IsInt32 is failed`)
			}
			if refHelper.IsInt64(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelper.IsString(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelper.IsUint8(i) {
				t.Error(`IsUint8 is failed`)
			}
			if refHelper.IsUint16(i) {
				t.Error(`IsUint16 is failed`)
			}
			if refHelper.IsUint32(i) {
				t.Error(`IsUint32 is failed`)
			}
			if refHelper.IsUint64(i) {
				t.Error(`IsUint64 is failed`)
			}
			break
		case 1:
			if refHelper.IsBool(i) {
				t.Error(`IsBool is failed`)
			}
			if !refHelper.IsFloat32(i) {
				t.Error(`IsFloat32 is failed`)
			}
			if refHelper.IsFloat64(i) {
				t.Error(`IsFloat64 is failed`)
			}
			if refHelper.IsInt8(i) {
				t.Error(`IsInt8 is failed`)
			}
			if refHelper.IsInt16(i) {
				t.Error(`IsInt16 is failed`)
			}
			if refHelper.IsInt32(i) {
				t.Error(`IsInt32 is failed`)
			}
			if refHelper.IsInt64(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelper.IsString(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelper.IsUint8(i) {
				t.Error(`IsUint8 is failed`)
			}
			if refHelper.IsUint16(i) {
				t.Error(`IsUint16 is failed`)
			}
			if refHelper.IsUint32(i) {
				t.Error(`IsUint32 is failed`)
			}
			if refHelper.IsUint64(i) {
				t.Error(`IsUint64 is failed`)
			}
			break
		case 2:
			if refHelper.IsBool(i) {
				t.Error(`IsBool is failed`)
			}
			if refHelper.IsFloat32(i) {
				t.Error(`IsFloat32 is failed`)
			}
			if !refHelper.IsFloat64(i) {
				t.Error(`IsFloat64 is failed`)
			}
			if refHelper.IsInt8(i) {
				t.Error(`IsInt8 is failed`)
			}
			if refHelper.IsInt16(i) {
				t.Error(`IsInt16 is failed`)
			}
			if refHelper.IsInt32(i) {
				t.Error(`IsInt32 is failed`)
			}
			if refHelper.IsInt64(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelper.IsString(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelper.IsUint8(i) {
				t.Error(`IsUint8 is failed`)
			}
			if refHelper.IsUint16(i) {
				t.Error(`IsUint16 is failed`)
			}
			if refHelper.IsUint32(i) {
				t.Error(`IsUint32 is failed`)
			}
			if refHelper.IsUint64(i) {
				t.Error(`IsUint64 is failed`)
			}
			break
		case 3:
			if refHelper.IsBool(i) {
				t.Error(`IsBool is failed`)
			}
			if refHelper.IsFloat32(i) {
				t.Error(`IsFloat32 is failed`)
			}
			if refHelper.IsFloat64(i) {
				t.Error(`IsFloat64 is failed`)
			}
			if refHelper.IsInt8(i) {
				t.Error(`IsInt8 is failed`)
			}
			if refHelper.IsInt16(i) {
				t.Error(`IsInt16 is failed`)
			}
			if refHelper.IsInt32(i) {
				t.Error(`IsInt32 is failed`)
			}
			if refHelper.IsInt64(i) {
				t.Error(`IsInt64 is failed`)
			}
			if !refHelper.IsString(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelper.IsUint8(i) {
				t.Error(`IsUint8 is failed`)
			}
			if refHelper.IsUint16(i) {
				t.Error(`IsUint16 is failed`)
			}
			if refHelper.IsUint32(i) {
				t.Error(`IsUint32 is failed`)
			}
			if refHelper.IsUint64(i) {
				t.Error(`IsUint64 is failed`)
			}
			break
		}
	}

	for i := 0; i < refHelperForLarge.NumField(); i++ {
		switch i {
		case 0:
			if refHelperForLarge.IsBool(i) {
				t.Error(`IsBool is failed`)
			}
			if refHelperForLarge.IsFloat32(i) {
				t.Error(`IsFloat32 is failed`)
			}
			if refHelperForLarge.IsFloat64(i) {
				t.Error(`IsFloat64 is failed`)
			}
			if refHelperForLarge.IsInt8(i) {
				t.Error(`IsInt8 is failed`)
			}
			if refHelperForLarge.IsInt16(i) {
				t.Error(`IsInt16 is failed`)
			}
			if !refHelperForLarge.IsInt32(i) {
				t.Error(`IsInt32 is failed`)
			}
			if refHelperForLarge.IsInt64(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelperForLarge.IsString(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelperForLarge.IsUint8(i) {
				t.Error(`IsUint8 is failed`)
			}
			if refHelperForLarge.IsUint16(i) {
				t.Error(`IsUint16 is failed`)
			}
			if refHelperForLarge.IsUint32(i) {
				t.Error(`IsUint32 is failed`)
			}
			if refHelperForLarge.IsUint64(i) {
				t.Error(`IsUint64 is failed`)
			}
			break
		case 1:
			if refHelperForLarge.IsBool(i) {
				t.Error(`IsBool is failed`)
			}
			if !refHelperForLarge.IsFloat32(i) {
				t.Error(`IsFloat32 is failed`)
			}
			if refHelperForLarge.IsFloat64(i) {
				t.Error(`IsFloat64 is failed`)
			}
			if refHelperForLarge.IsInt8(i) {
				t.Error(`IsInt8 is failed`)
			}
			if refHelperForLarge.IsInt16(i) {
				t.Error(`IsInt16 is failed`)
			}
			if refHelperForLarge.IsInt32(i) {
				t.Error(`IsInt32 is failed`)
			}
			if refHelperForLarge.IsInt64(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelperForLarge.IsString(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelperForLarge.IsUint8(i) {
				t.Error(`IsUint8 is failed`)
			}
			if refHelperForLarge.IsUint16(i) {
				t.Error(`IsUint16 is failed`)
			}
			if refHelperForLarge.IsUint32(i) {
				t.Error(`IsUint32 is failed`)
			}
			if refHelperForLarge.IsUint64(i) {
				t.Error(`IsUint64 is failed`)
			}
			break
		case 2:
			if refHelperForLarge.IsBool(i) {
				t.Error(`IsBool is failed`)
			}
			if refHelperForLarge.IsFloat32(i) {
				t.Error(`IsFloat32 is failed`)
			}
			if !refHelperForLarge.IsFloat64(i) {
				t.Error(`IsFloat64 is failed`)
			}
			if refHelperForLarge.IsInt8(i) {
				t.Error(`IsInt8 is failed`)
			}
			if refHelperForLarge.IsInt16(i) {
				t.Error(`IsInt16 is failed`)
			}
			if refHelperForLarge.IsInt32(i) {
				t.Error(`IsInt32 is failed`)
			}
			if refHelperForLarge.IsInt64(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelperForLarge.IsString(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelperForLarge.IsUint8(i) {
				t.Error(`IsUint8 is failed`)
			}
			if refHelperForLarge.IsUint16(i) {
				t.Error(`IsUint16 is failed`)
			}
			if refHelperForLarge.IsUint32(i) {
				t.Error(`IsUint32 is failed`)
			}
			if refHelperForLarge.IsUint64(i) {
				t.Error(`IsUint64 is failed`)
			}
			break
		case 3:
			if refHelperForLarge.IsBool(i) {
				t.Error(`IsBool is failed`)
			}
			if refHelperForLarge.IsFloat32(i) {
				t.Error(`IsFloat32 is failed`)
			}
			if refHelperForLarge.IsFloat64(i) {
				t.Error(`IsFloat64 is failed`)
			}
			if refHelperForLarge.IsInt8(i) {
				t.Error(`IsInt8 is failed`)
			}
			if refHelperForLarge.IsInt16(i) {
				t.Error(`IsInt16 is failed`)
			}
			if refHelperForLarge.IsInt32(i) {
				t.Error(`IsInt32 is failed`)
			}
			if refHelperForLarge.IsInt64(i) {
				t.Error(`IsInt64 is failed`)
			}
			if !refHelperForLarge.IsString(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelperForLarge.IsUint8(i) {
				t.Error(`IsUint8 is failed`)
			}
			if refHelperForLarge.IsUint16(i) {
				t.Error(`IsUint16 is failed`)
			}
			if refHelperForLarge.IsUint32(i) {
				t.Error(`IsUint32 is failed`)
			}
			if refHelperForLarge.IsUint64(i) {
				t.Error(`IsUint64 is failed`)
			}
			break
		case 4:
			if !refHelperForLarge.IsBool(i) {
				t.Error(`IsBool is failed`)
			}
			if refHelperForLarge.IsFloat32(i) {
				t.Error(`IsFloat32 is failed`)
			}
			if refHelperForLarge.IsFloat64(i) {
				t.Error(`IsFloat64 is failed`)
			}
			if refHelperForLarge.IsInt8(i) {
				t.Error(`IsInt8 is failed`)
			}
			if refHelperForLarge.IsInt16(i) {
				t.Error(`IsInt16 is failed`)
			}
			if refHelperForLarge.IsInt32(i) {
				t.Error(`IsInt32 is failed`)
			}
			if refHelperForLarge.IsInt64(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelperForLarge.IsString(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelperForLarge.IsUint8(i) {
				t.Error(`IsUint8 is failed`)
			}
			if refHelperForLarge.IsUint16(i) {
				t.Error(`IsUint16 is failed`)
			}
			if refHelperForLarge.IsUint32(i) {
				t.Error(`IsUint32 is failed`)
			}
			if refHelperForLarge.IsUint64(i) {
				t.Error(`IsUint64 is failed`)
			}
			break
		case 5:
			if refHelperForLarge.IsBool(i) {
				t.Error(`IsBool is failed`)
			}
			if refHelperForLarge.IsFloat32(i) {
				t.Error(`IsFloat32 is failed`)
			}
			if refHelperForLarge.IsFloat64(i) {
				t.Error(`IsFloat64 is failed`)
			}
			if !refHelperForLarge.IsInt8(i) {
				t.Error(`IsInt8 is failed`)
			}
			if refHelperForLarge.IsInt16(i) {
				t.Error(`IsInt16 is failed`)
			}
			if refHelperForLarge.IsInt32(i) {
				t.Error(`IsInt32 is failed`)
			}
			if refHelperForLarge.IsInt64(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelperForLarge.IsString(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelperForLarge.IsUint8(i) {
				t.Error(`IsUint8 is failed`)
			}
			if refHelperForLarge.IsUint16(i) {
				t.Error(`IsUint16 is failed`)
			}
			if refHelperForLarge.IsUint32(i) {
				t.Error(`IsUint32 is failed`)
			}
			if refHelperForLarge.IsUint64(i) {
				t.Error(`IsUint64 is failed`)
			}
			break
		case 6:
			if refHelperForLarge.IsBool(i) {
				t.Error(`IsBool is failed`)
			}
			if refHelperForLarge.IsFloat32(i) {
				t.Error(`IsFloat32 is failed`)
			}
			if refHelperForLarge.IsFloat64(i) {
				t.Error(`IsFloat64 is failed`)
			}
			if refHelperForLarge.IsInt8(i) {
				t.Error(`IsInt8 is failed`)
			}
			if !refHelperForLarge.IsInt16(i) {
				t.Error(`IsInt16 is failed`)
			}
			if refHelperForLarge.IsInt32(i) {
				t.Error(`IsInt32 is failed`)
			}
			if refHelperForLarge.IsInt64(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelperForLarge.IsString(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelperForLarge.IsUint8(i) {
				t.Error(`IsUint8 is failed`)
			}
			if refHelperForLarge.IsUint16(i) {
				t.Error(`IsUint16 is failed`)
			}
			if refHelperForLarge.IsUint32(i) {
				t.Error(`IsUint32 is failed`)
			}
			if refHelperForLarge.IsUint64(i) {
				t.Error(`IsUint64 is failed`)
			}
			break
		case 7:
			if refHelperForLarge.IsBool(i) {
				t.Error(`IsBool is failed`)
			}
			if refHelperForLarge.IsFloat32(i) {
				t.Error(`IsFloat32 is failed`)
			}
			if refHelperForLarge.IsFloat64(i) {
				t.Error(`IsFloat64 is failed`)
			}
			if refHelperForLarge.IsInt8(i) {
				t.Error(`IsInt8 is failed`)
			}
			if refHelperForLarge.IsInt16(i) {
				t.Error(`IsInt16 is failed`)
			}
			if refHelperForLarge.IsInt32(i) {
				t.Error(`IsInt32 is failed`)
			}
			if !refHelperForLarge.IsInt64(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelperForLarge.IsString(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelperForLarge.IsUint8(i) {
				t.Error(`IsUint8 is failed`)
			}
			if refHelperForLarge.IsUint16(i) {
				t.Error(`IsUint16 is failed`)
			}
			if refHelperForLarge.IsUint32(i) {
				t.Error(`IsUint32 is failed`)
			}
			if refHelperForLarge.IsUint64(i) {
				t.Error(`IsUint64 is failed`)
			}
			break
		case 8:
			if refHelperForLarge.IsBool(i) {
				t.Error(`IsBool is failed`)
			}
			if refHelperForLarge.IsFloat32(i) {
				t.Error(`IsFloat32 is failed`)
			}
			if refHelperForLarge.IsFloat64(i) {
				t.Error(`IsFloat64 is failed`)
			}
			if refHelperForLarge.IsInt8(i) {
				t.Error(`IsInt8 is failed`)
			}
			if refHelperForLarge.IsInt16(i) {
				t.Error(`IsInt16 is failed`)
			}
			if refHelperForLarge.IsInt32(i) {
				t.Error(`IsInt32 is failed`)
			}
			if refHelperForLarge.IsInt64(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelperForLarge.IsString(i) {
				t.Error(`IsInt64 is failed`)
			}
			if !refHelperForLarge.IsUint8(i) {
				t.Error(`IsUint8 is failed`)
			}
			if refHelperForLarge.IsUint16(i) {
				t.Error(`IsUint16 is failed`)
			}
			if refHelperForLarge.IsUint32(i) {
				t.Error(`IsUint32 is failed`)
			}
			if refHelperForLarge.IsUint64(i) {
				t.Error(`IsUint64 is failed`)
			}
			break
		case 9:
			if refHelperForLarge.IsBool(i) {
				t.Error(`IsBool is failed`)
			}
			if refHelperForLarge.IsFloat32(i) {
				t.Error(`IsFloat32 is failed`)
			}
			if refHelperForLarge.IsFloat64(i) {
				t.Error(`IsFloat64 is failed`)
			}
			if refHelperForLarge.IsInt8(i) {
				t.Error(`IsInt8 is failed`)
			}
			if refHelperForLarge.IsInt16(i) {
				t.Error(`IsInt16 is failed`)
			}
			if refHelperForLarge.IsInt32(i) {
				t.Error(`IsInt32 is failed`)
			}
			if refHelperForLarge.IsInt64(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelperForLarge.IsString(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelperForLarge.IsUint8(i) {
				t.Error(`IsUint8 is failed`)
			}
			if !refHelperForLarge.IsUint16(i) {
				t.Error(`IsUint16 is failed`)
			}
			if refHelperForLarge.IsUint32(i) {
				t.Error(`IsUint32 is failed`)
			}
			if refHelperForLarge.IsUint64(i) {
				t.Error(`IsUint64 is failed`)
			}
			break
		case 10:
			if refHelperForLarge.IsBool(i) {
				t.Error(`IsBool is failed`)
			}
			if refHelperForLarge.IsFloat32(i) {
				t.Error(`IsFloat32 is failed`)
			}
			if refHelperForLarge.IsFloat64(i) {
				t.Error(`IsFloat64 is failed`)
			}
			if refHelperForLarge.IsInt8(i) {
				t.Error(`IsInt8 is failed`)
			}
			if refHelperForLarge.IsInt16(i) {
				t.Error(`IsInt16 is failed`)
			}
			if refHelperForLarge.IsInt32(i) {
				t.Error(`IsInt32 is failed`)
			}
			if refHelperForLarge.IsInt64(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelperForLarge.IsString(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelperForLarge.IsUint8(i) {
				t.Error(`IsUint8 is failed`)
			}
			if refHelperForLarge.IsUint16(i) {
				t.Error(`IsUint16 is failed`)
			}
			if !refHelperForLarge.IsUint32(i) {
				t.Error(`IsUint32 is failed`)
			}
			if refHelperForLarge.IsUint64(i) {
				t.Error(`IsUint64 is failed`)
			}
			break
		case 11:
			if refHelperForLarge.IsBool(i) {
				t.Error(`IsBool is failed`)
			}
			if refHelperForLarge.IsFloat32(i) {
				t.Error(`IsFloat32 is failed`)
			}
			if refHelperForLarge.IsFloat64(i) {
				t.Error(`IsFloat64 is failed`)
			}
			if refHelperForLarge.IsInt8(i) {
				t.Error(`IsInt8 is failed`)
			}
			if refHelperForLarge.IsInt16(i) {
				t.Error(`IsInt16 is failed`)
			}
			if refHelperForLarge.IsInt32(i) {
				t.Error(`IsInt32 is failed`)
			}
			if refHelperForLarge.IsInt64(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelperForLarge.IsString(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelperForLarge.IsUint8(i) {
				t.Error(`IsUint8 is failed`)
			}
			if refHelperForLarge.IsUint16(i) {
				t.Error(`IsUint16 is failed`)
			}
			if refHelperForLarge.IsUint32(i) {
				t.Error(`IsUint32 is failed`)
			}
			if !refHelperForLarge.IsUint64(i) {
				t.Error(`IsUint64 is failed`)
			}
			break
		case 12:
			if refHelperForLarge.IsBool(i) {
				t.Error(`IsBool is failed`)
			}
			if refHelperForLarge.IsFloat32(i) {
				t.Error(`IsFloat32 is failed`)
			}
			if refHelperForLarge.IsFloat64(i) {
				t.Error(`IsFloat64 is failed`)
			}
			if refHelperForLarge.IsInt8(i) {
				t.Error(`IsInt8 is failed`)
			}
			if refHelperForLarge.IsInt16(i) {
				t.Error(`IsInt16 is failed`)
			}
			if refHelperForLarge.IsInt32(i) {
				t.Error(`IsInt32 is failed`)
			}
			if refHelperForLarge.IsInt64(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelperForLarge.IsString(i) {
				t.Error(`IsInt64 is failed`)
			}
			if refHelperForLarge.IsUint8(i) {
				t.Error(`IsUint8 is failed`)
			}
			if refHelperForLarge.IsUint16(i) {
				t.Error(`IsUint16 is failed`)
			}
			if refHelperForLarge.IsUint32(i) {
				t.Error(`IsUint32 is failed`)
			}
			if refHelperForLarge.IsUint64(i) {
				t.Error(`IsUint64 is failed`)
			}
			break
		}
	}

	// Method Call
	results := refHelperForSmall.Call("FuncA", []reflect.Value{reflect.ValueOf(9999), reflect.ValueOf(`kore`)})
	if len(results) != 2 {
		t.Error(`result length not matched`)
	} else if results[0].Int() != 9999 {
		t.Error(`result[0] not matched`)
	} else {
		if callError, castOK := results[1].Interface().(error); !castOK {
			t.Error(`result[1] not error`)
		} else if strings.Compare(callError.Error(), `kore`) != 0 {
			t.Error(`result[1] not matched: kore`)
		}
	}
}
