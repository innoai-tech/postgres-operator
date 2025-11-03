package pgctl

import (
	"context"

	"github.com/octohelm/x/logr"
	"github.com/octohelm/x/sync/singleflight"

	"github.com/innoai-tech/postgres-operator/pkg/exec"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl/archive"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl/internal"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl/pgconf"
)

type Postgres struct {
	pgconf.Conf

	sfg singleflight.Group[string]
	cmd *exec.Command
}

func (p *Postgres) Serve(ctx context.Context) error {
	if err := (&archive.Controller{DataDir: p.DataDir}).CommitRestore(ctx); err != nil {
		return err
	}
	cmd, err := internal.PostgresServeCommand(ctx, p.Conf)
	if err != nil {
		return err
	}
	p.cmd = cmd
	return cmd.Run(ctx)
}

func (s *Postgres) Shutdown(ctx context.Context) error {
	if cmd := s.cmd; cmd != nil {
		err, _ := s.sfg.Do("shutdown", func() error {
			logr.FromContext(ctx).Info("shutting down pg")
			// same as
			// pg_ctl stop -m fast
			return internal.StopPostgres(ctx, s.Conf)
		})
		return err
	}
	return nil
}
