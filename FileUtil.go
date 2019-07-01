package utility

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

func IsExist(path string) bool {
	var ret = false

	if target, openErr := os.Open(path); openErr == nil {
		defer target.Close()

		if _, statErr := target.Stat(); statErr == nil {
			ret = true
		} else {
			ret = os.IsExist(statErr)
		}
	}

	return ret
}

func IsFile(path string) bool {
	var ret = false

	if target, openErr := os.Open(path); openErr == nil {
		defer target.Close()

		if targetFileInfo, statErr := target.Stat(); statErr == nil {
			ret = !(targetFileInfo.IsDir())
		}
	}

	return ret
}

func IsDir(path string) bool {
	var ret = false

	if target, openErr := os.Open(path); openErr == nil {
		defer target.Close()

		if targetFileInfo, statErr := target.Stat(); statErr == nil {
			ret = targetFileInfo.IsDir()
		}
	}

	return ret
}

func GetFiles(path string, allowExtensions []string) []string {
	var ret []string = make([]string, 0)

	if target, openErr := os.Open(path); openErr == nil {
		defer target.Close()

		allowExtensionMap := map[string]bool{}
		for _, extension := range allowExtensions {
			lowerExtension := strings.ToLower(extension)
			if !strings.HasPrefix(lowerExtension, `.`) {
				lowerExtension = `.` + lowerExtension
			}
			allowExtensionMap[lowerExtension] = true
		}

		childElements := getFilesSub(target, allowExtensionMap)
		if childElements != nil && len(childElements) > 0 {
			for _, childElement := range childElements {
				ret = append(ret, childElement)
			}
		}
	}

	return ret
}

func getFilesSub(target *os.File, allowExtensionMap map[string]bool) []string {
	var ret []string = make([]string, 0)

	if targetStat, statErr := target.Stat(); statErr == nil {
		if targetStat.IsDir() {
			if childFileInfoArray, readErr := target.Readdir(0); readErr == nil {
				for _, childFileInfo := range childFileInfoArray {
					if childTarget, openErr := os.Open(target.Name() + string(os.PathSeparator) + childFileInfo.Name()); openErr == nil {
						defer childTarget.Close()

						childElements := getFilesSub(childTarget, allowExtensionMap)
						if childElements != nil {
							for _, childElement := range childElements {
								ret = append(ret, childElement)
							}
						}
					}
				}
			} else {
				lowerExt := strings.ToLower(path.Ext(target.Name()))
				if value, exist := allowExtensionMap[lowerExt]; exist && value {
					ret = append(ret, target.Name())
				}
			}
		} else {
			lowerExt := strings.ToLower(path.Ext(target.Name()))
			if value, exist := allowExtensionMap[lowerExt]; exist && value {
				ret = append(ret, target.Name())
			}
		}
	}

	return ret
}

func ChangeExtension(path string, newExtension string) string {
	var ret = path
	var fixedNewExtension string

	if !strings.HasPrefix(newExtension, `.`) {
		fixedNewExtension = `.` + newExtension
	} else {
		fixedNewExtension = newExtension
	}

	if oldExtension := filepath.Ext(path); len(oldExtension) > 0 {
		ret = ret[0:len(path)-len(oldExtension)] + fixedNewExtension
	} else {
		ret = ret + fixedNewExtension
	}

	return ret
}
