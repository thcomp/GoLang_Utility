package utility

import (
	"strings"
	"testing"
)

func Test_CreatePassword(t *testing.T) {
	lengthSlice := []int{8, 12, 16, 20, 24, 28, 32}
	validCharSlice := [][]int{
		[]int{PwValidSmallAlpha},
		[]int{PwValidSmallAlpha, PwValidLargeAlpha},
		[]int{PwValidSmallAlpha, PwValidLargeAlpha, PwValidNumeric},
		[]int{PwValidSmallAlpha, PwValidLargeAlpha, PwValidNumeric, PwValidSymbol},
	}
	validCharTextSlice := []string{
		"abcdefghijklmnopqrstuvwxyz",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\"#$%&'()=-~^|\\`@{[+;*:}]<,>.?/_\\",
	}

	for _, length := range lengthSlice {
		for charIndex, validChars := range validCharSlice {
			pw := CreatePassword(length, validChars...)
			t.Logf("%d, %v => %s\n", length, validChars, pw)

			if len(pw) != length {
				t.Fatalf("not match length: %d vs %d", len(pw), length)
			}
			for _, char := range []byte(pw) {
				if strings.Contains(validCharTextSlice[charIndex], string(char)) == false {
					t.Fatalf("not found in valid chars: %s @ %s", string(char), validCharTextSlice[charIndex])
				}
			}
		}
	}
}
