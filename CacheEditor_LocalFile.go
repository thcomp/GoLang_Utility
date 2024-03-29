package utility

import (
	"fmt"
	"os"
)

type LocalFileCacheEditorFactory struct {
	cacheRootFolderPath string
}

type LocalFileCache struct {
	localFile     *os.File
	localFilePath string
}

func NewLocalFileCacheEditorFactory(cacehRootFolder string) (ret *LocalFileCacheEditorFactory, retErr error) {
	if rootFileInfo, statErr := os.Stat(cacehRootFolder); statErr == nil {
		if rootFileInfo.IsDir() {
			ret = &LocalFileCacheEditorFactory{
				cacheRootFolderPath: cacehRootFolder,
			}
		} else {
			retErr = fmt.Errorf("%s has been exist by file", cacehRootFolder)
		}
	} else if os.IsNotExist(statErr) {
		if mkdirErr := os.MkdirAll(cacehRootFolder, 0700); mkdirErr == nil {
			if _, statErr := os.Stat(cacehRootFolder); statErr == nil {
				ret = &LocalFileCacheEditorFactory{
					cacheRootFolderPath: cacehRootFolder,
				}
			} else {
				retErr = statErr
			}
		} else {
			retErr = mkdirErr
		}
	} else {
		retErr = statErr
	}

	return ret, retErr
}

func (factory *LocalFileCacheEditorFactory) OpenCacheEditor(relPathUnderCacheRoot string, flag int, mode interface{}) (CacheEditor, error) {
	path := JoinPath(factory.cacheRootFolderPath, string(os.PathSeparator), relPathUnderCacheRoot)
	infHelper := NewInterfaceHelper(mode)
	if modeNum, getRet := infHelper.GetNumber(); getRet {
		perm := os.FileMode(modeNum)
		localFile, openErr := os.OpenFile(path, os.O_CREATE|os.O_RDWR, perm)
		if openErr == nil {
			return &LocalFileCache{localFilePath: path, localFile: localFile}, nil
		} else {
			return nil, fmt.Errorf("fail to open file: %v", openErr)
		}
	} else {
		return nil, fmt.Errorf("mode is not FileMode: %v", mode)
	}
}

func (factory *LocalFileCacheEditorFactory) Close() error {
	return os.RemoveAll(factory.cacheRootFolderPath)
}

func (file *LocalFileCache) Read(p []byte) (int, error) {
	return file.localFile.Read(p)
}

func (file *LocalFileCache) Seek(offset int64, whence int) (n int64, err error) {
	return file.localFile.Seek(offset, whence)
}

func (file *LocalFileCache) Write(p []byte) (n int, err error) {
	return file.localFile.Write(p)
}

func (file *LocalFileCache) Close() error {
	return file.localFile.Close()
}

func (file *LocalFileCache) Remove() error {
	file.localFile.Close()
	return os.Remove(file.localFilePath)
}

func (file *LocalFileCache) ID() string {
	return file.localFilePath
}
