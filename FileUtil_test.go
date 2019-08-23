package utility

import (
	"os"
	"testing"
)

var ExistTestFiles []string = []string{
	"test1.txt",
	"test2.jpeg",
	"test3.jpg",
	"test4.png",
	"test5.gif",
	"test7.dat",
	"test8.xls",
}
var ExistTestFolders []string = []string{
	"test6",
	"test9",
}
var ExistTestItems []string = []string{}

var NotExistTestFiles []string = []string{
	"test101.txt",
	"test102.jpeg",
	"test103.jpg",
	"test104.png",
	"test105.gif",
	"test107.dat",
	"test108.xls",
}
var NotExistTestFolders []string = []string{
	"test106",
	"test109",
}
var NotExistTestItems []string = []string{}

func execPath() string {
	/*
		ret := ``

		if exe, err := os.Executable(); err == nil {
			ret = filepath.Dir(exe)
		}
	*/
	ret := `C:\\Users\\Hisashi Tatsuguchi\\go\\src\\github.com\\thcomp\\GoLang_Utility`

	return ret
}

func init() {
	for index := range ExistTestFiles {
		ExistTestFiles[index] = "TestData" + string(os.PathSeparator) + ExistTestFiles[index]
		ExistTestItems = append(ExistTestItems, ExistTestFiles[index])
	}
	for index := range ExistTestFolders {
		ExistTestFolders[index] = "TestData" + string(os.PathSeparator) + ExistTestFolders[index]
		ExistTestItems = append(ExistTestItems, ExistTestFolders[index])
	}
	for index := range NotExistTestFiles {
		NotExistTestFiles[index] = "TestData" + string(os.PathSeparator) + NotExistTestFiles[index]
		NotExistTestItems = append(NotExistTestItems, NotExistTestFiles[index])
	}
	for index := range NotExistTestFolders {
		NotExistTestFolders[index] = "TestData" + string(os.PathSeparator) + NotExistTestFolders[index]
		NotExistTestItems = append(NotExistTestItems, NotExistTestFolders[index])
	}
}

func Test_IsExist(t *testing.T) {
	// relation
	for _, item := range ExistTestItems {
		if !IsExist(item) {
			t.Fatalf("%s is exist\n", item)
		}
	}

	// abusolution
	dirname := execPath()
	for _, item := range ExistTestItems {
		absItem := dirname + string(os.PathSeparator) + item
		if !IsExist(absItem) {
			t.Fatalf("%s is exist\n", absItem)
		}
	}

	// relation
	for _, item := range NotExistTestItems {
		if IsExist(item) {
			t.Fatalf("%s is not exist\n", item)
		}
	}

	// abusolution
	dirname = execPath()
	for _, item := range NotExistTestItems {
		absItem := dirname + string(os.PathSeparator) + item
		if IsExist(absItem) {
			t.Fatalf("%s is not exist\n", absItem)
		}
	}
}

func Test_IsFileOrDir(t *testing.T) {
	// relation
	for _, item := range ExistTestFiles {
		if !IsFile(item) {
			t.Fatalf("%s is a file\n", item)
		} else if IsDir(item) {
			t.Fatalf("%s is not a directory\n", item)
		}
	}
	for _, item := range ExistTestFolders {
		if IsFile(item) {
			t.Fatalf("%s is not a file\n", item)
		} else if !IsDir(item) {
			t.Fatalf("%s is a directory\n", item)
		}
	}
	for _, item := range NotExistTestItems {
		if IsFile(item) {
			t.Fatalf("%s is not exist\n", item)
		} else if IsDir(item) {
			t.Fatalf("%s is not exist\n", item)
		}
	}

	// abusolution
	dirname := execPath()
	for _, item := range ExistTestFiles {
		absItem := dirname + string(os.PathSeparator) + item
		if !IsFile(absItem) {
			t.Fatalf("%s is a file\n", absItem)
		} else if IsDir(absItem) {
			t.Fatalf("%s is not a directory\n", absItem)
		}
	}
	for _, item := range ExistTestFolders {
		absItem := dirname + string(os.PathSeparator) + item
		if IsFile(absItem) {
			t.Fatalf("%s is not a file\n", absItem)
		} else if !IsDir(absItem) {
			t.Fatalf("%s is a directory\n", absItem)
		}
	}
	for _, item := range NotExistTestItems {
		absItem := dirname + string(os.PathSeparator) + item
		if IsFile(absItem) {
			t.Fatalf("%s is not exist\n", absItem)
		} else if IsDir(absItem) {
			t.Fatalf("%s is not exist\n", absItem)
		}
	}
}

func Test_GetFiles(t *testing.T) {
	files := GetFilesV("TestData")
	if len(ExistTestFiles) != len(files) {
		t.Fatalf("file count not matched: %d vs %d\n", len(ExistTestFiles), len(files))
	}

	files = GetFilesV("TestData", "png")
	if 1 != len(files) {
		t.Fatalf("file count not matched: %d vs %d\n", 1, len(files))
	}
}
