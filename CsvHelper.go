package utility

import (
	"encoding/csv"
	"fmt"
	"io"
	"reflect"
)

type CsvReader struct {
	csvReader         csv.Reader
	headerOnFirstLine bool
	posFieldNameMap   map[int]string
	linePosition      int
}

func NewCsvReader(reader io.Reader, headerOnFirstLine bool) *CsvReader {
	return &CsvReader{
		csvReader:         *csv.NewReader(reader),
		headerOnFirstLine: headerOnFirstLine,
		linePosition:      0,
	}
}

func (reader *CsvReader) DecodeLine(recvr interface{}) error {
	retErr := error(nil)

	if reader.linePosition == 0 {
		reader.posFieldNameMap = map[int]string{}

		if reader.headerOnFirstLine {
			if fields, readErr := reader.csvReader.Read(); readErr == nil {
				for pos, field := range fields {
					reader.posFieldNameMap[pos] = field
				}
			} else {
				retErr = readErr
			}

			reader.linePosition++
		}
	}

	if retErr == nil {
		interfaceHelper := NewInterfaceHelper(recvr)
		if interfaceHelper.kind == reflect.Pointer {
			reflectHelper := NewReflectHelper(recvr)

			if fields, readErr := reader.csvReader.Read(); readErr == nil {
				for pos, field := range fields {
					switch interfaceHelper.GetKind() {
					case reflect.Array, reflect.Slice:
						reflectHelper.SetOnList(pos, field)
					case reflect.Struct:
						if fieldTagName, exist := reader.posFieldNameMap[pos]; exist {
							if !reflectHelper.SetByTagName("csv", fieldTagName, field) {
								if !reflectHelper.SetByName(fieldTagName, field) {
									LogfE("fail to set value on structure: %s, %v", fieldTagName, field)
								}
							}
						}
					case reflect.Map:
						if fieldTagName, exist := reader.posFieldNameMap[pos]; exist {
							if !reflectHelper.SetOnMap(fieldTagName, field) {
								LogfE("fail to set value on map: %s, %v", fieldTagName, field)
							}
						}
					}
				}
			} else {
				retErr = readErr
			}
		} else {
			retErr = fmt.Errorf("recvr is not pointer")
		}
	}

	return retErr
}
