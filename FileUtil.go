package utility

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
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
	".cer":    "application/pkix-cert",
	".crl":    "application/pkix-crl",
	".crt":    "application/x-x509-user-cert",
	".csh":    "application/x-csh",
	".csr":    "application/pkcs10",
	".css":    "text/css",
	".csv":    "text/csv",
	".der":    "application/x-x509-ca-cert",
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
	".key":    "application/pkcs8",
	".mid":    "audio/x-midi",
	".midi":   "audio/x-midi",
	".mjs":    "text/javascript",
	".mov":    "video/quicktime",
	".mp3":    "audio/mpeg",
	".mp4":    "video/mpeg",
	".mpeg":   "video/mpeg",
	".mpkg":   "application/vnd.apple.installer+xml",
	".odp":    "application/vnd.oasis.opendocument.presentation",
	".ods":    "application/vnd.oasis.opendocument.spreadsheet",
	".odt":    "application/vnd.oasis.opendocument.text",
	".oga":    "audio/ogg",
	".ogv":    "video/ogg",
	".ogx":    "application/ogg",
	".otf":    "font/otf",
	".p10":    "application/pkcs10",
	".p12":    "application/x-pkcs12",
	".p7b":    "application/x-pkcs7-certificates",
	".p7c":    "application/pkcs7-mime",
	".p7r":    "application/x-pkcs7-certreqresp",
	".p8":     "application/pkcs8",
	".pem":    "application/x-pem-file",
	".pfx":    "application/x-pkcs12",
	".pdf":    "application/pdf",
	".php":    "appliction/php",
	".png":    "image/png",
	".ppt":    "application/vnd.ms-powerpoint",
	".pptx":   "application/vnd.openxmlformats-officedocument.presentationml.presentation",
	".rar":    "application/x-rar-compressed",
	".rtf":    "application/rtf",
	".sh":     "application/x-sh",
	".spc":    "application/x-pkcs7-certificates",
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

var sMIMEExtensionMap map[string]string

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

	if targetFileInfo, openErr := os.Stat(path); openErr == nil {
		ret = !(targetFileInfo.IsDir())
	}

	return ret
}

func IsFileByFileInfo(fileInfo os.FileInfo) bool {
	return !fileInfo.IsDir()
}

func IsDir(path string) bool {
	var ret = false

	if targetFileInfo, openErr := os.Stat(path); openErr == nil {
		ret = targetFileInfo.IsDir()
	}

	return ret
}

func IsSymlink(path string) bool {
	var ret = false

	if targetFileInfo, openErr := os.Stat(path); openErr == nil {
		ret = (targetFileInfo.Mode()&os.ModeSymlink == os.ModeSymlink)
	}

	return ret
}

