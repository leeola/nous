package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/leeola/nous"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli"
)

func ServeCmd(ctx *cli.Context) error {
	srv, err := nous.NewServer()
	if err != nil {
		return fmt.Errorf("newserver: %v", err)
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		log.Info().Msg("sigint triggered")

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Info().Err(err).Msg("server shutdown failed")
		}

		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Info().Err(err).Msg("http server failed")
	}

	<-idleConnsClosed

	return nil
}

func nousFromContext(clictx *cli.Context) (*nous.Nous, error) {
	//path := ctx.GlobalString("path")
	//if path == "" {
	//	return errors.New("missing required path flag value")
	//}

	return nil, errors.New("not implemented")
}
