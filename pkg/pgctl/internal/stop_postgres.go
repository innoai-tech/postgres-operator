package internal

import (
	"context"

	"github.com/innoai-tech/postgres-operator/pkg/exec"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl/pgconf"
)

func StopPostgres(ctx context.Context, s pgconf.Conf) error {
	user, err := lookupPostgresUser()
	if err != nil {
		return err
	}

	cmd := &exec.Command{
		Name:    "pg_ctl",
		UID:     user.UID,
		GID:     user.GID,
		WorkDir: string(s.DataDir),
		Args: []string{
			"-D", s.DataDir.PgDataPath(),
			"stop",
		},
		Flags: exec.Flags{
			"-m": {"fast"},
		},
	}

	return cmd.Run(ctx)
}
