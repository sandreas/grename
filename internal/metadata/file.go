package metadata

import (
	"github.com/gabriel-vasile/mimetype"
	"path/filepath"
	"strings"
)

type FileMeta struct {
	Extension string
	MimeType  *mimetype.MIME
}

func (f *FileMeta) ReadFromFile(filename string) error {
	f.Extension = strings.TrimPrefix(filepath.Ext(filename), ".")
	mime, err := mimetype.DetectFile(filename)
	if err == nil {
		f.MimeType = mime
	}
	return err
}
