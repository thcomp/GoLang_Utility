package utility

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/xid"
)

const c_TemporaryFolderPath = "." + string(os.PathSeparator) + "multipart_helper" + string(os.PathSeparator)

var sLocalCacheEditorFactory *LocalFileCacheEditorFactory = nil

type FormData struct {
	cacheEditor CacheEditor
	mimeType    string
}

func (formData *FormData) Part() (ret []byte, retErr error) {
	buffer := bytes.NewBuffer([]byte{})
	readBuffer := make([]byte, 10*1024)

	formData.cacheEditor.Seek(0, io.SeekStart)
	for {
		if readSize, readErr := formData.cacheEditor.Read(readBuffer); readSize > 0 {
			if _, writeErr := buffer.Write(readBuffer[0:readSize]); writeErr != nil {
				retErr = writeErr
				break
			}

			if readErr == io.EOF {
				break
			} else {
				retErr = readErr
			}
		} else if readErr == io.EOF {
			break
		} else {
			retErr = readErr
			break
		}
	}

	if retErr == nil {
		ret = buffer.Bytes()
	}

	return
}

func (formData *FormData) Filename() string {
	return formData.cacheEditor.ID()
}

func (formData *FormData) MimeType() string {
	return formData.mimeType
}

type MultipartHelper struct {
	formDataMap        map[string](*FormData)
	cacheEditorFactory CacheEditorFactory
}

func NewMultipartHelperFromHttpRequest(r *http.Request) (*MultipartHelper, error) {
	if sLocalCacheEditorFactory == nil {
		sLocalCacheEditorFactory, _ = NewLocalFileCacheEditorFactory(c_TemporaryFolderPath)
	}
	return NewMultipartHelperFromHttpRequestWithFactory(r, sLocalCacheEditorFactory)
}

func NewMultipartHelperFromHttpRequestWithFactory(r *http.Request, cacheEditorFactory CacheEditorFactory) (*MultipartHelper, error) {
	ret := (*MultipartHelper)(nil)
	retErr := error(nil)

	if mediaType, params, parseErr := mime.ParseMediaType(r.Header.Get(`Content-Type`)); parseErr == nil {
		if strings.HasPrefix(mediaType, `multipart/`) {
			ret, retErr = NewMultipartHelper(r.Body, params[`boundary`], cacheEditorFactory)
		} else {
			retErr = fmt.Errorf("not support type: %s", mediaType)
		}
	} else {
		retErr = parseErr
	}

	return ret, retErr
}

func NewMultipartHelper(reader io.Reader, boundary string, cacheEditorFactory CacheEditorFactory) (*MultipartHelper, error) {
	ret := &MultipartHelper{cacheEditorFactory: cacheEditorFactory}
	retError := error(nil)

	multipartReader := multipart.NewReader(reader, boundary)
	ret.formDataMap = map[string](*FormData){}

	sameFormNameIndicateMap := map[string]int{}
	for {
		if tempPart, readErr := multipartReader.NextPart(); readErr == nil {
			if partBytes, readErr := ioutil.ReadAll(tempPart); readErr == nil {
				temporaryFilepath := xid.New().String()
				if cacheEditor, err := ret.cacheEditorFactory.OpenCacheEditor(temporaryFilepath, os.O_WRONLY, 0400); err == nil {
					if _, writeErr := cacheEditor.Write(partBytes); writeErr == nil {
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
							cacheEditor: cacheEditor,
							mimeType:    mimeType,
						}
					} else {
						retError = writeErr
						LogfE("fail to write part: %v", retError)
						break
					}
				} else {
					retError = err
					LogfE("fail to write part: %v", retError)
					break
				}
			} else {
				retError = readErr
				LogfE("fail to read part: %v", retError)
				break
			}
		} else if readErr == io.EOF {
			break
		} else {
			// error
			retError = readErr
			LogfE("fail to get part: %v", retError)
			break
		}
	}

	return ret, retError
}

func NewMultipartWriteHelper(cacheEditorFactory CacheEditorFactory) (ret *MultipartHelper) {
	ret = &MultipartHelper{
		formDataMap:        map[string](*FormData){},
		cacheEditorFactory: cacheEditorFactory,
	}

	if cacheEditorFactory == nil {
		if sLocalCacheEditorFactory == nil {
			sLocalCacheEditorFactory, _ = NewLocalFileCacheEditorFactory(c_TemporaryFolderPath)
		}
		ret.cacheEditorFactory = sLocalCacheEditorFactory
	}

	return ret
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
	if len(helper.formDataMap) > 0 {
		for _, formData := range helper.formDataMap {
			if formData.cacheEditor != nil {
				formData.cacheEditor.Close()
				formData.cacheEditor.Remove()
				formData.cacheEditor = nil
			}
		}
	}

	return helper.cacheEditorFactory.Close()
}

func (helper *MultipartHelper) AppendFormData(formName, mimeType string, data []byte) error {
	retErr := error(nil)
	temporaryFilepath := xid.New().String()

	if cacheEditor, err := helper.cacheEditorFactory.OpenCacheEditor(temporaryFilepath, os.O_WRONLY, 0400); err == nil {
		if _, retErr = cacheEditor.Write(data); retErr == nil {
			helper.formDataMap[formName] = &FormData{
				cacheEditor: cacheEditor,
				mimeType:    mimeType,
			}
		}
	} else {
		retErr = err
	}

	return retErr
}

func (helper *MultipartHelper) Write(writer io.Writer) (boundary string, retErr error) {
	multipartWriter := multipart.NewWriter(writer)
	defer multipartWriter.Close()

	for formName, formData := range helper.formDataMap {
		formPartWriter := (io.Writer)(nil)
		if formData.mimeType != "" {
			quoteEscaper := strings.NewReplacer("\\", "\\\\", `"`, "\\\"")
			formName = quoteEscaper.Replace(formName)

			header := textproto.MIMEHeader{}
			header.Set("Content-Type", formData.mimeType)
			header.Set("Content-Disposition", fmt.Sprintf("form-data; name=\"%s\"; filename=\"%s\"", formName, filepath.Base(formData.Filename())))

			if partWriter, createErr := multipartWriter.CreatePart(header); createErr == nil {
				formPartWriter = partWriter
			} else {
				retErr = createErr
			}
		} else {
			formPartWriter, retErr = multipartWriter.CreateFormField(formName)
		}

		if formPartWriter == nil {
			if retErr == nil {
				retErr = fmt.Errorf("internal server error")
			}
			break
		} else {
			if partData, openErr := formData.Part(); openErr == nil {
				if _, writeErr := formPartWriter.Write(partData); writeErr != nil {
					retErr = writeErr
				}
			} else {
				retErr = openErr
			}
		}

		if retErr != nil {
			break
		}
	}

	return multipartWriter.Boundary(), retErr
}
