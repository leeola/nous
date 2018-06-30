package nous

import (
	"context"
	"net/http"
)

type Server struct {
	http http.Server
}

func NewServer() (*Server, error) {
	return &Server{}, nil
}

func (s *Server) ListenAndServe() error {
	return s.http.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}
