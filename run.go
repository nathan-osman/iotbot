package main

import (
	"github.com/urfave/cli/v2"
)

var runCommand = &cli.Command{
	Name:  "run",
	Usage: "run the specified script(s)",
	Action: func(c *cli.Context) error {
		return nil
	},
}
