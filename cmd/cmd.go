package cmd

import (
	"fmt"
	"net"

	"github.com/urfave/cli/v2"
)

var appName = "dnat"

func InitCmd() cli.App {
	app := cli.App{
		Name:  appName,
		Usage: "a DNAT management tool",
		Commands: []*cli.Command{
			appendCommand,
			deleteCommand,
			listCommand,
			getCommand,
			syncCommand,
			versionCommand,
		},
	}
	return app
}

var appendCommand = &cli.Command{
	Name:    "append",
	Aliases: []string{},
	Usage:   "append a rule to the end of nat chain if it does not exist",
	Action:  Append,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "in-interface",
			Aliases:  []string{"i"},
			Required: true,
			Action: func(ctx *cli.Context, s string) error {
				return nil
			},
		},
		&cli.IntFlag{
			Name:     "dport",
			Aliases:  []string{"p"},
			Required: true,
			Action: func(ctx *cli.Context, port int) error {
				if port > 65536 {
					return fmt.Errorf("port value %d out of range[0-65535]", port)
				}
				return nil
			},
		},
		&cli.StringFlag{
			Name:     "to-destination",
			Aliases:  []string{"d"},
			Required: true,
			Action: func(ctx *cli.Context, dest string) error {
				_, err := net.ResolveTCPAddr("tcp", dest)
				return err
			},
		},
	},
}

var deleteCommand = &cli.Command{
	Name:    "delete",
	Aliases: []string{},
	Usage:   "delete a rule by id",
	Action:  Delete,
	Flags: []cli.Flag{
		&cli.Int64Flag{
			Name:     "id",
			Usage:    "rule id",
			Required: true,
		},
	},
}

var listCommand = &cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
	Usage:   "list all rules",
	Action:  List,
}

var getCommand = &cli.Command{
	Name:    "get",
	Aliases: []string{},
	Usage:   "get one rule by id",
	Action:  Get,
	Flags: []cli.Flag{
		&cli.Int64Flag{
			Name:     "id",
			Usage:    "rule id",
			Required: true,
		},
	},
}

var syncCommand = &cli.Command{
	Name:    "sync",
	Aliases: []string{},
	Usage:   "sync rules to local machine",
	Action:  Sync,
}
