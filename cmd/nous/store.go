package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/leeola/nous"
	"github.com/leeola/nous/nous/flatdisk"
	"github.com/urfave/cli"
)

func StoreCmd(ctx *cli.Context) error {
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
	n, err := flatdisk.New(nConf)
	if err != nil {
		return fmt.Errorf("failed to create Flatdisk Nous: %s", err)
	}

	info := nous.Information{
		Content: strings.Join(ctx.Args(), " "),
		Tags:    ctx.StringSlice("tag"),
	}

	hash, err := n.Store(info)
	if err != nil {
		return fmt.Errorf("failed to show info: %s", err)
	}

	fmt.Println(hash)

	return nil
}
