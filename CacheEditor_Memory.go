package utility

import (
	"bytes"
	"fmt"
	"io"
)

type MemoryCacheEditorFactory struct {
	totalCacheSize int64
	limitCacheSize int64
	cacheMap       map[string](*iMemoryBuffer)
}

var sMemoryCacheEditorFactory *MemoryCacheEditorFactory

func NewMemoryCacheEditorFactory() *MemoryCacheEditorFactory {
	if sMemoryCacheEditorFactory == nil {
		sMemoryCacheEditorFactory = &MemoryCacheEditorFactory{limitCacheSize: -1}
	}
	return sMemoryCacheEditorFactory
}

func (factory *MemoryCacheEditorFactory) SizeChanged(id string, diffSize int64) {
	factory.totalCacheSize += diffSize
}

func (factory *MemoryCacheEditorFactory) IsAppendCache(appendSize int64) (ret bool) {
	if factory.limitCacheSize > 0 {
		if factory.totalCacheSize+appendSize <= factory.limitCacheSize {
			ret = true
		}
	} else {
		ret = true
	}

	return ret
}

func (factory *MemoryCacheEditorFactory) Remove(id string) (retErr error) {
	delete(factory.cacheMap, id)
	return nil
}

func (factory *MemoryCacheEditorFactory) GetLimitCacheSize() int64 {
	return factory.limitCacheSize
}

func (factory *MemoryCacheEditorFactory) SetLimitCacheSize(limitCacheSize int64) {
	factory.limitCacheSize = limitCacheSize
}

func (factory *MemoryCacheEditorFactory) OpenLocalFileCacheEditor(id string, flag int, mode interface{}) (CacheEditor, error) {
	ret := CacheEditor(nil)
	retErr := error(nil)

	if factory.cacheMap == nil {
		factory.cacheMap = map[string](*iMemoryBuffer){}
	}

	if buffer, exist := factory.cacheMap[id]; exist {
		ret = buffer
	} else {
		buffer = newMemoryBuffer(id, factory)
		factory.cacheMap[id] = buffer
		ret = buffer
	}

	return ret, retErr
}

type iController interface {
	SizeChanged(id string, diffSize int64)
	IsAppendCache(appendSize int64) bool
	Remove(id string) error
}

type iMemoryBuffer struct {
	id               string
	buffer           bytes.Buffer
	offset           int64
	controller       iController
	needChangeBuffer bool
}

func newMemoryBuffer(id string, controller iController) *iMemoryBuffer {
	return &iMemoryBuffer{
		buffer:     *bytes.NewBuffer([]byte{}),
		controller: controller,
	}
}

func (buffer *iMemoryBuffer) Read(p []byte) (n int, err error) {
	reader := bytes.NewReader(buffer.buffer.Bytes())
	n, err = reader.ReadAt(p, int64(buffer.offset))
	if err == nil {
		buffer.offset += int64(n)
	}

	return n, err
}

func (buffer *iMemoryBuffer) Seek(offset int64, whence int) (n int64, err error) {
	switch whence {
	case io.SeekStart:
		buffer.offset = offset
	case io.SeekCurrent:
		buffer.offset += offset
	case io.SeekEnd:
		buffer.offset = int64(buffer.buffer.Len()) + offset
	}

	currentBufferSize := int64(buffer.buffer.Len())
	if buffer.offset < 0 || buffer.offset >= currentBufferSize {
		buffer.offset = currentBufferSize + (buffer.offset % currentBufferSize)
	}

	if currentBufferSize == buffer.offset {
		buffer.needChangeBuffer = false
	} else {
		buffer.needChangeBuffer = true
	}

	return
}

func (buffer *iMemoryBuffer) Write(p []byte) (n int, err error) {
	if buffer.needChangeBuffer {
		buffer.needChangeBuffer = false
		currentBufferSize := buffer.buffer.Len()
		buffer.buffer = *bytes.NewBuffer(buffer.buffer.Bytes()[0:buffer.offset])
		buffer.controller.SizeChanged(buffer.id, int64(buffer.buffer.Len()-currentBufferSize))
	}

	if buffer.controller.IsAppendCache(int64(len(p))) {
		n, err = buffer.buffer.Write(p)
		if err == nil {
			buffer.controller.SizeChanged(buffer.id, int64(n))
			buffer.offset += int64(n)
		}
	} else {
		err = fmt.Errorf("cache size is over the limit")
	}

	return
}

func (buffer *iMemoryBuffer) Close() error {
	buffer.buffer.Reset()
	return nil
}

func (buffer *iMemoryBuffer) Remove() error {
	ret := buffer.controller.Remove(buffer.id)
	buffer.Close()
	return ret
}

func (buffer *iMemoryBuffer) ID() string {
	return buffer.id
}
