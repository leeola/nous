package main

import (
	"context"
	"fmt"

	"github.com/leeola/nous"
	"github.com/urfave/cli"
)

func StoreCmd(clictx *cli.Context) error {
	n, err := newNous(clictx)
	if err != nil {
		return err // no wrap
	}

	info := nous.Data{
		Name: "foo",
		Type: nous.TypeText,
		Text: &nous.DataText{
			Content: "foo",
		},
	}

	ctx := context.Background()

	if err := n.Store(ctx, info); err != nil {
		return fmt.Errorf("store: %s", err)
	}

	fmt.Println("success")

	return nil
}
