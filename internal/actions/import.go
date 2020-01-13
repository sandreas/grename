package actions

import (
	"github.com/urfave/cli/v2"
	"grename/internal/database"
	"grename/internal/metadata"
)

type Import struct {
}

func (action *Import) Execute(c *cli.Context) error {
	source := c.Args().First()
	destination := c.Args().Get(1)
	renameTemplate := c.String("tpl")

	m, err := metadata.ReadFromFile("var/samples/Nikon_D70.jpg")
	if err != nil {
		return err
	}

	db, err := database.InitDatabase(&database.Credentials{
		Driver:   "sqlite3",
		Database: "./var/app.db",
	})
	println(db)
	if err != nil {
		println("DB error")
		println(err.Error())
		return err
	}
	//var f models.File
	//db.First(f)
	println("renameTemplate: " + renameTemplate)
	println("source: " + source)
	println("destination: " + destination)
	println(m.File.Extension)
	println(m.File.MimeType.String())
	println(m.Exif.Make)
	println(m.Exif.Model)
	println(m.Exif.DateTime.Format("2006-01-02 03:04:05"))
	return nil
}
