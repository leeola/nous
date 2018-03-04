package main // import "github.com/leeola/nous/cmd/nous"

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "nous"
	app.Usage = ""
	app.Flags = []cli.Flag{
		// Nous config disabled currently.
		// cli.StringFlag{
		// 	Name:   "config, c",
		// 	Usage:  "load config from `PATH`",
		// 	EnvVar: "NOUS_CONFIG",
		// },

		cli.StringFlag{
			Name:   "path, p",
			Usage:  "load Nous flatdisk from `PATH`",
			EnvVar: "NOUS_PATH",
		},
	}

	app.Commands = cli.Commands{
		{
			Name: "store",
		},
		{
			Name:   "show",
			Action: ShowCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
