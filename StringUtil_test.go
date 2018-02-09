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
