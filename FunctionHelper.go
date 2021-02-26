package utility

import (
	"math/rand"
	"reflect"
	"time"
)

func CallAfterDelay(delayTime time.Duration, funcValue reflect.Value, parameters ...interface{}) []reflect.Value {
	time.Sleep(delayTime)

	paramValueSlice := []reflect.Value(nil)

	if parameters != nil && len(parameters) > 0 {
		paramValueSlice = []reflect.Value{}
		for _, parameter := range parameters {
			paramValueSlice = append(paramValueSlice, reflect.ValueOf(parameter))
		}
	}

	return funcValue.Call(paramValueSlice)
}

func CallAfterRandomDelay(funcValue reflect.Value, parameters ...interface{}) []reflect.Value {
	delayTime := time.Duration(rand.New(rand.NewSource(time.Now().UnixNano())).Int31() % 1000)
	return CallAfterDelay(delayTime, funcValue, parameters...)
}
