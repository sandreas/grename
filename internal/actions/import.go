package actions

import (
	"bytes"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/sandreas/afero"
	"github.com/urfave/cli/v2"
	"grename/internal/crypt"
	"grename/internal/database"
	"grename/internal/log"
	"grename/internal/metadata"
	"lukechampine.com/blake3"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"text/template/parse"
)

type ImportSetting struct {
	MimeTypes    []string
	NameTemplate string
}

type ImportFile struct {
	Source        string
	ImportSetting *ImportSetting
	File          *database.File
}

// TODO
// - dry run
// - ignore all-tags on import
// - import --copy or import --move with --keep-duplicate-files and --ignore-iptc-keywords
// - skip-inclomplete-exif by default
// - type Importer struct with ImportSettings
// - type Finder struct with FindSettings (or Generic Settings)
// - find --add-tag and find --remove-tag
// - dump tags to iptc keywords (no lib available) - before hash remove iptc keywords
// - web / serve start a webserver with static

type Import struct {
}

func ListTemplFields(t *template.Template) []string {
	return listNodeFields(t.Tree.Root, nil)
}
func listNodeFields(node parse.Node, res []string) []string {
	if node.Type() == parse.NodeAction {
		res = append(res, node.String())
	}

	if ln, ok := node.(*parse.ListNode); ok {
		for _, n := range ln.Nodes {
			res = listNodeFields(n, res)
		}
	}
	return res
}

func (action *Import) Execute(c *cli.Context) error {
	var err error

	log.WithTargets(
		log.NewColorTerminalTarget(os.Stdout, log.LevelDebug, log.LevelInfo),
		log.NewColorTerminalTarget(os.Stderr, log.LevelWarn, log.LevelFatal),
		log.NewFileWrapperTarget(CreateDefaultLogFileName(ProjectName, "main.log"), log.LevelDebug, log.LevelFatal),
	)
	defer log.Flush()

	log.Debug("Debug")
	log.Info("Info")
	log.Warn("Warn")
	log.Error("Error")
	log.Fatal("Fatal")

	importPath := filepath.Clean(c.Args().First())
	destinationPath := filepath.Clean(c.Args().Get(1))
	dbPath := filepath.Clean(destinationPath + "/grename.db")
	importSettings := []*ImportSetting{
		{
			MimeTypes:    []string{"image"},
			NameTemplate: "{{.File.MimeType}}/{{.Exif.Model}}/{{.Exif.YYYY}}/{{.Exif.MM}}/{{.Exif.DD}}/{{.Exif.YYYY}}{{.Exif.MM}}{{.Exif.DD}}_{{.Exif.Hh}}{{.Exif.Mm}}{{.Exif.Ss}}.{{.File.Extension}}",
		},
		{
			MimeTypes:    []string{"video"},
			NameTemplate: "{{.File.MimeType}}/{{.Exif.Model}}/{{.Exif.YYYY}}/{{.Exif.MM}}/{{.Exif.DD}}/{{.Exif.YYYY}}{{.Exif.MM}}{{.Exif.DD}}_{{.Exif.Hh}}{{.Exif.Mm}}{{.Exif.Ss}}.{{.File.Extension}}",
		},
	}

	importFiles := []*ImportFile{}
	db, err := database.InitDatabase(&database.Credentials{
		Driver:   "sqlite3",
		Database: dbPath,
	})
	if err != nil {
		return err
	}

	fs := afero.NewOsFs()

	err = afero.Walk(fs, importPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		mime, err := mimetype.DetectFile(path)
		if err != nil {
			return nil
		}

		fileMimeType := mime.String()
		mimeParts := strings.Split(fileMimeType, "/")
		if len(mimeParts) != 2 {
			return nil
		}

		fileMediaType := mimeParts[0]

		var foundImportSetting *ImportSetting
		for _, importSetting := range importSettings {
			if ContainsStrings(importSetting.MimeTypes, fileMediaType) || ContainsStrings(importSetting.MimeTypes, fileMimeType) {
				foundImportSetting = importSetting
				break
			}
		}

		if foundImportSetting == nil {
			return nil
		}

		fileHash, err := crypt.HashFileString(path, blake3.New(64, nil))
		if err != nil {
			return err
		}
		importFiles = append(importFiles, &ImportFile{
			Source:        path,
			ImportSetting: foundImportSetting,
			File: &database.File{
				MimeMediaType: fileMediaType,
				MimeSubType:   mimeParts[1],
				Hash:          fileHash,
				Location:      "",
				Tags:          nil,
			},
		})

		return nil
	})

	for _, importFile := range importFiles {
		f := importFile.File

		m, err := metadata.ReadFromFile(importFile.Source)
		if err != nil {
			return err
		}

		t, err := template.New("destination").Parse(importFile.ImportSetting.NameTemplate)

		templateVars := ListTemplFields(t)
		for _, templateVar := range templateVars {
			tmpTemplate, err := template.New("tmp").Parse(templateVar)
			if err != nil {
				println("failed to check template var " + templateVar)
				continue
			}
			var tmpTplOutput bytes.Buffer
			err = tmpTemplate.Execute(&tmpTplOutput, m)
			if err != nil {
				println(err.Error())
				continue
			}
			if tmpTplOutput.String() == "" {
				println("template variable empty, skipping import" + templateVar)
				continue
			}
		}

		if err != nil {
			println("!!! template parsing error - " + importFile.Source + ": " + err.Error())
			continue
		}

		var tplOutput bytes.Buffer
		err = t.Execute(&tplOutput, m)
		if err != nil {
			println("!!! template rendering error - " + importFile.Source + ": " + err.Error())
			continue
		}

		var fullDestinationPath string
		// TODO replace newlines
		location := strings.TrimPrefix(filepath.ToSlash(tplOutput.String()), "/")
		extension := filepath.Ext(location)
		location = strings.TrimSuffix(location, extension)
		locationSuffix := ""
		i := 0
		for {
			i++
			if i > 999 {
				panic("999 is the max index for existing files that is supported")
			}
			relativeDestinationPath := location + locationSuffix + extension
			fullDestinationPath = filepath.Clean(destinationPath + "/" + relativeDestinationPath)
			stat, err := fs.Stat(fullDestinationPath)
			if stat == nil || os.IsNotExist(err) {
				f.Location = relativeDestinationPath
				break
			}
			// println(stat)
			locationSuffix = fmt.Sprintf("-%d", i)
		}

		tagBasePath := filepath.ToSlash(strings.TrimPrefix(filepath.Dir(importFile.Source), importPath))
		rawTagStrings := strings.Split(tagBasePath, "/")
		tagStrings := UniqueStrings(FilterEmptyStrings(rawTagStrings))

		// TODO if hash AND imported file still exists
		db.Where("hash = ?", f.Hash).FirstOrCreate(f)

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

		if db.NewRecord(f) {
			// os.Rename(importFile.Source, fullDestinationPath)
			println("rename: " + importFile.Source + " => " + fullDestinationPath)

		} else {
			// os.Remove(importFile.Source)
			println("remove: " + importFile.Source)
		}

		//var f models.File
		//db.First(f)
		//println("renameTemplate: " + renameTemplate)
		//println("importPath: " + importPath)
		//println("destinationPath: " + f.Location)
		//println(m)
		//println(m.File.MimeType.String())
		//println(m.Exif.Make)
		//println(m.Exif.Model)
		//println(m.Exif.DateTime.Format("2006-01-02 03:04:05"))

	}

	return err

}

func ContainsStrings(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
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
