package utility

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

var sExtensionMIMEMap map[string]string = map[string]string{
	".aac":    "audio/aac",
	".abw":    "application/x-abiword",
	".arc":    "application/x-freearc",
	".avi":    "video/x-msvideo",
	".azw":    "application/vnd.amazon.ebook",
	".bin":    "application/octet-stream",
	".bmp":    "image/bmp",
	".bz":     "application/x-bzip",
	".bz2":    "application/x-bzip2",
	".csh":    "application/x-csh",
	".css":    "text/css",
	".csv":    "text/csv",
	".doc":    "application/msword",
	".docx":   "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	".eot":    "application/vnd.ms-fontobject",
	".epub":   "application/epub+zip",
	".gz":     "application/gzip",
	".gif":    "image/gif",
	".htm":    "text/html",
	".html":   "text/html",
	".ico":    "image/vnd.microsoft.icon",
	".ics":    "text/calendar",
	".jar":    "application/java-archive",
	".jpeg":   "image/jpeg",
	".jpg":    "image/jpeg",
	".js":     "text/javascript",
	".json":   "application/json",
	".jsonld": "application/ld+json",
	".mid":    "audio/x-midi",
	".midi":   "audio/x-midi",
	".mjs":    "text/javascript",
	".mp3":    "audio/mpeg",
	".mpeg":   "video/mpeg",
	".mpkg":   "application/vnd.apple.installer+xml",
	".odp":    "application/vnd.oasis.opendocument.presentation",
	".ods":    "application/vnd.oasis.opendocument.spreadsheet",
	".odt":    "application/vnd.oasis.opendocument.text",
	".oga":    "audio/ogg",
	".ogv":    "video/ogg",
	".ogx":    "application/ogg",
	".otf":    "font/otf",
	".png":    "image/png",
	".pdf":    "application/pdf",
	".php":    "appliction/php",
	".ppt":    "application/vnd.ms-powerpoint",
	".pptx":   "application/vnd.openxmlformats-officedocument.presentationml.presentation",
	".rar":    "application/x-rar-compressed",
	".rtf":    "application/rtf",
	".sh":     "application/x-sh",
	".svg":    "image/svg+xml",
	".swf":    "application/x-shockwave-flash",
	".tar":    "application/x-tar",
	".tif":    "image/tiff",
	".tiff":   "image/tiff",
	".ts":     "video/mp2t",
	".ttf":    "font/ttf",
	".txt":    "text/plain",
	".vsd":    "application/vnd.visio",
	".wav":    "audio/wav",
	".weba":   "audio/webm",
	".webm":   "video/webm",
	".webp":   "image/webp",
	".woff":   "font/woff",
	".woff2":  "font/woff2",
	".xhtml":  "application/xhtml+xml",
	".xls":    "application/vnd.ms-excel",
	".xlsx":   "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	".xml":    "application/xml",
	".xul":    "application/vnd.mozilla.xul+xml",
	".zip":    "application/zip",
	".3gp":    "video/3gpp",
	".3g2":    "video/3gpp2",
	".7z":     "application/x-7z-compressed",
}

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

func GetFilesV(path string, allowExtensions ...string) []string {
	var ret []string = make([]string, 0)

	if target, openErr := os.Open(path); openErr == nil {
		defer target.Close()

		allowExtensionMap := map[string]bool{}

		if allowExtensions == nil || len(allowExtensions) == 0 {
			// allow all extension
		} else {
			for _, extension := range allowExtensions {
				lowerExtension := strings.ToLower(extension)
				if !strings.HasPrefix(lowerExtension, `.`) {
					lowerExtension = `.` + lowerExtension
				}
				allowExtensionMap[lowerExtension] = true
			}
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

func GetFiles(path string, allowExtensions []string) []string {
	var ret []string

	if allowExtensions == nil {
		ret = GetFilesV(path)
	} else {
		ret = GetFilesV(path, allowExtensions...)
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

				if len(allowExtensionMap) == 0 {
					// allow all extension
					ret = append(ret, target.Name())
				} else {
					if value, exist := allowExtensionMap[lowerExt]; exist && value {
						ret = append(ret, target.Name())
					}
				}
			}
		} else {
			lowerExt := strings.ToLower(path.Ext(target.Name()))

			if len(allowExtensionMap) == 0 {
				// allow all extension
				ret = append(ret, target.Name())
			} else {
				if value, exist := allowExtensionMap[lowerExt]; exist && value {
					ret = append(ret, target.Name())
				}
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

func AppendSuffix(path, suffix string) string {
	var ret = path

	if extension := filepath.Ext(path); len(extension) > 0 {
		ret = ret[0:len(path)-len(extension)] + suffix + extension
	} else {
		ret = ret + suffix
	}

	return ret
}

func AppendPrefix(path, prefix string) string {
	var ret = path

	if base := filepath.Base(path); len(base) > 0 {
		ret = ret[0:len(path)-len(base)] + prefix + base
	} else {
		ret = prefix + ret
	}

	return ret
}

func GetMIMETypeFromExtension(filePath string) string {
	ret := ""
	extension := path.Ext(filePath)
	extension = strings.ToLower(extension)

	if mime, exit := sExtensionMIMEMap[extension]; exit {
		ret = mime
	} else {
		ret = "applicaion/octet-stream"
	}

	return ret
}
