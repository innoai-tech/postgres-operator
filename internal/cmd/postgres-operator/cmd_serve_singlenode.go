package main

import (
	"github.com/innoai-tech/infra/pkg/cli"
	"github.com/innoai-tech/infra/pkg/otel"
	"github.com/innoai-tech/postgres-operator/internal/cmd/postgres-operator/sidecar"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl"
)

func init() {
	cli.AddTo(Serve, &SingleNode{})
}

// SingleNode serve single-node
type SingleNode struct {
	cli.C `component:"postgres-operator,kind=StatefulSet" envprefix:"POSTGRES_"`
	otel.Otel

	pgctl.Controller

	Server

	Daemon pgctl.Daemon

	AutoArchiver sidecar.AutoArchiver
}
