package main

import (
	"context"

	infrahttp "github.com/innoai-tech/infra/pkg/http"
	"github.com/innoai-tech/postgres-operator/internal/cmd/postgres-operator/apis"
)

// +gengo:injectable
type Server struct {
	infrahttp.Server
}

func (s *Server) SetDefaults() {
	if s.Addr == "" {
		s.Addr = ":80"
	}
}

func (s *Server) beforeInit(ctx context.Context) error {
	s.ApplyRouter(apis.R)
	s.Server.SetName("postgres-operator")
	return nil
}
