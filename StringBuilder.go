package utility

import(
	"bytes"
//	"log"
);

var EnableDebug bool = true;

type StringBuilderIF interface {
	Append(text string) StringBuilder;
	Delete();
	String() string;
};

type StringBuilder struct {
	mBuffer *bytes.Buffer;
};

func (this *StringBuilder) Append(text string) *StringBuilder {
	if(this.mBuffer == nil){
		this.mBuffer = bytes.NewBufferString("");
	}

	this.mBuffer.WriteString(text);
	return this;
}

func (this *StringBuilder) Delete(){
	this.mBuffer = bytes.NewBufferString("");
}

func (this *StringBuilder) String() string{
	var ret string;

	if(this.mBuffer != nil){
		ret = this.mBuffer.String();
	}else{
		ret = ``;
	}

	return ret;
}
