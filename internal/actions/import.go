package actions

import (
	"github.com/gabriel-vasile/mimetype"
	"github.com/urfave/cli/v2"
	"grename/internal/database"
	"grename/internal/metadata"
	"os"
	"path/filepath"
	"strings"
)

type Import struct {
}

func (action *Import) Execute(c *cli.Context) error {
	var err error

	source := filepath.Clean(c.Args().First())
	destination := filepath.Clean(c.Args().Get(1))
	mediaTypes := strings.Split(c.String("include-media-types"), ",")
	// renameTemplate := c.String("tpl")

	importFiles := []*database.File{}

	dbPath := destination + "/grename.db"
	db, err := database.InitDatabase(&database.Credentials{
		Driver:   "sqlite3",
		Database: dbPath,
	})
	if err != nil {
		return err
	}

	err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		mime, err := mimetype.DetectFile(path)
		if err != nil {
			return nil
		}

		mimeParts := strings.Split(mime.String(), "/")
		if len(mimeParts) != 2 {
			return nil
		}

		fileMediaType := mimeParts[0]

		matchFound := false
		for _, includedMediaType := range mediaTypes {
			if fileMediaType == includedMediaType {
				matchFound = true
				break
			}
		}

		if !matchFound {
			return nil
		}

		importFiles = append(importFiles, &database.File{
			MimeMediaType: fileMediaType,
			MimeSubType:   mimeParts[1],
			Hash:          "",
			Location:      path,
			Tags:          nil,
		})

		return nil
	})

	for _, f := range importFiles {

		m, err := metadata.ReadFromFile(f.Location)
		if err != nil {
			return err
		}

		tagBasePath := filepath.ToSlash(strings.TrimPrefix(filepath.Dir(f.Location), source))
		rawTagStrings := strings.Split(tagBasePath, "/")
		tagStrings := UniqueStrings(FilterEmptyStrings(rawTagStrings))
		db.FirstOrCreate(f)

		for _, str := range tagStrings {
			tag := &database.Tag{
				Value: str,
			}
			fileTag := &database.FileTag{
				Tag:   tag,
				File:  f,
				Group: database.TagGroupDefault,
				Type:  database.TagTypeDefault,
			}

			db.FirstOrCreate(tag).FirstOrCreate(fileTag)
		}

		//var f models.File
		//db.First(f)
		//println("renameTemplate: " + renameTemplate)
		//println("source: " + source)
		println("destination: " + f.Location)
		println(m)
		//println(m.File.MimeType.String())
		//println(m.Exif.Make)
		//println(m.Exif.Model)
		//println(m.Exif.DateTime.Format("2006-01-02 03:04:05"))

	}

	return err

}

// Ints returns a unique subset of the int slice provided.
func UniqueStrings(input []string) []string {
	u := make([]string, 0, len(input))
	m := make(map[string]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}

func FilterEmptyStrings(strings []string) []string {
	var ln int
	for _, str := range strings {
		if str == "" {
			continue // drop IPv4 address
		}
		strings[ln] = str
		ln++
	}
	return strings[:ln]
}
