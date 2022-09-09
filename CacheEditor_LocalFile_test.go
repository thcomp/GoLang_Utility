package utility

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func Test_LocalFileCacheEditor(t *testing.T) {
	if factory, err := NewLocalFileCacheEditorFactory(fmt.Sprintf(".%scache", string(os.PathSeparator))); err == nil {
		for i := 0; i < 100; i++ {
			if editor, openErr := factory.OpenCacheEditor(fmt.Sprintf("%d", i), os.O_CREATE|os.O_RDWR, os.FileMode(0600)); openErr == nil {
				appendText := fmt.Sprintf("%02X", i)
				textBuilder := StringBuilder{}

				for j := 0; j < (i + 10); j++ {
					textBuilder.Append(appendText)
					readSize, writeErr := editor.Write([]byte(appendText))
					if readSize != len(appendText) {
						t.Fatalf("write size is not matched: %d vs %d", readSize, len(appendText))
					}

					if writeErr != nil {
						t.Fatalf("fail to write: %v", writeErr)
					}
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

				if _, statErr := os.Stat(fmt.Sprintf(".%scache%s%d", string(os.PathSeparator), string(os.PathSeparator), i)); statErr == nil {
					t.Fatalf("error is not occurred")
				} else if !os.IsNotExist(statErr) {
					t.Fatalf("cache file not removed")
				}
			} else {
				t.Fatalf("fail to open: %v", openErr)
			}
		}
	}
}
