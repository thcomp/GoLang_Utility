package utility

import (
	"reflect"
	"strconv"
	"strings"
	"time"
)

type ReflectHelper struct {
	mTargetValue     reflect.Value
	mTargetInterface interface{}
}

func NewReflectHelper(targetInterface interface{}) *ReflectHelper {
	var ret ReflectHelper

	ret.mTargetValue = reflect.Indirect(reflect.ValueOf(targetInterface))
	ret.mTargetInterface = targetInterface
	return &ret
}

func (this *ReflectHelper) NumField() int {
	return this.mTargetValue.NumField()
}

func (this *ReflectHelper) InterfaceName() string {
	return reflect.TypeOf(this.mTargetInterface).Name()
}

func (this *ReflectHelper) FieldName(index int) string {
	return this.mTargetValue.Type().Field(index).Name
}

func (this *ReflectHelper) FieldIndex(name string) int {
	ret := -1

	for i := 0; i < this.mTargetValue.NumField(); i++ {
		if this.mTargetValue.Type().Field(i).Name == name {
			ret = i
			break
		}
	}

	return ret
}

func (this *ReflectHelper) Tag(index int) reflect.StructTag {
	if reflect.TypeOf(this.mTargetInterface).Kind() == reflect.Ptr {
		return reflect.TypeOf(this.mTargetInterface).Elem().Field(index).Tag
	} else {
		return reflect.TypeOf(this.mTargetValue).Field(index).Tag
	}
}

func (this *ReflectHelper) BoolValue(index int) bool {
	return this.mTargetValue.Field(index).Bool()
}

func (this *ReflectHelper) StringValue(index int) string {
	return this.mTargetValue.Field(index).String()
}

func (this *ReflectHelper) String(index int) string {
	var ret string
	var value reflect.Value = this.mTargetValue.Field(index)

	if this.IsBool(index) {
		ret = strconv.FormatBool(value.Bool())
	} else if this.IsInt32(index) || this.IsInt64(index) {
		ret = strconv.FormatInt(value.Int(), 10)
	} else if this.IsUint32(index) || this.IsUint64(index) {
		ret = strconv.FormatUint(value.Uint(), 10)
	} else if this.IsFloat32(index) {
		ret = strconv.FormatFloat(value.Float(), 'e', -1, 32)
	} else if this.IsFloat64(index) {
		ret = strconv.FormatFloat(value.Float(), 'e', -1, 64)
	}

	return ret
}

func (this *ReflectHelper) IntValue(index int) int64 {
	return this.mTargetValue.Field(index).Int()
}

func (this *ReflectHelper) UintValue(index int) uint64 {
	return this.mTargetValue.Field(index).Uint()
}

func (this *ReflectHelper) FloatValue(index int) float64 {
	return this.mTargetValue.Field(index).Float()
}

func (this *ReflectHelper) TimeValue(index int) time.Time {
	fieldInf := this.mTargetValue.Field(index).Interface()
	ret, assertionOK := fieldInf.(time.Time)

	if !assertionOK {
		LogfE("fail to assertion: %v -> time.Time", fieldInf)
	}

	return ret
}

func (this *ReflectHelper) ValueKind(index int) reflect.Kind {
	return this.mTargetValue.Field(index).Kind()
}

func (this *ReflectHelper) IsBool(index int) bool {
	return this.mTargetValue.Field(index).Kind() == reflect.Bool
}

func (this *ReflectHelper) IsString(index int) bool {
	return this.mTargetValue.Field(index).Kind() == reflect.String
}

func (this *ReflectHelper) IsInt8(index int) bool {
	return this.mTargetValue.Field(index).Kind() == reflect.Int8
}

func (this *ReflectHelper) IsInt16(index int) bool {
	return this.mTargetValue.Field(index).Kind() == reflect.Int16
}

func (this *ReflectHelper) IsInt32(index int) bool {
	return this.mTargetValue.Field(index).Kind() == reflect.Int32
}

func (this *ReflectHelper) IsInt64(index int) bool {
	return this.mTargetValue.Field(index).Kind() == reflect.Int64
}

func (this *ReflectHelper) IsUint8(index int) bool {
	return this.mTargetValue.Field(index).Kind() == reflect.Uint8
}

func (this *ReflectHelper) IsUint16(index int) bool {
	return this.mTargetValue.Field(index).Kind() == reflect.Uint16
}

