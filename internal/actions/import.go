package actions

import (
	"github.com/urfave/cli/v2"
)

type Import struct {
}

func (action *Import) Execute(c *cli.Context) error {
	println("import executed")

	return nil
}
