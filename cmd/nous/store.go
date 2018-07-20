package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/leeola/nous"
	"github.com/urfave/cli"
)

func StoreCmd(clictx *cli.Context) error {
	n, err := newNous(clictx)
	if err != nil {
		return err // no wrap
	}

	content := strings.Join(clictx.Args(), " ")

	name := clictx.String("name")

	if name == "" && !clictx.Bool("dont-name-from-content") {
		name = content
	}

	info := nous.Data{
		Name: name,
		Type: nous.TypeText,
		Text: &nous.DataText{
			Content: content,
		},
	}

	ctx := context.Background()

	if err := n.Store(ctx, info); err != nil {
		return fmt.Errorf("store: %s", err)
	}

	fmt.Println("success")

	return nil
}
