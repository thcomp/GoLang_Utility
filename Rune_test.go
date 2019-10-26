package utility

import (
	"testing"
)

func TestIndexRunes(t *testing.T) {
	testTextArray := []string{
		"ＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMNＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMN",
		"ＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMNＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMN",
		"ＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMNＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMN",
		"ＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMNＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMN",
		"ＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMNＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMN",
		"ＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMNＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMN",
		"ＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMNＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMN",
		"ＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMNＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMN",
	}
	testSubTextArray := []string{
		"1",
		"１",
		"A",
		"Ａ",
		"ABCDE",
		"ＡＢＣＤＥ",
		"ABDE",
		"ＡＢＤＥ",
	}
	expectResultsArray := [][]int{
		{-1, -1},
		{-1, -1},
		{14, 42},
		{0, 28},
		{14, 42},
		{0, 28},
		{-1, -1},
		{-1, -1},
	}

	for i := range testTextArray {
		index := IndexRunes([]rune(testTextArray[i]), []rune(testSubTextArray[i]))
		lastIndex := LastIndexRunes([]rune(testTextArray[i]), []rune(testSubTextArray[i]))

		if index != expectResultsArray[i][0] {
			t.Fatalf("index not matched(%d vs %d): %s @ %s\n", index, expectResultsArray[i][0], testTextArray[i], testSubTextArray[i])
		}
		if lastIndex != expectResultsArray[i][1] {
			t.Fatalf("last index not matched(%d vs %d): %s @ %s\n", lastIndex, expectResultsArray[i][1], testTextArray[i], testSubTextArray[i])
		}
	}
}

func TestPrefixSuffix(t *testing.T) {
	testTextArray := []string{
		"ＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMNＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMN",
		"ＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMNＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMN",
		"ＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMNＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMN",
		"ＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMNＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMN",
		"ＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMNＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMN",
		"ＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMNＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMN",
		"ＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMNＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMN",
		"ＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMNＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMN",

		"ＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMNＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJＫＬＭＮ",
		"ＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMNＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJＫＬＭＮ",
		"ＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMNＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJＫＬＭＮ",
		"ＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMNＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJＫＬＭＮ",
		"ＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMNＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJＫＬＭＮ",
		"ＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJKLMNＡＢＣＤＥＦＧＨＩＪＫＬＭＮABCDEFGHIJＫＬＭＮ",
	}
	testSubTextArray := []string{
		"1",
		"１",
		"A",
		"Ａ",
		"ABCDE",
		"ＡＢＣＤＥ",
		"ABDE",
		"ＡＢＤＥ",

		"N",
		"Ｎ",
		"KLMN",
		"ＫＬＭＮ",
		"IJＫＬＭＮ",
		"IJＫＬＭＮ",
	}
	expectResultsArray := [][]bool{
		{false, false},
		{false, false},
		{false, false},
		{true, false},
		{false, false},
		{true, false},
		{false, false},
		{false, false},

		{false, false},
		{false, true},
		{false, false},
		{false, true},
		{false, true},
		{false, true},
	}

	for i := range testTextArray {
		hasPrefix := HasPrefix([]rune(testTextArray[i]), []rune(testSubTextArray[i]))
		hasSuffix := HasSuffix([]rune(testTextArray[i]), []rune(testSubTextArray[i]))

		if hasPrefix != expectResultsArray[i][0] {
			t.Fatalf("prefix not matched(%t vs %t): %s @ %s\n", hasPrefix, expectResultsArray[i][0], testTextArray[i], testSubTextArray[i])
		}
		if hasSuffix != expectResultsArray[i][1] {
			t.Fatalf("suffix not matched(%t vs %t): %s @ %s\n", hasSuffix, expectResultsArray[i][0], testTextArray[i], testSubTextArray[i])
		}
	}
}
