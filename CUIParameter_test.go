package utility

import (
	"flag"
	"fmt"
	"testing"
)

type ParameterCommon struct {
	Thread int  `cui_param:"name: 'thread', init: 10, desc: 'working thread count'"`
	Debug  bool `cui_param:"name: 'debug', init: false, desc: 'enable debug log or not'"`
}

type Parameter struct {
	ParameterCommon
	InputFile   string `cui_param:"{\"name\": \"i_file\", \"init\": \"\", \"desc\": \"\", \"expect\": \"file|exist\"}"`
	InputFolder string `cui_param:"name: 'i_folder', init: '', desc: '', expect: 'folder|exist'"`
}

func (param *Parameter) String() string {
	return fmt.Sprintf("InputFile(i_file): %v, InputFolder(i_folder): %v, Thread(thread): %v, Debug(debug): %v", param.InputFile, param.InputFolder, param.Thread, param.Debug)
}

func Test_GetCUIParameter(t *testing.T) {
	tests := [][]struct {
		name  string
		value string
	}{
		{
			{name: "i_file", value: "C:\\Users\\Hisashi Tatsuguchi\\ownCloud\\全社共通\\test.txt"},
			{name: "i_folder", value: "C:\\Users\\Hisashi Tatsuguchi\\ownCloud\\全社共通"},
			{name: "debug", value: "true"},
			{name: "thread", value: "11"},
		},
	}

	for _, paramSlice := range tests {
		for _, param := range paramSlice {
			flag.CommandLine.Init(param.name, flag.ContinueOnError)
			flag.CommandLine.Set(param.name, param.value)
		}

		param := &Parameter{}
		if errors := GetCUIParameter(param, false); len(errors) == 0 {
			if param.InputFile != paramSlice[0].value {
				t.Fatalf("not matched: %s vs %s", param.InputFile, paramSlice[0].value)
			}
			if param.InputFolder != paramSlice[1].value {
				t.Fatalf("not matched: %s vs %s", param.InputFolder, paramSlice[1].value)
			}
			if fmt.Sprintf("%v", param.Debug) != paramSlice[2].value {
				t.Fatalf("not matched: %s vs %s", fmt.Sprintf("%v", param.Debug), paramSlice[2].value)
			}
			if fmt.Sprintf("%v", param.Thread) != paramSlice[3].value {
				t.Fatalf("not matched: %s vs %s", fmt.Sprintf("%v", param.Thread), paramSlice[3].value)
			}
		}
	}
}

func Test_IsValidNumberParameter(t *testing.T) {
	tests := []struct {
		expect       string
		targetValue  float64
		expectResult bool
	}{
		{expect: "0-1", targetValue: 0, expectResult: true},
		{expect: "0-1", targetValue: 1, expectResult: true},
		{expect: "0-1", targetValue: 0.5, expectResult: true},
		{expect: "0<v<1", targetValue: 0, expectResult: false},
		{expect: "0<v<1", targetValue: 1, expectResult: false},
		{expect: "0<v<1", targetValue: 0.1, expectResult: true},
		{expect: "0<v<1", targetValue: 0.9, expectResult: true},
		{expect: "0<=v<=1", targetValue: 0, expectResult: true},
		{expect: "0<=v<=1", targetValue: 1, expectResult: true},
		{expect: "-10--9", targetValue: -10, expectResult: true},
		{expect: "-10--9", targetValue: -9, expectResult: true},
		{expect: "-10--9", targetValue: -9.5, expectResult: true},
		{expect: "-9>v>-10", targetValue: -10, expectResult: false},
		{expect: "-9>v>-10", targetValue: -9, expectResult: false},
		{expect: "-9>v>-10", targetValue: -9.5, expectResult: true},
		{expect: "-9>=v>=-10", targetValue: -10, expectResult: true},
		{expect: "-9>=v>=-10", targetValue: -9, expectResult: true},

		{expect: "0-0x1", targetValue: 0, expectResult: true},
		{expect: "0-0x1", targetValue: 1, expectResult: true},
		{expect: "0-0x1", targetValue: 0.5, expectResult: true},
		{expect: "0<v<0x1", targetValue: 0, expectResult: false},
		{expect: "0<v<0x1", targetValue: 1, expectResult: false},
		{expect: "0<v<0x1", targetValue: 0.1, expectResult: true},
		{expect: "0<v<0x1", targetValue: 0.9, expectResult: true},
		{expect: "0<=v<=0x1", targetValue: 0, expectResult: true},
		{expect: "0<=v<=0x1", targetValue: 1, expectResult: true},
		{expect: "-0x10--0x09", targetValue: -0x10, expectResult: true},
		{expect: "-0x10--0x09", targetValue: -0x09, expectResult: true},
		{expect: "-0x09>v>-0x10", targetValue: -0x10, expectResult: false},
		{expect: "-0x09>v>-0x10", targetValue: -0x09, expectResult: false},
		{expect: "-0x09>=v>=-0x10", targetValue: -0x10, expectResult: true},
		{expect: "-0x09>=v>=-0x10", targetValue: -0x09, expectResult: true},
	}

	for _, test := range tests {
		err := isValidNumberParameter(test.targetValue, test.expect)

		if test.expectResult && err != nil {
			t.Fatalf("not matched: %v, %v", test, err)
		} else if !test.expectResult && err == nil {
			t.Fatalf("not matched: %v, %v", test, err)
		}
	}

}
