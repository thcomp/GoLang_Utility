package utility

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/rs/xid"
)

const c_TemporaryFolderPath = "multipart_helper" + string(os.PathSeparator)

type FormData struct {
	temporaryFilepath string
	originalFilename  string
	mimeType          string
}

func (formData *FormData) Part() ([]byte, error) {
	return ioutil.ReadFile(formData.temporaryFilepath)
}

func (formData *FormData) Filename() string {
	return formData.originalFilename
}

func (formData *FormData) MimeType() string {
	return formData.mimeType
}

type MultipartHelper struct {
	formDataMap map[string](*FormData)
}

func NewMultipartHelperFromHttpRequest(r *http.Request) (*MultipartHelper, error) {
	ret := (*MultipartHelper)(nil)
	retErr := error(nil)

	if mediaType, params, parseErr := mime.ParseMediaType(r.Header.Get(`Content-Type`)); parseErr == nil {
		if strings.HasPrefix(mediaType, `multipart/`) {
			ret, retErr = NewMultipartHelper(r.Body, params[`boundary`])
		} else {
			retErr = fmt.Errorf("not support type: %s", mediaType)
		}
	} else {
		retErr = parseErr
	}

	return ret, retErr
}

func NewMultipartHelper(reader io.Reader, boundary string) (*MultipartHelper, error) {
	ret := &MultipartHelper{}
	retError := error(nil)

	multipartReader := multipart.NewReader(reader, boundary)
	ret.formDataMap = map[string](*FormData){}

	os.MkdirAll(c_TemporaryFolderPath, 0700)

	sameFormNameIndicateMap := map[string]int{}
	for {
		if tempPart, readErr := multipartReader.NextPart(); readErr == nil {
			if partBytes, readErr := ioutil.ReadAll(tempPart); readErr == nil {
				temporaryFilepath := c_TemporaryFolderPath + xid.New().String()
				if writeErr := ioutil.WriteFile(temporaryFilepath, partBytes, 0600); writeErr == nil {
					contentDisposition := tempPart.Header.Get("Content-Disposition")
					fileName := ``
					if len(contentDisposition) > 0 {
						contentDispositionItemArray := strings.Split(contentDisposition, ";")
						for _, contentDispositionItem := range contentDispositionItemArray {
							trimedContentDispositionItem := strings.Trim(contentDispositionItem, " ")
							lowerTrimedContentDispositionItem := strings.ToLower(trimedContentDispositionItem)

							if strings.HasPrefix(lowerTrimedContentDispositionItem, "filename") {
								fileName = trimedContentDispositionItem[len("filename"):]
								fileName = strings.Trim(fileName, " ")
								fileName = strings.Trim(fileName, "=")
								fileName = strings.Trim(fileName, " ")
								fileName = strings.Trim(fileName, "\"")
							}
						}
					}
					mimeType := tempPart.Header.Get("Content-Type")
					formName := tempPart.FormName()
					if strings.HasSuffix(formName, "[]") {
						index := 0
						index, _ = sameFormNameIndicateMap[formName]
						index = index + 1
						sameFormNameIndicateMap[formName] = index

						formName = fmt.Sprintf("%s[%d]", formName[0:len(formName)-len("[]")], index)
					}

					ret.formDataMap[formName] = &FormData{
						temporaryFilepath: temporaryFilepath,
						originalFilename:  fileName,
						mimeType:          mimeType,
					}
				} else {
					retError = writeErr
					LogE(fmt.Sprintf("fail to write part: %s", retError.Error()))
					break
				}
			} else {
				retError = readErr
				LogE(fmt.Sprintf("fail to read part: %s", retError.Error()))
				break
			}
		} else if readErr == io.EOF {
			break
		} else {
			// error
			retError = readErr
			LogE(fmt.Sprintf("fail to get part: %s", retError.Error()))
			break
		}
	}

	return ret, retError
}

func (helper *MultipartHelper) PartNames() []string {
	ret := []string{}

	for partName, _ := range helper.formDataMap {
		ret = append(ret, partName)
	}

	return ret
}

func (helper *MultipartHelper) GetByName(partName string) (*FormData, error) {
	ret := (*FormData)(nil)
	retErr := error(nil)

	if formData, exist := helper.formDataMap[partName]; exist {
		ret = formData
	} else {
		retErr = fmt.Errorf("not found this part: %s", partName)
	}

	return ret, retErr
}

func (helper *MultipartHelper) GetByIndex(index int) (*FormData, error) {
	ret := (*FormData)(nil)
	retErr := error(nil)

	if len(helper.formDataMap) > index {
		for _, formData := range helper.formDataMap {
			if index == 0 {
				ret = formData
				break
			}
			index--
		}
	} else {
		retErr = fmt.Errorf("index over: %d of %d", index, len(helper.formDataMap))
	}

	return ret, retErr
}

func (helper *MultipartHelper) Count() int {
	return len(helper.formDataMap)
}

func (helper *MultipartHelper) Close() error {
	ret := error(nil)

	if len(helper.formDataMap) > 0 {
		for _, formData := range helper.formDataMap {
			if len(formData.temporaryFilepath) > 0 {
				if IsExist(formData.temporaryFilepath) {
					if ret = os.Remove(formData.temporaryFilepath); ret == nil {
						formData.temporaryFilepath = ""
					}
				} else {
					ret = fmt.Errorf("not found temporary file: %s", formData.temporaryFilepath)
					formData.temporaryFilepath = ""
				}
			}
		}
	}

	return ret
}
