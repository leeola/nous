package main // import "github.com/leeola/nous/cmd/nous"

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/leeola/fixity"
	"github.com/leeola/fixity/config"
	_ "github.com/leeola/fixity/defaultpkg"
	"github.com/leeola/nous"
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
			Value:  "./_store/fixity.json",
			EnvVar: "NOUS_PATH",
		},
	}

	app.Commands = cli.Commands{
		{
			Name:      "store",
			Action:    StoreCmd,
			ArgsUsage: "CONTENT",
			Usage:     "store CONTENT in nous",
			Flags: []cli.Flag{
				cli.StringSliceFlag{
					Name:  "tag, t",
					Usage: "apply the given tags to the information",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "define the name for this content",
				},
				cli.StringFlag{
					Name:  "value, v",
					Usage: "an optional value for this content",
				},
				cli.BoolFlag{
					Name:  "dont-name-from-content",
					Usage: "dont infer name from content if unspecified",
				},
			},
		},
		{
			Name:      "show",
			Action:    ShowCmd,
			Aliases:   []string{"s"},
			ArgsUsage: "QUERY",
			Usage:     "show QUERY",
			Flags: []cli.Flag{
				cli.StringSliceFlag{
					Name:  "tag, t",
					Usage: "query with tags",
				},
			},
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

func newNous(clictx *cli.Context) (*nous.Nous, error) {
	path := clictx.GlobalString("path")
	if path == "" {
		return nil, errors.New("missing required path flag value")
	}

	// this feels a bit weird here, but cmds should only be called once,
	// so i guess it's okay..
	config.Configure(func(c config.Config) (config.Config, error) {
		c.RootPath = filepath.Dir(path)

		return c, nil
	})

	fixi, err := fixity.NewFromPath("", path)
	if err != nil {
		return nil, fmt.Errorf("new fixity: %v", err)
	}

	return nous.New(fixi)
}
