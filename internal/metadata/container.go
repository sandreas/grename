package metadata

import "github.com/sandreas/log"

type MetaContainer struct {
	File *FileMeta
	Exif *Exif
}

func (container *MetaContainer) ReadFromFile(filename string) error {
	var err error

	container.Exif = &Exif{}
	err = container.Exif.ReadFromFile(filename)
	if err != nil {
		log.Warnf("could not read exif from file %s: %s", filename, err.Error())
	}

	container.File = &FileMeta{}
	err = container.File.ReadFromFile(filename)
	if err != nil {
		log.Warnf("could not file meta from file %s: %s", filename, err.Error())
	}
	return err
}
