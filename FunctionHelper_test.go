package utility

import (
	"reflect"
	"testing"
	"time"
)

func Test_RunAfterRandomDelay(t *testing.T) {
	valueInt := []int{5, 8, 13}
	retSlice1 := CallAfterRandomDelay(reflect.ValueOf(targetFunc1), t, valueInt[0], valueInt[1])
	if len(retSlice1) != 3 {
		t.Fatalf("ret not match: %v", retSlice1)
	} else {
		for index, retValue := range retSlice1 {
			if int(retValue.Int()) != valueInt[index] {
				t.Fatalf("ret not match: %v", retSlice1)
			}
		}
	}

	test := sFuncTest{t: t}
	retSlice1 = CallAfterRandomDelay(reflect.ValueOf(test.Test), valueInt[0], valueInt[1])
	if len(retSlice1) != 3 {
		t.Fatalf("ret not match: %v", retSlice1)
	} else {
		for index, retValue := range retSlice1 {
			if int(retValue.Int()) != valueInt[index] {
				t.Fatalf("ret not match: %v", retSlice1)
			}
		}
	}
}

func targetFunc1(t *testing.T, valueInt1, valueInt2 int) (int, int, int) {
	t.Logf("%d: value: %d, %d", time.Now().UnixNano(), valueInt1, valueInt2)
	return valueInt1, valueInt2, valueInt1 + valueInt2
}

type sFuncTest struct {
	t *testing.T
}

func (test *sFuncTest) Test(valueInt1, valueInt2 int) (int, int, int) {
	test.t.Logf("%d: value: %d, %d", time.Now().UnixNano(), valueInt1, valueInt2)
	return valueInt1, valueInt2, valueInt1 + valueInt2
}
