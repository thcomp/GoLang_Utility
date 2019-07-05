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
