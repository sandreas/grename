package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"grename/internal/actions"
	"os"
)

func main() {
	globalFlags := []cli.Flag{
		&cli.BoolFlag{Name: "quiet, q", Usage: "do not show any output"},                                           // does quiet make sense in find?
		&cli.BoolFlag{Name: "force, f", Usage: "force the requested action - even if it might be not a good idea"}, // does force make sense in find?
		&cli.BoolFlag{Name: "debug", Usage: "debug mode with logging to Stdout and into $HOME/.graft/application.log"},
	}

	importFlags := []cli.Flag{
		// &cli.BoolFlag{Name: "keep-duplicates", Usage: "keep duplicate files"},
		// &cli.StringFlag{Name: "tpl", Usage: "filename template"},
		// &cli.StringFlag{Name: "include-media-types", Value: "image,video", Usage: "media types to include"},
	}

	app := cli.NewApp()
	app.Name = "grename"
	app.Version = "0.1"
	app.Usage = "rename files based on metadata"

	app.Commands = []*cli.Command{
		{
			Name:    "import",
			Aliases: []string{"i"},
			Action:  new(actions.Import).Execute,
			Usage:   "import files",
			Flags:   mergeFlags(globalFlags, importFlags),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		println(fmt.Errorf("error: %s", err))
	}
}

func mergeFlags(flagsToMerge ...[]cli.Flag) []cli.Flag {
	var mergedFlags []cli.Flag
	for _, flags := range flagsToMerge {
		mergedFlags = append(mergedFlags, flags...)
	}
	return mergedFlags
}

/*
func (action *AbstractAction) initLogging() {
	if !action.CliParameters.Debug {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
		return
	}
	log.SetOutput(os.Stdout)


	logFileName := homeDir + "/graft.log"
	logFile, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("could not open logfile: ", logFile, err)
		return
	}
	defer logFile.Close()
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
}
*/
