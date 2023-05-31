package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var (
	Version string
	Time    string
)

var versionCommand = &cli.Command{
	Name:    "version",
	Aliases: []string{"v"},
	Usage:   "print version",
	Action:  PrintVersion,
}

func PrintVersion(cCtx *cli.Context) error {
	fmt.Println("Version:", Version)
	return nil
}
