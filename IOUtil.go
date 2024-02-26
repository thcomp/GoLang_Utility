package utility

import (
	"fmt"
	"io"
)

type NopCloser struct {
	target interface{}
}

func NewNopCloser(target interface{}) *NopCloser {
	if _, assertionOK := target.(io.Reader); assertionOK {
		return &NopCloser{
			target: target,
		}
	} else if _, assertionOK := target.(io.Seeker); assertionOK {
		return &NopCloser{
			target: target,
		}
	} else if _, assertionOK := target.(io.Writer); assertionOK {
		return &NopCloser{
			target: target,
		}
	}

	return &NopCloser{}
}

func (closer *NopCloser) Close() (err error) {
	if hiddenCloser, assertionOK := closer.target.(io.Closer); assertionOK {
		err = hiddenCloser.Close()
	}

	closer.target = nil
	return err
}

func (closer *NopCloser) Read(p []byte) (n int, err error) {
	if closer.target != nil {
		if reader, assertionOK := closer.target.(io.Reader); assertionOK {
			n, err = reader.Read(p)
		} else {
			err = fmt.Errorf("target not io.Reader")
		}
	} else {
		err = fmt.Errorf("target not opened")
	}

	return
}

func (closer *NopCloser) Seek(offset int64, whence int) (n int64, err error) {
	if closer.target != nil {
		if seeker, assertionOK := closer.target.(io.Seeker); assertionOK {
			n, err = seeker.Seek(offset, whence)
		} else {
			err = fmt.Errorf("fail to assert to io.Seeker")
		}
	} else {
		err = fmt.Errorf("target not opened")
	}

	return
}

func (closer *NopCloser) Write(p []byte) (n int, err error) {
	if closer.target != nil {
		if writer, assertionOK := closer.target.(io.Writer); assertionOK {
			n, err = writer.Write(p)
		} else {
			err = fmt.Errorf("fail to assert to io.Writer")
		}
	} else {
		err = fmt.Errorf("target not opened")
	}

	return
}
