package utility

import (
	"io"
)

type CacheEditor interface {
	io.Reader
	io.Writer
	io.Seeker
	io.Closer
	Remove() error
	ID() string
}

type CacheEditorFactory interface {
	OpenCacheEditor(path string, flag int, mode interface{}) (CacheEditor, error)
}