func IsSymlinkByFileInfo(fileInfo os.FileInfo) bool {
	return (fileInfo.Mode()&os.ModeSymlink == os.ModeSymlink)
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

		// deferではリリースが遅れてリソース枯渇が発生するので、明示的にも開放
		target.Close()
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

						// deferではリリースが遅れてリソース枯渇が発生するので、明示的にも開放
						childTarget.Close()
					} else {
						LogfE("fail to open file: %s, %v", target.Name()+string(os.PathSeparator)+childFileInfo.Name(), openErr)
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

func GetExtensionFromMIMEType(mimeType string) string {
	if sMIMEExtensionMap == nil {
		sMIMEExtensionMap = map[string]string{}

		for key, value := range sExtensionMIMEMap {
			sMIMEExtensionMap[value] = key
		}
	}

	ret := "dat"

	if ext, exist := sMIMEExtensionMap[mimeType]; exist {
		ret = ext
	}

	if strings.HasPrefix(ret, ".") {
		ret = ret[1:]
	}

	return ret
}

func GetParent(itemPath string) string {
	parent := ""
	pathSeparators := []string{"\\", "/"}
	foundIndex := -1
	exist := false

	for _, pathSeparator := range pathSeparators {
		index := strings.LastIndex(itemPath, pathSeparator)
		if index >= 0 {
			foundIndex = index
			exist = true
			break
		}
	}

	if exist {
		parent = itemPath[0:foundIndex]
		parent = strings.TrimLeft(parent, " \t")

		if len(parent) == 0 {
			parent = "."
		}
	} else {
		// only file name
		parent = "."
	}

	return parent
}

func PrepareOutputFilepath(inputTopFolderpath, inputFilepath, outputTopFolderpath, distExtension string) string {
	ret := ""
	specifiedFolder := inputTopFolderpath

	if len(specifiedFolder) > 0 {
		ret = outputTopFolderpath + string(os.PathSeparator) + inputFilepath[len(inputTopFolderpath):]
	} else {
		ret = outputTopFolderpath + string(os.PathSeparator) + filepath.Base(inputFilepath)
	}

	if len(distExtension) > 0 {
		ret = ChangeExtension(ret, distExtension)
	}

	newFileParentFolder := ret[0 : len(ret)-len(filepath.Base(ret))]
	if !IsExist(newFileParentFolder) {
		os.MkdirAll(newFileParentFolder, 0755)
	}

	return ret
}

func Open(filePath string) (*os.File, error) {
	newFilepath := ExchangePath(filePath)
	return os.Open(newFilepath)
}

func OpenFile(filePath string, flag int, perm os.FileMode) (*os.File, error) {
	newFilepath := ExchangePath(filePath)
	return os.OpenFile(newFilepath, flag, perm)
}

func SkipUTF8BOM(readseeker io.ReadSeeker) (err error) {
	if _, seekErr := readseeker.Seek(0, io.SeekStart); seekErr == nil {
		readBuffer := make([]byte, 3)
		if _, readErr := readseeker.Read(readBuffer); readErr == nil {
			if bytes.Equal(readBuffer, []byte{0xEF, 0xBB, 0xBF}) {
				// UTF-8 BOM, no operation
			} else {
				_, err = readseeker.Seek(0, io.SeekStart)
			}
		}
	} else {
		err = seekErr
	}

	return err
}

func ReadFile(filePath string) ([]byte, error) {
	newFilepath := ExchangePath(filePath)
	return ioutil.ReadFile(newFilepath)
}

func ExchangePath(originalPath string) string {
	oldnewPathSeparaters := []byte{}

	switch runtime.GOOS {
	case "windows":
		oldnewPathSeparaters = append(oldnewPathSeparaters, '/')
		oldnewPathSeparaters = append(oldnewPathSeparaters, '\\')
		break
	case "darwin", "linux", "freebsd":
		oldnewPathSeparaters = append(oldnewPathSeparaters, '\\')
		oldnewPathSeparaters = append(oldnewPathSeparaters, '/')
		break
	}

	return exchangePath(originalPath, oldnewPathSeparaters)
}

func exchangePath(originalPath string, oldnewPathSeparaters []byte) string {
	ret := ""
	absolute := false

	if len(oldnewPathSeparaters) >= 2 {
		if strings.Index(originalPath, string(oldnewPathSeparaters[0])) >= 0 {
			if strings.HasPrefix(originalPath, string(oldnewPathSeparaters[0])) {
				absolute = true

				if oldnewPathSeparaters[0] == '\\' {
					// set default drive letter
					originalPath = "/mnt/c" + originalPath
				} else if strings.HasPrefix(originalPath, "/mnt/") {
					originalPath = string(originalPath[len("/mnt/")]) + ":" + originalPath[len("/mnt/")+1:]
				}
			} else if len(originalPath) > 1 && strings.HasPrefix(originalPath[1:], ":") {
				absolute = true

				originalPath = strings.ToLower(originalPath[0:1]) + originalPath[1:]
				originalPath = strings.Replace(originalPath, ":", "", 1)
				originalPath = "/mnt/" + originalPath
			}
		} else if strings.Index(originalPath, string(oldnewPathSeparaters[1])) >= 0 {
			// nop
			ret = originalPath
		} else {
			// current folder's file
		}

		if len(ret) == 0 {
			if absolute {
				tempPath := strings.Replace(originalPath, string(oldnewPathSeparaters[0]), string(oldnewPathSeparaters[1]), -1)
				ret = tempPath
			} else {
				ret = strings.Replace(originalPath, string(oldnewPathSeparaters[0]), string(oldnewPathSeparaters[1]), -1)
			}
		}
	}

	return ret
}

func GetRealFilepath(filePath string, ignoreNameCase, ignoreExtCase bool) (outputFilepath string, exist bool) {
	outputFilepath = filePath

	if ignoreNameCase || ignoreExtCase {
		fileName := filepath.Base(filePath)
		fileExt := filepath.Ext(fileName)
		parentPath := filePath[0 : len(filePath)-len(fileName)]
		fileName = fileName[0 : len(fileName)-len(fileExt)]

		if strings.HasSuffix(fileName, ".") {
			fileName = fileName[0 : len(fileName)-1]
		}

		if ignoreNameCase {
			fileName = strings.ToLower(fileName)
		}
		if ignoreExtCase {
			fileExt = strings.ToLower(fileExt)
		}

		if items, readErr := ioutil.ReadDir(parentPath); readErr == nil {
			for _, item := range items {
				itemName := item.Name()
				itemExt := filepath.Ext(itemName)
				itemName = itemName[0 : len(itemName)-len(itemExt)]

				if ignoreNameCase {
					itemName = strings.ToLower(itemName)
				}
				if ignoreExtCase {
					itemExt = strings.ToLower(itemExt)
				}

				if itemName == fileName && itemExt == fileExt {
					outputFilepath = parentPath + item.Name()
					exist = true
					break
				}
			}
		} else {
			LogfE("fail to readdir: %s,%v", parentPath, readErr)
		}
	} else {
		if _, statErr := os.Stat(filePath); statErr == nil {
			exist = true
		}
	}

	LogfV("output: %s, exist: %t", outputFilepath, exist)
	return
}

func JoinPath(base string, separator string, part string, otherParts ...string) string {
	builder := StringBuilder{}
	builder.Append(base)
	parts := []string{part}

	if otherParts != nil && len(otherParts) > 0 {
		parts = append(parts, otherParts...)
	}

	for _, part := range parts {
		builtText := builder.String()

		hasSuffix := strings.HasSuffix(builtText, separator)
		hasPrefix := strings.HasPrefix(part, separator)

		if (hasSuffix && !hasPrefix) || (!hasSuffix && hasPrefix) {
			builder.Append(part)
		} else if hasSuffix && hasPrefix {
			builder.Append(part[len(separator):])
		} else {
			builder.Append(separator).Append(part)
		}
	}

	return builder.String()
}
