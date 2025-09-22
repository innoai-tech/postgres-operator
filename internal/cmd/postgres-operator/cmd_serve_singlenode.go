package main

import (
	"github.com/innoai-tech/infra/pkg/cli"
	"github.com/innoai-tech/infra/pkg/otel"
	"github.com/innoai-tech/postgres-operator/internal/cmd/postgres-operator/domain/auth"
	"github.com/innoai-tech/postgres-operator/internal/cmd/postgres-operator/sidecar"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl"
	"github.com/octohelm/jwx/pkg/encryption"
	"github.com/octohelm/jwx/pkg/jwk"
	"github.com/octohelm/jwx/pkg/sign"
	"github.com/octohelm/objectkind/pkg/idgen"
)

func init() {
	cli.AddTo(Serve, &SingleNode{})
}

// SingleNode serve single-node
type SingleNode struct {
	cli.C `component:"postgres-operator,kind=StatefulSet" envprefix:"POSTGRES_"`
	otel.Otel
	idgen.IDGen

	KeySet jwk.KeySet
	Sign   sign.JWTSigner
	Enc    encryption.Encrypter

	pgctl.Controller

	AuthService auth.Service

	Server

	Daemon pgctl.Daemon

	AutoArchiver sidecar.AutoArchiver
}
