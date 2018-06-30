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
			Name:   "store",
			Action: StoreCmd,
			Flags: []cli.Flag{
				cli.StringSliceFlag{
					Name:  "tag, t",
					Usage: "apply the given tags to the information",
				},
			},
		},
		{
			Name:   "show",
			Action: ShowCmd,
		},
		{
			Name:   "serve",
			Action: ServeCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
