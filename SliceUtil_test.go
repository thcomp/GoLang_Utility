package utility

import (
	"reflect"
	"testing"
)

func Test_RemoveSliceItem1(t *testing.T) {
	originalSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	originalSliceLen := len(originalSlice)
	sliceInterface := interface{}(originalSlice)

	for i := 0; i < originalSliceLen; i++ {
		result := false
		if sliceInterface, result = RemoveSliceItem(sliceInterface, 0); result {
			if sliceValue, assertionOK := sliceInterface.([]int); assertionOK {
				if len(sliceValue) != originalSliceLen-i-1 {
					t.Fatalf("not match length: %v", sliceValue)
				} else {
					for j := 0; j < originalSliceLen-i-1; j++ {
						if sliceValue[j] != i+j+1 {
							t.Fatalf("not match value: %v[%d]", sliceValue, j)
						}
					}
				}

				sliceInterface = sliceValue
			} else {
				t.Fatalf("not match type: %v", reflect.TypeOf(sliceInterface))
			}
		} else {
			t.Fatalf("fail to remove: %v", sliceInterface)
		}
	}
}

func Test_RemoveSliceItem2(t *testing.T) {
	originalSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	originalSliceLen := len(originalSlice)
	sliceInterface := interface{}(originalSlice)

	for i := 0; i < originalSliceLen; i++ {
		result := false
		if sliceInterface, result = RemoveSliceItem(sliceInterface, originalSliceLen-i-1); result {
			if sliceValue, assertionOK := sliceInterface.([]int); assertionOK {
				if len(sliceValue) != originalSliceLen-i-1 {
					t.Fatalf("not match length: %v", sliceValue)
				} else {
					for j := 0; j < originalSliceLen-i-1; j++ {
						if sliceValue[j] != j {
							t.Fatalf("not match value: %v[%d]", sliceValue, j)
						}
					}
				}

				sliceInterface = sliceValue
			} else {
				t.Fatalf("not match type: %v", reflect.TypeOf(sliceInterface))
			}
		} else {
			t.Fatalf("fail to remove: %v", sliceInterface)
		}
	}
}

func Test_RemoveSliceItem3(t *testing.T) {
	originalSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	originalSliceLen := len(originalSlice)

	for i := 0; i < originalSliceLen; i++ {
		sliceInterface := interface{}(originalSlice)
		result := false
		if sliceInterface, result = RemoveSliceItem(sliceInterface, i); result {
			if sliceValue, assertionOK := sliceInterface.([]int); assertionOK {
				if len(sliceValue) != originalSliceLen-1 {
					t.Fatalf("not match length: %v", sliceValue)
				} else {
					for j := 0; j < originalSliceLen-1; j++ {
						if j >= i {
							if sliceValue[j] != (j + 1) {
								t.Fatalf("not match value: %v[i: %d, j: %d]", sliceValue, i, j)
							}
						} else {
							if sliceValue[j] != j {
								t.Fatalf("not match value: %v[i: %d, j: %d]", sliceValue, i, j)
							}
						}
					}
				}
			} else {
				t.Fatalf("not match type: %v", reflect.TypeOf(sliceInterface))
			}
		} else {
			t.Fatalf("fail to remove: %v", sliceInterface)
		}
	}
}
