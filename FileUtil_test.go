package utility

import (
	"os"
	"path"
	"path/filepath"
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

func Test_AppendSuffix(t *testing.T) {
	testDataArray := []string{
		"/var/www/html/index.html",
		"index.html",
		"./index.html",
		"../html/index.html",
		"c:\\var\\www\\html\\index.html",
		".\\index.html",
		"..\\html\\index.html",
	}
	expectDataArray := []string{
		"/var/www/html/index_suffiX.html",
		"index_suffiX.html",
		"./index_suffiX.html",
		"../html/index_suffiX.html",
		"c:\\var\\www\\html\\index_suffiX.html",
		".\\index_suffiX.html",
		"..\\html\\index_suffiX.html",
	}

	for pos := range testDataArray {
		if AppendSuffix(testDataArray[pos], "_suffiX") != expectDataArray[pos] {
			t.Fatalf("not matched: %s vs %s", AppendSuffix(testDataArray[pos], "_suffiX"), expectDataArray[pos])
		}
	}
}

func Test_AppendPrefix(t *testing.T) {
	testDataArray := []string{
		"/var/www/html/index.html",
		"index.html",
		"./index.html",
		"../html/index.html",
		"c:\\var\\www\\html\\index.html",
		".\\index.html",
		"..\\html\\index.html",
	}
	expectDataArray := []string{
		"/var/www/html/preFix_index.html",
		"preFix_index.html",
		"./preFix_index.html",
		"../html/preFix_index.html",
		"c:\\var\\www\\html\\preFix_index.html",
		".\\preFix_index.html",
		"..\\html\\preFix_index.html",
	}

	for pos := range testDataArray {
		if AppendPrefix(testDataArray[pos], "preFix_") != expectDataArray[pos] {
			t.Fatalf("not matched: %s vs %s", AppendPrefix(testDataArray[pos], "preFix_"), expectDataArray[pos])
		}
	}
}

func TestExt(t *testing.T) {
	testDataArray := []string{
		"/var/www/html/index.html",
		"index.html",
		"./index.html",
		"../html/index.html",
		"c:\\var\\www\\html\\index.html",
		".\\index.html",
		"..\\html\\index.html",
		"..\\html\\index",
	}
	expectDataArray := []string{
		".html",
		".html",
		".html",
		".html",
		".html",
		".html",
		".html",
		".\\html\\index",
	}

	for pos := range testDataArray {
		t.Logf("%s -> %s\n", testDataArray[pos], path.Ext(testDataArray[pos]))

		if path.Ext(testDataArray[pos]) != expectDataArray[pos] {
			t.Fatalf("not matched: %s vs %s\n", path.Ext(testDataArray[pos]), expectDataArray[pos])
		}
	}
}

func Test_GetMIMETypeFromExtension(t *testing.T) {
	testDataArray := []string{
		"/var/www/html/image.png",
		"test.pdf",
	}
	expectDataArray := []string{
		"image/png",
		"application/pdf",
	}

	for pos := range testDataArray {
		if GetMIMETypeFromExtension(testDataArray[pos]) != expectDataArray[pos] {
			t.Fatalf("not matched: %s vs %s\n", GetMIMETypeFromExtension(testDataArray[pos]), expectDataArray[pos])
		}
	}
}

func TestGetParent(t *testing.T) {
	var testPathMap map[string]string = map[string]string{
		"/home/a/ss.text.txt":       "/home/a",
		"ss.text.dat":               ".",
		"./ss.text.Jpeg":            ".",
		"../a/ss.text.png":          "../a",
		"/ss.text.png":              ".",
		"C:\\home\\a\\ss.text.a":    "C:\\home\\a",
		".\\ss.text.bn":             ".",
		"..\\a\\ss.text.koredemoii": "..\\a",
		"C:\\ss.text.a":             "C:",
		"\\ss.text.a":               ".",
	}

	for key, value := range testPathMap {
		if GetParent(key) != value {
			t.Fatalf("not matched: %s vs %s", filepath.Dir(key), value)
		}
	}
}

func Test_exchangePath(t *testing.T) {
	type TestData struct {
		OriginalFilepath     string
		ExpectFilepath       string
		OldNewPathSeparaters []byte
	}

	testDataArray := []TestData{
		{
			OriginalFilepath:     "C:\\Windows\\system\\drivers\\log.txt",
			ExpectFilepath:       "/mnt/c/Windows/system/drivers/log.txt",
			OldNewPathSeparaters: []byte{'\\', '/'},
		},
		{
			OriginalFilepath:     "\\Windows\\system\\drivers\\log.txt",
			ExpectFilepath:       "/mnt/c/Windows/system/drivers/log.txt",
			OldNewPathSeparaters: []byte{'\\', '/'},
		},
		{
			OriginalFilepath:     "/mnt/c/Windows/system/drivers/log.txt",
			ExpectFilepath:       "c:\\Windows\\system\\drivers\\log.txt",
			OldNewPathSeparaters: []byte{'/', '\\'},
		},

		{
			OriginalFilepath:     "C:\\Windows\\system\\drivers\\log.txt",
			ExpectFilepath:       "C:\\Windows\\system\\drivers\\log.txt",
			OldNewPathSeparaters: []byte{'/', '\\'},
		},
		{
			OriginalFilepath:     "\\Windows\\system\\drivers\\log.txt",
			ExpectFilepath:       "\\Windows\\system\\drivers\\log.txt",
			OldNewPathSeparaters: []byte{'/', '\\'},
		},
		{
			OriginalFilepath:     "/mnt/c/Windows/system/drivers/log.txt",
			ExpectFilepath:       "/mnt/c/Windows/system/drivers/log.txt",
			OldNewPathSeparaters: []byte{'\\', '/'},
		},

		{
			OriginalFilepath:     "system\\drivers\\log.txt",
			ExpectFilepath:       "system/drivers/log.txt",
			OldNewPathSeparaters: []byte{'\\', '/'},
		},
		{
			OriginalFilepath:     ".\\system\\drivers\\log.txt",
			ExpectFilepath:       "./system/drivers/log.txt",
			OldNewPathSeparaters: []byte{'\\', '/'},
		},
		{
			OriginalFilepath:     "system/drivers/log.txt",
			ExpectFilepath:       "system\\drivers\\log.txt",
			OldNewPathSeparaters: []byte{'/', '\\'},
		},
		{
			OriginalFilepath:     "./system/drivers/log.txt",
			ExpectFilepath:       ".\\system\\drivers\\log.txt",
			OldNewPathSeparaters: []byte{'/', '\\'},
		},
		{
			OriginalFilepath:     "..\\system\\drivers\\log.txt",
			ExpectFilepath:       "../system/drivers/log.txt",
			OldNewPathSeparaters: []byte{'\\', '/'},
		},
		{
			OriginalFilepath:     "../system/drivers/log.txt",
			ExpectFilepath:       "..\\system\\drivers\\log.txt",
			OldNewPathSeparaters: []byte{'/', '\\'},
		},
	}

	for index, testData := range testDataArray {
		result := exchangePath(testData.OriginalFilepath, testData.OldNewPathSeparaters)
		if testData.ExpectFilepath != result {
			t.Fatalf("%d: not matched, %s to %s(expect: %s)", index, testData.OriginalFilepath, result, testData.ExpectFilepath)
		}
	}
}
