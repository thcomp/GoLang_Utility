package utility

import (
	"testing"
)

func TestStringUtil(t *testing.T) {
	var testCaseArray [][]string = [][]string{
		{`a_bcd_efg_hijklmnopq_r_s_t_u`, `ABcdEfgHijklmnopqRSTU`},
		{`a_bｃd_eｆg_hijklmnopq_r_s_t_u`, `ABｃdEｆgHijklmnopqRSTU`}, // 全角文字含む(ｃ、ｆ)
		{`content_url`, `ContentUrl`},
		{`start_time`, `StartTime`},
		{`end_time`, `EndTime`},
	}

	for _, testCase := range testCaseArray {
		var result string

		result = SnakeToCamelCase(testCase[0])
		if result != testCase[1] {
			t.Error(`SnakeToCamelCase is failed: result=` + result + `, ` + testCase[0] + `->` + testCase[1])
		}
		result = CamelToSnakeCase(testCase[1])
		if result != testCase[0] {
			t.Error(`CamelToSnakeCase is failed: result=` + result + `, ` + testCase[1] + `->` + testCase[0])
		}
	}
}

func TestStringUtil2(t *testing.T) {
	var testCaseArray [][]string = [][]string{
		{`Word`, `word`},
	}

	for _, testCase := range testCaseArray {
		var result string

		result = ToLowerFirst(testCase[0])
		if result != testCase[1] {
			t.Error(`ToLowerFirst is failed: result=` + result + `, ` + testCase[0] + `->` + testCase[1])
		}
		result = ToUpperFirst(testCase[1])
		if result != testCase[0] {
			t.Error(`ToLowerFirst is failed: result=` + result + `, ` + testCase[1] + `->` + testCase[0])
		}
	}
}

// func TestStringUtil3(t *testing.T) {
// 	var testCaseArray []string = []string{
// 		`col1 = 1 OR col3="ANDO"`,
// 		`(col1 = 1 AND col2 = "col2 value") OR col3="ANDO"`,
// 	}
// 	var testCaseExpectArray [][][]rune = [][][]rune{
// 		{[]rune(`col1 = 1 `), []rune(`OR`), []rune(` col3="ANDO"`)},
// 		{[]rune(`(`), []rune(`col1 = 1 `), []rune(`AND`), []rune(` col2 = "col2 value"`), []rune(`)`), []rune(` `), []rune(`OR`), []rune(` col3="ANDO"`)},
// 	}

// 	for i, testCase := range testCaseArray {
// 		ret := Split(testCase, []string{"AND", "OR", "(", ")"}, []rune{'"', '\''}, false)
// 		if len(ret) != len(testCaseExpectArray[i]) {
// 			t.Errorf("length not matched: %d vs %d", len(ret), len(testCaseExpectArray[i]))
// 		} else {
// 			for j := 0; j < len(ret); j++ {
// 				if strings.Compare(string(ret[j]), string(testCaseExpectArray[i][j])) != 0 {
// 					t.Errorf("not matched: %s vs %s", string(ret[j]), string(testCaseExpectArray[i][j]))
// 				}
// 			}
// 		}
// 	}
// }

func TestStringUtil4(t *testing.T) {
	var testCaseArray []string = []string{
		`"json1": "value1, value2", "json2": "value3, value4"`,
		`"json1": 'value1, value2', "json2": "value3, value4", "json3": "value3, 'value4'"`,
	}
	var expectArray = [][]string{
		[]string{`"json1": "value1, value2"`, ` "json2": "value3, value4"`},
		[]string{`"json1": 'value1, value2'`, ` "json2": "value3, value4"`, ` "json3": "value3, 'value4'"`},
	}

	for index, testCase := range testCaseArray {
		resultRunesSlice := Split(testCase, []string{","}, "\"")

		if len(resultRunesSlice) == len(expectArray[index]) {
			for index2, resultRunes := range resultRunesSlice {
				if string(resultRunes) != expectArray[index][index2] {
					t.Fatalf("not matched: %s vs %s", string(resultRunes), expectArray[index][index2])
				}
			}
		} else {
			t.Fatalf("not matched: `%s`: %v(%d) vs %v(%d)", testCaseArray[index], resultRunesSlice, len(resultRunesSlice), expectArray[index], len(expectArray[index]))
		}
	}
}

func TestStringUtil5(t *testing.T) {
	var testCaseArray []string = []string{
		`"json1": "value1, value2", "json2": "value3, value4"`,
		`"json1": 'value1, value2', "json2": "value3, value4", "json3": "value3, 'value4'"`,
	}
	var expectArray = [][]string{
		[]string{`"json1"`, ` "value1, value2"`, ` "json2"`, ` "value3, value4"`},
		[]string{`"json1"`, ` 'value1`, ` value2'`, ` "json2"`, ` "value3, value4"`, ` "json3"`, ` "value3, 'value4'"`},
	}

	for index, testCase := range testCaseArray {
		resultRunesSlice := Split(testCase, []string{",", ":"}, "\"")

		if len(resultRunesSlice) == len(expectArray[index]) {
			for index2, resultRunes := range resultRunesSlice {
				if string(resultRunes) != expectArray[index][index2] {
					t.Fatalf("not matched: %s vs %s", string(resultRunes), expectArray[index][index2])
				}
			}
		} else {
			t.Fatalf("not matched: `%s`: %v(%d) vs %v(%d)", testCaseArray[index], resultRunesSlice, len(resultRunesSlice), expectArray[index], len(expectArray[index]))
		}
	}
}

func Test_ParseNumber(t *testing.T) {
	testSlice := []struct {
		Str         string
		ExpectValue float64
	}{
		{Str: "100", ExpectValue: float64(100)},
		{Str: "-200", ExpectValue: float64(-200)},
		{Str: "1.112", ExpectValue: float64(1.112)},
		{Str: "-0.112", ExpectValue: float64(-0.112)},
		{Str: "0x1123", ExpectValue: float64(0x1123)},
		{Str: "-0x112F", ExpectValue: float64(-0x112F)},
		{Str: "b'1100'", ExpectValue: float64(12)},
		{Str: "-b'1100'", ExpectValue: float64(-12)},
	}

	for _, test := range testSlice {
		if value, err := ParseNumber(test.Str); err == nil {
			if value != test.ExpectValue {
				t.Fatalf("not matched: %f vs %f", value, test.ExpectValue)
			}
		} else {
			t.Fatalf("%v: %v\n", test, err)
		}
	}
}
