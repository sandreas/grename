package metadata

import (
	"path/filepath"
	"strings"
)

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
	}

	f, err := fileReadFromFile(filepathArgument)
	if err == nil {
		container.File = f
	}

	tagStrings := UniqueStrings(FilterEmptyStrings(strings.Split(filepath.ToSlash(filepath.Dir(filepathArgument)), "/")))
	for _, str := range tagStrings {
		container.Tags = append(container.Tags, &Tag{
			Type:  TagTypeCustom,
			Value: str,
		})
	}

	return container, nil
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
