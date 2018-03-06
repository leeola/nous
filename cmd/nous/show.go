package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/leeola/nous/nous/flatdisk"
	"github.com/urfave/cli"
)

func ShowCmd(ctx *cli.Context) error {
	path := ctx.GlobalString("path")
	if path == "" {
		return errors.New("missing required path flag value")
	}

	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("failed to make flatdisk path %s: %s", path, err)
	}

	nConf := flatdisk.Config{
		Path: path,
	}
	nous, err := flatdisk.New(nConf)
	if err != nil {
		return fmt.Errorf("failed to create Flatdisk Nous: %s", err)
	}

	infos, err := nous.Retrieve(ctx.Args()...)
	if err != nil {
		return fmt.Errorf("failed to show info: %s", err)
	}

	for _, info := range infos {
		fmt.Printf("match: %s\n\n", info.Content)
	}

	return nil
}
