package utility

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type ReadHelper struct {
	reader      io.Reader
	readCloser  io.ReadCloser
	lineBuffer  []string
	lineBuilder StringBuilder
}

func NewReadHelper(reader io.Reader) *ReadHelper {
	ret := ReadHelper{
		reader: reader,
	}

	return &ret
}

func NewReadHelperFromReadCloser(reader io.ReadCloser) *ReadHelper {
	ret := NewReadHelper(reader)
	ret.readCloser = reader
	return ret
}

func NewReadHelperFromFile(filePath string) (ret *ReadHelper, retErr error) {
	if inFileReader, openErr := os.Open(filePath); openErr == nil {
		ret = &ReadHelper{
			reader:     inFileReader,
			readCloser: inFileReader,
		}
	} else {
		retErr = openErr
	}

	return
}

func (helper *ReadHelper) ReadLine() (ret string, retErr error) {
	if helper.reader != nil {

		if helper.lineBuffer == nil {
			helper.lineBuffer = []string{}
		}

		if len(helper.lineBuffer) == 0 {
			buffer := make([]byte, 1024)

			for {
				LogfV("left line: %s", helper.lineBuilder.String())

				lines := []string(nil)
				leftText := ""
				brokenLoop := false

				size, readErr := helper.reader.Read(buffer)
				if size > 0 && readErr == nil {
					helper.lineBuilder.Append(string(buffer[0:size]))
					lines, leftText = separateLines(helper.lineBuilder.String())
				} else if readErr == io.EOF {
					retErr = readErr
					helper.lineBuilder.Append(string(buffer[0:size]))
					lines, leftText = separateLines(helper.lineBuilder.String())
					brokenLoop = true
				} else {
					LogfE("fail to read: %v", readErr)
					retErr = readErr
					break
				}

				helper.lineBuilder.Delete()

				if len(leftText) > 0 {
					helper.lineBuilder.Append(leftText)
				}

				if len(lines) > 0 {
					helper.lineBuffer = append(helper.lineBuffer, lines...)
					break
				} else {
					if brokenLoop {
						break
					}
				}
			}
		}

		if len(helper.lineBuffer) > 0 {
			ret = helper.lineBuffer[0]
			helper.lineBuffer = helper.lineBuffer[1:]
		}
	} else {
		retErr = fmt.Errorf("reader is not set, already closed?")
	}

	return
}

func (helper *ReadHelper) Close() (ret error) {
	if helper.readCloser != nil {
		ret = helper.readCloser.Close()
		helper.readCloser = nil
	}
	helper.reader = nil

	return
}

func separateLines(buffer string) (lines []string, leftText string) {
	tempLines := strings.Split(buffer, "\r\n")

	if len(tempLines) == 1 && len(tempLines[0]) == len(buffer) {
		// not found "\r\n"
		leftText = buffer
	} else {
		for index, tempLine := range tempLines {
			LogfV("temp line: %s", tempLine)
			if strings.HasSuffix(buffer, "\r\n") {
				if len(tempLine) > 0 {
					lines = append(lines, tempLine)
				}
				leftText = ""
			} else {
				if index == len(tempLines)-1 {
					leftText = tempLine
				} else {
					lines = append(lines, tempLine)
					leftText = ""
				}
			}
		}
	}

	return
}
