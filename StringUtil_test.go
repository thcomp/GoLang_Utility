package utility

import (
	"testing"
)

func TestStringUtil(t *testing.T) {
	var testCaseArray [][]string = [][]string{
		{`a_bcd_efg_hijklmnopq_r_s_t_u`, `ABcdEfgHijklmnopqRSTU`},
		{`a_bｃd_eｆg_hijklmnopq_r_s_t_u`, `ABｃdEｆgHijklmnopqRSTU`}, // 全角文字含む(ｃ、ｆ)
	}

	for _, testCase := range testCaseArray {
		var result string

		result = SnakeToCamelCase(testCase[0])
		if result != testCase[1] {
			t.Error(`SnakeToCamelCase is failed: ` + result)
		}
		result = CamelToSnakeCase(testCase[1])
		if result != testCase[0] {
			t.Error(`CamelToSnakeCase is failed: ` + result)
		}
	}
}
