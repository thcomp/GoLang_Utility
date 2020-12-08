package utility

import (
	"os"
	"path"
	"path/filepath"
	"testing"
)

func TestPathExt(t *testing.T) {
	var testPathMap map[string]string = map[string]string{
		"/home/a/ss.text.txt":       ".txt",
		"ss.text.dat":               ".dat",
		"./ss.text.Jpeg":            ".Jpeg",
		"../a/ss.text.png":          ".png",
		"C:\\home\\a\\ss.text.a":    ".a",
		".\\ss.text.bn":             ".bn",
		"..\\a\\ss.text.koredemoii": ".koredemoii",
	}

	for key, value := range testPathMap {
		if path.Ext(key) != value {
			t.Errorf("not matched: %s vs %s", path.Ext(key), value)
		}
	}
}

func TestPathBase(t *testing.T) {
	var testPathMap map[string]string = map[string]string{
		"/home/a/ss.text.txt":       "ss.text.txt",
		"ss.text.dat":               "ss.text.dat",
		"./ss.text.Jpeg":            "ss.text.Jpeg",
		"../a/ss.text.png":          "ss.text.png",
		"C:\\home\\a\\ss.text.a":    "ss.text.a",
		".\\ss.text.bn":             "ss.text.bn",
		"..\\a\\ss.text.koredemoii": "ss.text.koredemoii",
	}

	for key, value := range testPathMap {
		if path.Base(key) != value {
			t.Errorf("not matched(%s): %s vs %s", key, path.Ext(key), value)
		}
	}
}

func TestOsName(t *testing.T) {
	var testPathMap map[string]string = map[string]string{
		"/home/a/ss.text.txt":       "ss.text.txt",
		"ss.text.dat":               "ss.text.dat",
		"./ss.text.Jpeg":            "ss.text.Jpeg",
		"../a/ss.text.png":          "ss.text.png",
		"C:\\home\\a\\ss.text.a":    "ss.text.a",
		".\\ss.text.bn":             "ss.text.bn",
		"..\\a\\ss.text.koredemoii": "ss.text.koredemoii",
	}

	for key, value := range testPathMap {
		if keyFile, openErr := os.Open(key); openErr == nil {
			defer keyFile.Close()

			if keyFileInfo, statErr := keyFile.Stat(); statErr == nil {
				if keyFileInfo.Name() != value {
					t.Errorf("not matched(%s): %s vs %s", key, keyFileInfo.Name(), value)
				}
			}
		}
	}
}

func TestPathDir(t *testing.T) {
	var testPathMap map[string]string = map[string]string{
		"/home/a/ss.text.txt":       "/home/a",
		"ss.text.dat":               ".",
		"./ss.text.Jpeg":            ".",
		"../a/ss.text.png":          "../a",
		"C:\\home\\a\\ss.text.a":    "C:\\home\\a",
		".\\ss.text.bn":             ".",
		"..\\a\\ss.text.koredemoii": "..\\a",
	}

	for key, value := range testPathMap {
		if path.Dir(key) != value {
			t.Fatalf("not matched: %s vs %s", path.Dir(key), value)
		}
	}
}

func TestFilepathDir(t *testing.T) {
	var testPathMap map[string]string = map[string]string{
		"/home/a/ss.text.txt":       "/home/a",
		"ss.text.dat":               ".",
		"./ss.text.Jpeg":            ".",
		"../a/ss.text.png":          "../a",
		"C:\\home\\a\\ss.text.a":    "C:\\home\\a",
		".\\ss.text.bn":             ".",
		"..\\a\\ss.text.koredemoii": "..\\a",
	}

	for key, value := range testPathMap {
		if filepath.Dir(key) != value {
			t.Fatalf("not matched: %s vs %s", filepath.Dir(key), value)
		}
	}
}

// func Test(t *testing.T) {
// 	var testPathMap map[string]string = map[string]string{
// 		"/home/a/ss.text.txt":       "/home/a",
// 		"ss.text.dat":               ".",
// 		"./ss.text.Jpeg":            ".",
// 		"../a/ss.text.png":          "../a",
// 		"C:\\home\\a\\ss.text.a":    "C:\\home\\a",
// 		".\\ss.text.bn":             ".",
// 		"..\\a\\ss.text.koredemoii": "..\\a",
// 	}

// 	for key, value := range testPathMap {
// 		if itemInfo, statErr := os.Stat(key); statErr == nil {
// 			itemInfo.
// 		} else {
// 			t.Fatalf(statErr.Error())
// 		}
// 		if (key) != value {
// 			t.Fatalf("not matched: %s vs %s", filepath.Dir(key), value)
// 		}
// 	}
// }
