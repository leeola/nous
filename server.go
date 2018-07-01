package nous

import (
	"context"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog/log"
)

type Server struct {
	bindAddr string
	http     http.Server
}

func NewServer() (*Server, error) {
	srv := &Server{
		bindAddr: ":8081",
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", srv.handleGet)
	r.Post("/", srv.handlePost)

	srv.http = http.Server{
		Addr:    srv.bindAddr,
		Handler: r,
	}

	return srv, nil
}

func (s *Server) ListenAndServe() error {
	log.Info().Str("addr", s.bindAddr).Msg("listening..")
	return s.http.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Info().Msg("shutting down..")
	return s.http.Shutdown(ctx)
}

func (s *Server) handleGet(w http.ResponseWriter, r *http.Request) {
	t := template.New("fieldname example")
	t, _ = t.Parse(`

<html>
  <body>
    <form method="post" enctype="multipart/form-data">
      <input type="file" name="file"/>
      <input type="submit" class="submit" value="Submit" />
    </form>
  </body>
</html>
`)
	p := map[string]string{"UserName": "Astaxie"}
	t.Execute(w, p)
}

func (s *Server) handlePost(w http.ResponseWriter, r *http.Request) {
	mr, err := r.MultipartReader()
	if err != nil {
		log.Error().Err(err).Msg("multipart failed")
		return
	}

	// just parsing one part for now.
	part, err := mr.NextPart()
	if err != nil {
		log.Error().Err(err).Msg("nextpart failed")
		return
	}

	log.Info().Str("filename", part.FileName()).Str("formname", part.FormName()).Msg("wee")

	f, err := os.OpenFile("foo.image", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Error().Err(err).Msg("open failed")
		return
	}
	defer f.Close()

	if io.Copy(f, part); err != nil {
		log.Error().Err(err).Msg("copy failed")
		return
	}

	if err := f.Sync(); err != nil {
		log.Error().Err(err).Msg("sync failed")
		return
	}
}
