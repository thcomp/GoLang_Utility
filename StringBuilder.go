package utility

import (
	"bytes"
	"fmt"
	//	"log"
)

var EnableDebug bool = true

type StringBuilderIF interface {
	Append(text string) StringBuilder
	AppendByte(text byte) StringBuilder
	Delete()
	String() string
	Length() int
	Bytes() []byte
}

type StringBuilder struct {
	mBuffer *bytes.Buffer
}

func (this *StringBuilder) Append(text string) *StringBuilder {
	if this.mBuffer == nil {
		this.mBuffer = bytes.NewBufferString("")
	}

	this.mBuffer.WriteString(text)
	return this
}

func (this *StringBuilder) Appendf(fmtText string, params ...interface{}) *StringBuilder {
	if this.mBuffer == nil {
		this.mBuffer = bytes.NewBufferString("")
	}

	this.mBuffer.WriteString(fmt.Sprintf(fmtText, params...))
	return this
}

func (this *StringBuilder) AppendByte(text byte) *StringBuilder {
	if this.mBuffer == nil {
		this.mBuffer = bytes.NewBufferString("")
	}

	this.mBuffer.WriteByte(text)
	return this
}

func (this *StringBuilder) Delete() {
	this.mBuffer = bytes.NewBufferString("")
}

func (this *StringBuilder) String() string {
	var ret string

	if this.mBuffer != nil {
		ret = this.mBuffer.String()
	} else {
		ret = ``
	}

	return ret
}

func (this *StringBuilder) Length() int {
	var ret int = 0

	if this.mBuffer != nil {
		ret = len(this.mBuffer.String())
	}

	return ret
}

func (this *StringBuilder) Bytes() []byte {
	var ret []byte = nil

	if this.mBuffer != nil {
		ret = this.mBuffer.Bytes()
	}

	return ret
}
