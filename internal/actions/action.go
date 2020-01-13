package actions

import "github.com/urfave/cli/v2"

type Action interface {
	Execute(c *cli.Context) error
}
