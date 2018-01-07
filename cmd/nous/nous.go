package main

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
		cli.StringSliceFlag{
			Name:   "config, c",
			Usage:  "load config from `PATH`",
			EnvVar: "NOUS_CONFIG",
		},
	}

	app.Commands = cli.Commands{
		{
			Name: "store",
		},
		{
			Name: "retain",
		},
		{
			Name: "research",
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