func (this *ReflectHelper) IsUint32(index int) bool {
	return this.mTargetValue.Field(index).Kind() == reflect.Uint32
}

func (this *ReflectHelper) IsUint64(index int) bool {
	return this.mTargetValue.Field(index).Kind() == reflect.Uint64
}

func (this *ReflectHelper) IsFloat32(index int) bool {
	return this.mTargetValue.Field(index).Kind() == reflect.Float32
}

func (this *ReflectHelper) IsFloat64(index int) bool {
	return this.mTargetValue.Field(index).Kind() == reflect.Float64
}

func (this *ReflectHelper) IsFloat(index int) bool {
	return this.mTargetValue.Field(index).Kind() == reflect.Float32 || this.mTargetValue.Field(index).Kind() == reflect.Float64
}

func (this *ReflectHelper) IsInt(index int) bool {
	ret := false

	switch this.mTargetValue.Field(index).Kind() {
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		ret = true
		break
	}

	return ret
}

func (this *ReflectHelper) IsUint(index int) bool {
	ret := false

	switch this.mTargetValue.Field(index).Kind() {
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		ret = true
		break
	}

	return ret
}

func (this *ReflectHelper) IsIntUint(index int) bool {
	return this.IsInt(index) || this.IsUint(index)
}

func (this *ReflectHelper) IsNumber(index int) bool {
	return this.IsInt(index) || this.IsUint(index) || this.IsFloat(index)
}

func (this *ReflectHelper) IsTime(index int) bool {
	ret := false

	if this.mTargetValue.Field(index).Kind() == reflect.Struct {
		_, ret = this.mTargetValue.Field(index).Interface().(time.Time)
	}

	return ret
}

func (this *ReflectHelper) SetByName(name string, value interface{}) bool {
	var ret bool = false
	var tempValue reflect.Value = this.mTargetValue.FieldByName(name)

	if tempValue.Kind() == reflect.ValueOf(value).Kind() {
		tempValue.Set(reflect.ValueOf(value))
		ret = true
	}

	return ret
}

func (this *ReflectHelper) SetByIndex(index int, valueInf interface{}) bool {
	var ret bool = false
	var tempValue reflect.Value = this.mTargetValue.Field(index).Addr().Elem()

	if tempValue.CanSet() {
		v := reflect.ValueOf(valueInf)
		if v.Kind() == reflect.Ptr {
			v = reflect.Indirect(v)
		}

		if tempValue.Kind() == v.Kind() {
			tempValue.Set(v)
			ret = true
		}
	}

	return ret
}

func (this *ReflectHelper) GetByName(name string) interface{} {
	var ret interface{} = nil

	tempValue := this.mTargetValue.FieldByName(name)
	if tempValue.IsValid() {
		ret = tempValue.Interface()
	}

	return ret
}

func (this *ReflectHelper) GetByTagName(key string, name string) interface{} {
	var ret interface{} = nil

	for i := 0; i < this.NumField(); i++ {
		tagData := this.Tag(i)
		if tempName, exist := tagData.Lookup(key); exist {
			if strings.Compare(name, tempName) == 0 {
				ret = this.GetByIndex(i)
				break
			}
		}
	}

	return ret
}

func (this *ReflectHelper) GetByIndex(index int) interface{} {
	return this.mTargetValue.Field(index).Interface()
}

func (this *ReflectHelper) GetValueByIndex(index int) reflect.Value {
	return this.mTargetValue.Field(index)
}

func (this *ReflectHelper) GetPtrByIndex(index int) reflect.Value {
	return this.mTargetValue.Field(index).Addr().Elem()
}

func (this *ReflectHelper) GetAddrByIndex(index int) reflect.Value {
	return this.mTargetValue.Field(index).Addr()
}

func (this *ReflectHelper) Call(methodName string, values []reflect.Value) []reflect.Value {
	return reflect.ValueOf(this.mTargetInterface).MethodByName(methodName).Call(values)
}

func (this *ReflectHelper) CanSet(name string, value interface{}) bool {
	var ret bool = false
	var tempValue reflect.Value = this.mTargetValue.FieldByName(name)

	if tempValue.CanSet() && (tempValue.Type() == reflect.TypeOf(value)) {
		ret = true
	}

	return ret
}
