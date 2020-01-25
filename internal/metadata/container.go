package metadata

type Container struct {
	File *File
	Exif *Exif
	Tags []*Tag
}

type TagType int

const (
	TagTypeCustom TagType = 0
)

type Tag struct {
	Type  TagType
	Value string
}

func ReadFromFile(filepathArgument string) (*Container, error) {
	container := new(Container)
	e, err := exifReadFromFile(filepathArgument)
	if err == nil {
		container.Exif = e
	} else {
		container.Exif = &Exif {}
	}


	f, err := fileReadFromFile(filepathArgument)
	if err == nil {
		container.File = f
	}  else {
		container.File = &File{}
	}

	return container, nil
}
