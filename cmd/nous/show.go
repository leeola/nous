package main

import (
	"context"
	"strings"

	"github.com/urfave/cli"
)

func ShowCmd(clictx *cli.Context) error {
	n, err := newNous(clictx)
	if err != nil {
		return err // no wrap
	}

	ctx := context.Background()

	queryStr := strings.Join(clictx.Args(), " ")

	err = n.Show(ctx, queryStr)
	if err != nil {
		return err
	}

	return nil
}
