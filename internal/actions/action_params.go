package actions

import "github.com/urfave/cli/v2"

type ActionParams struct {
	Quiet bool `arg:"help:do not show any output"`
	Force bool `arg:"help:force the requested action - even if it might be not a good idea"`
	Debug bool `arg:"-d,help:debug mode with logging to Stdout and into $HOME/.graft/application.log"`
}

func parseActionParams(c *cli.Context) *ActionParams {
	return &ActionParams{
		Quiet: c.Bool("quiet"),
		Force: c.Bool("force"),
		Debug: c.Bool("debug"),
	}
}
