package utility

import (
	"bytes"
	"net/http"
)

type HttpResponseHelper struct {
	httpRes   *http.Response
	resEntity *bytes.Buffer
}

func NewHttpResponseHelper(params ...interface{}) *HttpResponseHelper {
	httpRes := (*http.Response)(nil)
	if len(params) > 0 {
		for _, param := range params {
			if tempHttpRes, assertionOK := param.(*http.Response); assertionOK {
				httpRes = tempHttpRes
			}
		}
	}

	if httpRes == nil {
		httpRes = &http.Response{}
	}

	return &HttpResponseHelper{
		httpRes: httpRes,
	}
}

func (res *HttpResponseHelper) Header() http.Header {
	return res.httpRes.Header
}

func (res *HttpResponseHelper) Write(data []byte) (int, error) {
	if res.resEntity == nil {
		res.resEntity = bytes.NewBuffer([]byte{})
		res.httpRes.Body = NewNopCloser(res.resEntity)
	}

	return res.resEntity.Write(data)
}

func (res *HttpResponseHelper) WriteHeader(statusCode int) {
	res.httpRes.StatusCode = statusCode
}

func (res *HttpResponseHelper) Status(status string) {
	res.httpRes.Status = status
}

func (res *HttpResponseHelper) ExportHttpResponse() *http.Response {
	res.httpRes.Body = NewNopCloser(bytes.NewReader(res.resEntity.Bytes()))
	return res.httpRes
}
