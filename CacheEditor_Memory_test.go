package utility

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func Test_MemoryCacheEditor(t *testing.T) {
	factory := NewMemoryCacheEditorFactory()
	limitSize := 100
	factory.SetLimitCacheSize(int64(limitSize))

	writtenSize := 0
	for i := 0; i < 100; i++ {
		if editor, openErr := factory.OpenCacheEditor(fmt.Sprintf("%d", i), os.O_CREATE|os.O_RDWR, os.FileMode(0600)); openErr == nil {
			appendText := fmt.Sprintf("%02X", i)
			textBuilder := StringBuilder{}

			for j := 0; j < (i + 10); j++ {
				writeSize, writeErr := editor.Write([]byte(appendText))
				if writeErr != nil {
					if writtenSize+len(appendText) <= limitSize {
						t.Fatalf("fail to write: %v, %d vs %d", writeErr, writtenSize+len(appendText), limitSize)
					} else {
						limitSize += limitSize
						factory.SetLimitCacheSize(int64(limitSize))
						continue
					}
				} else {
					if writeSize != len(appendText) {
						t.Fatalf("write size is not matched: i=%d, j=%d, %d vs %d", i, j, writeSize, len(appendText))
					}
				}

				textBuilder.Append(appendText)
				writtenSize += writeSize
			}

			editor.Seek(0, io.SeekStart)
			if data, readErr := ioutil.ReadAll(editor); readErr == nil {
				if string(data) != textBuilder.String() {
					t.Fatalf("not matched: %s vs %s", string(data), textBuilder.String())
				}
			} else {
				t.Fatalf("fail to read: %v", readErr)
			}

			editor.Seek(0, io.SeekStart)
			if data, readErr := ioutil.ReadAll(editor); readErr == nil {
				if string(data) != textBuilder.String() {
					t.Fatalf("not matched: %s vs %s", string(data), textBuilder.String())
				}
			} else {
				t.Fatalf("fail to read: %v", readErr)
			}

			if closeErr := editor.Close(); closeErr != nil {
				t.Fatalf("fail to close: %v", closeErr)
			}

			if removeErr := editor.Remove(); removeErr != nil {
				t.Fatalf("fail to remove: %v", removeErr)
			}
		} else {
			t.Fatalf("fail to open: %v", openErr)
		}
	}
}
