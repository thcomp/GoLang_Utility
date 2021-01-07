package utility

import (
	"flag"
	"fmt"
	"testing"
)

type Parameter struct {
	InputFile   string `cui_param:"{\"name\": \"i_file\", \"init\": \"\", \"desc\": \"\", \"expect\": \"file|exist\"}"`
	InputFolder string `cui_param:"name: 'i_folder', init: '', desc: '', expect: 'folder|exist'"`
	Thread      int    `cui_param:"name: 'thread', init: 10, desc: 'working thread count'"`
	Debug       bool   `cui_param:"name: 'debug', init: false, desc: 'enable debug log or not'"`
}

func (param *Parameter) String() string {
	return fmt.Sprintf("InputFile(i_file): %v, InputFolder(i_folder): %v, Thread(thread): %v, Debug(debug): %v", param.InputFile, param.InputFolder, param.Thread, param.Debug)
}

func Test(t *testing.T) {
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
		if errors := GetCUIParameter(param); len(errors) == 0 {
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
