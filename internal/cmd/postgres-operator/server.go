package main

import (
	"context"
	"net/http"

	infrahttp "github.com/innoai-tech/infra/pkg/http"
	"github.com/innoai-tech/postgres-operator/internal/cmd/postgres-operator/apis"
	"github.com/innoai-tech/postgres-operator/internal/webapp"
	"github.com/octohelm/jwx/pkg/encryption"
)

// +gengo:injectable
type Server struct {
	WebUI webapp.WebUI

	infrahttp.Server
}

func (s *Server) SetDefaults() {
	if s.Addr == "" {
		s.Addr = ":80"
	}

	s.WebUI.Use("console")
	s.WebUI.SetDefaults()
}

func (s *Server) beforeInit(ctx context.Context) error {
	enc := &encryption.MiddlewareProvider{}
	if err := enc.Init(ctx); err != nil {
		return err
	}

	s.Server.SetName("postgres-operator")
	s.Server.ApplyRouterHandlers(enc.Wrap)

	apiHandler, err := s.NewHandler(ctx, apis.R)
	if err != nil {
		return err
	}

	s.ApplyGlobalHandlers(func(http.Handler) http.Handler {
		mux := http.NewServeMux()

		mux.HandleFunc("/api/", apiHandler.ServeHTTP)
		mux.HandleFunc("/", s.WebUI.ServeHTTP)

		return mux
	})

	return nil
}
