package utility

import (
	//"strconv"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"
	//"time"
)

type SmallData struct {
	Value0 int32
}

type Data struct {
	Value0 int32
	Value1 float32 `json:"value1"`
	Value2 float64
	Value3 string
	Value4 time.Time
}

type LargeData struct {
	Value0  int32     `csv:"value0"`
	Value1  float32   `csv:"value1"`
	Value2  float64   `csv:"value2"`
	Value3  string    `csv:"value3"`
	Value4  bool      `csv:"value4"`
	Value5  int8      `csv:"value5"`
	Value6  int16     `csv:"value6"`
	Value7  int64     `csv:"value7"`
	Value8  uint8     `csv:"value8"`
	Value9  uint16    `csv:"value9"`
	Value10 uint32    `csv:"value10"`
	Value11 uint64    `csv:"value11"`
	Value12 SmallData `csv:"value12"`
	Value13 time.Time `csv:"value13"`
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
	if refHelper.NumField() != 5 {
		t.Error(`NumField is failed`)
	}
	if refHelperForLarge.NumField() != 14 {
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
				t.Error(`IsString is failed`)
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
				t.Error(`IsString is failed`)
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
			if refHelper.IsTime(i) {
				t.Error(`IsTime is failed`)
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
				t.Error(`IsString is failed`)
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
			if refHelper.IsTime(i) {
				t.Error(`IsTime is failed`)
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
				t.Error(`IsString is failed`)
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
			if refHelper.IsTime(i) {
				t.Error(`IsTime is failed`)
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
				t.Error(`IsString is failed`)
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
			if refHelper.IsTime(i) {
				t.Error(`IsTime is failed`)
			}
			break
		case 4:
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
			if refHelper.IsString(i) {
				t.Error(`IsString is failed`)
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
			if !refHelper.IsTime(i) {
				t.Error(`IsTime is failed`)
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
				t.Error(`IsString is failed`)
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
			if refHelperForLarge.IsTime(i) {
				t.Error(`IsTime is failed`)
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
				t.Error(`IsString is failed`)
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
			if refHelperForLarge.IsTime(i) {
				t.Error(`IsTime is failed`)
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
				t.Error(`IsString is failed`)
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
			if refHelperForLarge.IsTime(i) {
				t.Error(`IsTime is failed`)
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
				t.Error(`IsString is failed`)
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
			if refHelperForLarge.IsTime(i) {
				t.Error(`IsTime is failed`)
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
				t.Error(`IsString is failed`)
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
			if refHelperForLarge.IsTime(i) {
				t.Error(`IsTime is failed`)
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
				t.Error(`IsString is failed`)
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
			if refHelperForLarge.IsTime(i) {
				t.Error(`IsTime is failed`)
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
				t.Error(`IsString is failed`)
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
			if refHelperForLarge.IsTime(i) {
				t.Error(`IsTime is failed`)
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
				t.Error(`IsString is failed`)
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
			if refHelperForLarge.IsTime(i) {
				t.Error(`IsTime is failed`)
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
				t.Error(`IsString is failed`)
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
			if refHelperForLarge.IsTime(i) {
				t.Error(`IsTime is failed`)
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
				t.Error(`IsString is failed`)
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
			if refHelperForLarge.IsTime(i) {
				t.Error(`IsTime is failed`)
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
				t.Error(`IsString is failed`)
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
			if refHelperForLarge.IsTime(i) {
				t.Error(`IsTime is failed`)
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
				t.Error(`IsString is failed`)
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
			if refHelperForLarge.IsTime(i) {
				t.Error(`IsTime is failed`)
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
				t.Error(`IsString is failed`)
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
			if refHelperForLarge.IsTime(i) {
				t.Error(`IsTime is failed`)
			}
			break
		case 13:
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
				t.Error(`IsString is failed`)
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
			if !refHelperForLarge.IsTime(i) {
				t.Error(`IsTime is failed`)
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

	// get field by tag name
	targetField := refHelper.GetByTagName(`json`, `value1`)
	if targetField == nil {
		t.Error(`not found field by GetByTagName`)
	}
}

func Test_SetByTagName1(t *testing.T) {
	data := LargeData{}
	reflectHelper := NewReflectHelper(&data)
	value13, _ := DateInJST(2024, 2, 19, 16, 24, 35, 1)
	values := []interface{}{
		0,
		1,
		2,
		"3",
		true,
		5,
		6,
		7,
		8,
		9,
		10,
		11,
		&SmallData{
			Value0: 12,
		},
		value13,
	}

	for pos, value := range values {
		name := fmt.Sprintf("value%d", pos)
		if !reflectHelper.SetByTagName("csv", name, value) && pos != 12 {
			t.Fatalf("fail to set by tag name: %s, %v", name, value)
		}
	}

	reflectHelper2 := NewReflectHelper(&data)
	for pos, value := range values {
		name := fmt.Sprintf("value%d", pos)
		value2 := reflectHelper2.GetByTagName("csv", name)

		infHelper := NewInterfaceHelper(value)
		if !infHelper.EqualWithRoughClassification(value2) && pos != 12 {
			t.Fatalf("not matched value: %d, %v, %v", pos, value, value2)
		}
	}
}

func Test_SetOnList1(t *testing.T) {
	values := []int{
		0,
		1,
		2,
		3,
	}
	fromValues := []int{
		10,
		11,
		12,
		13,
	}

	reflectHelper := NewReflectHelper(&values)
	for pos, fromValue := range fromValues {
		if !reflectHelper.SetOnList(pos, fromValue) {
			t.Fatalf("fail to set on list: %d, %d", pos, fromValue)
		}
	}

	for pos := range values {
		if values[pos] != fromValues[pos] {
			t.Fatalf("not matched value: %d, %v, %v", pos, values[pos], fromValues[pos])
		}
	}
}

func Test_SetOnList2(t *testing.T) {
	values := []int{
		0,
		1,
		2,
		3,
	}
	fromValues := []int{
		10,
		11,
		12,
		13,
	}

	reflectHelper := NewReflectHelper(&values)
	for pos, fromValue := range fromValues {
		if !reflectHelper.SetOnList(len(values), fromValue) {
			t.Fatalf("fail to append on list: %d, %d", pos, fromValue)
		}
	}

	for pos := range fromValues {
		if values[len(values)-(len(fromValues)-pos)] != fromValues[pos] {
			t.Fatalf("not matched value: %d, %v, %v", pos, values[len(values)-(len(fromValues)-pos)], fromValues[pos])
		}
	}
}

func Test_SetOnList3(t *testing.T) {
	values := []string{
		"0",
		"1",
		"2",
		"3",
	}
	fromValues := []string{
		"10",
		"11",
		"12",
		"13",
	}

	reflectHelper := NewReflectHelper(&values)
	for pos, fromValue := range fromValues {
		if !reflectHelper.SetOnList(pos, fromValue) {
			t.Fatalf("fail to set on list: %d, %v", pos, fromValue)
		}
	}

	for pos := range values {
		if values[pos] != fromValues[pos] {
			t.Fatalf("not matched value: %d, %v, %v", pos, values[pos], fromValues[pos])
		}
	}
}

func Test_SetOnList4(t *testing.T) {
	values := []string{
		"0",
		"1",
		"2",
		"3",
	}
	fromValues := []string{
		"10",
		"11",
		"12",
		"13",
	}

	reflectHelper := NewReflectHelper(&values)
	for pos, fromValue := range fromValues {
		if !reflectHelper.SetOnList(len(values), fromValue) {
			t.Fatalf("fail to append on list: %d, %v", pos, fromValue)
		}
	}

	for pos := range fromValues {
		if values[len(values)-(len(fromValues)-pos)] != fromValues[pos] {
			t.Fatalf("not matched value: %d, %v, %v", pos, values[len(values)-(len(fromValues)-pos)], fromValues[pos])
		}
	}
}

func Test_SetOnList5(t *testing.T) {
	values := []interface{}{
		"0",
		"1",
		"2",
		"3",
	}
	fromValues := []interface{}{
		10,
		11,
		12,
		13,
	}

	reflectHelper := NewReflectHelper(&values)
	for pos, fromValue := range fromValues {
		if !reflectHelper.SetOnList(pos, fromValue) {
			t.Fatalf("fail to set on list: %d, %v", pos, fromValue)
		}
	}

	for pos := range values {
		if values[pos] != fromValues[pos] {
			t.Fatalf("not matched value: %d, %v, %v", pos, values[pos], fromValues[pos])
		}
	}
}

func Test_SetOnList6(t *testing.T) {
	values := []interface{}{
		"0",
		"1",
		"2",
		"3",
	}
	fromValues := []interface{}{
		10,
		11,
		12,
		13,
	}

	reflectHelper := NewReflectHelper(&values)
	for pos, fromValue := range fromValues {
		if !reflectHelper.SetOnList(len(values), fromValue) {
			t.Fatalf("fail to append on list: %d, %v", pos, fromValue)
		}
	}

	for pos := range fromValues {
		if values[len(values)-(len(fromValues)-pos)] != fromValues[pos] {
			t.Fatalf("not matched value: %d, %v, %v", pos, values[len(values)-(len(fromValues)-pos)], fromValues[pos])
		}
	}
}

func test_SetOnList7(t *testing.T) { /* not support */
	values := []int{
		0,
		1,
		2,
		3,
	}
	fromValues := []float64{
		10,
		11,
		12,
		13,
	}

	reflectHelper := NewReflectHelper(&values)
	for pos, fromValue := range fromValues {
		if !reflectHelper.SetOnList(pos, fromValue) {
			t.Fatalf("fail to set on list: %d, %v", pos, fromValue)
		}
	}

	for pos := range values {
		if values[pos] != int(fromValues[pos]) {
			t.Fatalf("not matched value: %d, %v, %v", pos, values[pos], fromValues[pos])
		}
	}
}

func test_SetOnList8(t *testing.T) { /* not support */
	values := []int{
		0,
		1,
		2,
		3,
	}
	fromValues := []float64{
		10,
		11,
		12,
		13,
	}

	reflectHelper := NewReflectHelper(&values)
	for pos, fromValue := range fromValues {
		if !reflectHelper.SetOnList(len(values), fromValue) {
			t.Fatalf("fail to append on list: %d, %v", pos, fromValue)
		}
	}

	for pos := range fromValues {
		if values[len(values)-(len(fromValues)-pos)] != int(fromValues[pos]) {
			t.Fatalf("not matched value: %d, %v, %v", pos, values[len(values)-(len(fromValues)-pos)], fromValues[pos])
		}
	}
}

func Test_SetOnMap1(t *testing.T) {
	values := map[int]string{
		0: "0",
		1: "1",
		2: "2",
		3: "3",
	}
	fromValues := []string{
		"10",
		"11",
		"12",
		"13",
	}

	reflectHelper := NewReflectHelper(&values)
	for pos, fromValue := range fromValues {
		if !reflectHelper.SetOnMap(pos, fromValue) {
			t.Fatalf("fail to set on list: %d, %v", pos, fromValue)
		}
	}

	for pos := range values {
		if values[pos] != fromValues[pos] {
			t.Fatalf("not matched value: %d, %v, %v", pos, values[pos], fromValues[pos])
		}
	}
}

func Test_SetOnMap2(t *testing.T) {
	values := map[int]string{
		0: "0",
		1: "1",
		2: "2",
		3: "3",
	}
	fromValues := []string{
		"0",
		"1",
		"2",
		"3",
		"10",
		"11",
		"12",
		"13",
	}

	reflectHelper := NewReflectHelper(&values)
	for pos, fromValue := range fromValues {
		if !reflectHelper.SetOnMap(pos, fromValue) {
			t.Fatalf("fail to set on list: %d, %v", pos, fromValue)
		}
	}

	for pos := range values {
		if values[pos] != fromValues[pos] {
			t.Fatalf("not matched value: %d, %v, %v", pos, values[pos], fromValues[pos])
		}
	}
}
