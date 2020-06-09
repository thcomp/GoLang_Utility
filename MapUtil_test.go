package utility

import (
	"reflect"
	"testing"
)

func TestMapKeys(t *testing.T) {
	testMap1 := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	testMap2 := map[int]int{
		1: 3,
		2: 4,
	}
	testMap3 := map[string]int{}

	key1Inf := MapKeys(testMap1)
	key2Inf := MapKeys(testMap2)
	key3Inf := MapKeys(testMap3)

	if _, assertOK := key1Inf.([]string); !assertOK {
		t.Errorf("not matched type: []string vs %s", reflect.TypeOf(key1Inf).Kind().String())
	}
	if _, assertOK := key2Inf.([]int); !assertOK {
		t.Errorf("not matched type: []int vs %s", reflect.TypeOf(key2Inf).Kind().String())
	}
	if key3Inf != nil {
		t.Errorf("not nil")
	}
}

func TestGetInterface(t *testing.T) {
	testMapArray := []interface{}{
		map[string]string{
			"1": "test_value1",
			"2": "test_value2",
			"3": "test_value3",
			"4": "test_value4",
			"5": "test_value5",
		},
		map[int]string{
			1: "test_value1",
			2: "test_value2",
			3: "test_value3",
			4: "test_value4",
			5: "test_value5",
		},
		map[float32]string{
			float32(1): "test_value1",
			float32(2): "test_value2",
			float32(3): "test_value3",
			float32(4): "test_value4",
			float32(5): "test_value5",
		},
		map[float64]string{
			float64(1): "test_value1",
			float64(2): "test_value2",
			float64(3): "test_value3",
			float64(4): "test_value4",
			float64(5): "test_value5",
		},
	}
	keyArray := []interface{}{
		"1",
		int(1),
		float32(1),
		float64(1),
	}
	notFoundKeyArray := []interface{}{
		"10",
		int(10),
		float32(10),
		float64(10),
	}
	notMatchedKeyArray := []interface{}{
		int(10),
		float32(10),
		float64(10),
		"10",
	}

	for index, mapObject := range testMapArray {
		if _, getErr := GetInterface(mapObject, keyArray[index]); getErr != nil {
			t.Fatalf(getErr.Error())
		}
	}

	for index, mapObject := range testMapArray {
		if _, getErr := GetInterface(mapObject, notFoundKeyArray[index]); getErr == nil {
			t.Fatalf("it should be error")
		} else {
			if !IsNotFound(getErr) {
				t.Fatalf("it should be error for NotFound: %s", getErr.Error())
			}
		}
	}

	for index, mapObject := range testMapArray {
		if _, getErr := GetInterface(mapObject, notMatchedKeyArray[index]); getErr == nil {
			t.Fatalf("it should be error")
		} else {
			if IsNotFound(getErr) {
				t.Fatalf("it should not be not found error: %s", getErr.Error())
			}
		}
	}
}

func TestGetString(t *testing.T) {
	testMapArray := []interface{}{
		map[string]string{
			"1": "test_value1",
			"2": "test_value2",
			"3": "test_value3",
			"4": "test_value4",
			"5": "test_value5",
		},
		map[int]string{
			1: "test_value1",
			2: "test_value2",
			3: "test_value3",
			4: "test_value4",
			5: "test_value5",
		},
		map[float32]string{
			float32(1): "test_value1",
			float32(2): "test_value2",
			float32(3): "test_value3",
			float32(4): "test_value4",
			float32(5): "test_value5",
		},
		map[float64]string{
			float64(1): "test_value1",
			float64(2): "test_value2",
			float64(3): "test_value3",
			float64(4): "test_value4",
			float64(5): "test_value5",
		},
	}
	keyArray := []interface{}{
		"1",
		int(1),
		float32(1),
		float64(1),
	}

	for index, mapObject := range testMapArray {
		if stringValue, getErr := GetString(mapObject, keyArray[index]); getErr != nil {
			t.Fatalf(getErr.Error())
		} else if stringValue != "test_value1" {
			t.Fatalf("not matched %s vs %s", stringValue, "test_value1")
		}
	}
}

func TestGetNumber(t *testing.T) {
	testMapArray := []interface{}{
		map[float64]interface{}{
			float64(1):  float64(1),
			float64(2):  float32(2),
			float64(3):  int64(3),
			float64(4):  uint64(4),
			float64(5):  int32(5),
			float64(6):  uint32(6),
			float64(7):  int16(7),
			float64(8):  uint16(8),
			float64(9):  int8(9),
			float64(10): uint8(10),
			float64(11): int(11),
			float64(12): uint(12),
		},
	}

	for index, mapObject := range testMapArray {
		if numberValue, getErr := GetNumber(mapObject, float64(index+1)); getErr != nil {
			t.Fatalf(getErr.Error())
		} else if numberValue != float64(index+1) {
			t.Fatalf("not matched %v vs %v", numberValue, float64(index+1))
		}
	}
}
