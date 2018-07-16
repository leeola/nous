package main

import (
	"errors"
	"fmt"
	"os"

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

	return nil
}
