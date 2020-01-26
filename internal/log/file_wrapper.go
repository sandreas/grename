package log

import (
	"os"
)

func NewFileWrapper(filename string) *FileWrapper {
	return &FileWrapper{
		filename,
		nil,
		nil,
	}
}

type FileWrapper struct {
	filename    string
	filePointer *os.File
	err         error
}

func (fw *FileWrapper) Write(b []byte) (int, error) {
	if fw.err != nil {
		return -1, fw.err
	}
	if fw.filePointer == nil {
		f, err := os.OpenFile(fw.filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fw.err = err
			return -1, err
		}
		fw.filePointer = f
	}
	return fw.filePointer.Write(b)
}

func (fw *FileWrapper) CurrentFile() *os.File {
	return fw.filePointer
}

func (fw *FileWrapper) Flush() error {
	if fw.filePointer != nil {
		return fw.filePointer.Close()
	}
	return nil
}
