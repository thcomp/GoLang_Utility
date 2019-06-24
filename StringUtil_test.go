package utility

import (
	"strings"
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

func TestStringUtil3(t *testing.T) {
	var testCaseArray []string = []string{
		`col1 = 1 OR col3="ANDO"`,
		`(col1 = 1 AND col2 = "col2 value") OR col3="ANDO"`,
	}
	var testCaseExpectArray [][][]rune = [][][]rune{
		{[]rune(`col1 = 1 `), []rune(`OR`), []rune(` col3="ANDO"`)},
		{[]rune(`(`), []rune(`col1 = 1 `), []rune(`AND`), []rune(` col2 = "col2 value"`), []rune(`)`), []rune(` `), []rune(`OR`), []rune(` col3="ANDO"`)},
	}

	for i, testCase := range testCaseArray {
		ret := Split(testCase, []string{"AND", "OR", "(", ")"}, []rune{'"', '\''}, false)
		if len(ret) != len(testCaseExpectArray[i]) {
			t.Errorf("length not matched: %d vs %d", len(ret), len(testCaseExpectArray[i]))
		} else {
			for j := 0; j < len(ret); j++ {
				if strings.Compare(string(ret[j]), string(testCaseExpectArray[i][j])) != 0 {
					t.Errorf("not matched: %s vs %s", string(ret[j]), string(testCaseExpectArray[i][j]))
				}
			}
		}
	}
}
