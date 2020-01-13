package metadata

import (
	"github.com/gabriel-vasile/mimetype"
	"path/filepath"
	"strings"
)

type File struct {
	Extension string
	MimeType  *mimetype.MIME
}

func fileReadFromFile(filepathArgument string) (*File, error) {
	f := new(File)
	f.Extension = strings.TrimPrefix(filepath.Ext(filepathArgument), ".")
	mime, err := mimetype.DetectFile(filepathArgument)
	if err == nil {
		f.MimeType = mime
	}
	return f, err
}
