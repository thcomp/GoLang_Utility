package utility

import (
	"fmt"
	"io"
)

type NopCloser struct {
	reader io.Reader
}

func NewNopCloser(reader io.Reader) *NopCloser {
	return &NopCloser{
		reader: reader,
	}
}

func (closer *NopCloser) Close() error {
	closer.reader = nil
	return nil
}

func (closer *NopCloser) Read(p []byte) (n int, err error) {
	if closer.reader != nil {
		n, err = closer.Read(p)
	} else {
		err = fmt.Errorf("reader not opened")
	}

	return
}
