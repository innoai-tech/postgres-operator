package internal

import (
	"context"
	"fmt"
	"maps"
	"slices"

	"github.com/innoai-tech/postgres-operator/pkg/exec"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl/pgconf"
)

func PostgresServeCommand(ctx context.Context, s pgconf.Conf) (*exec.Command, error) {
	user, err := lookupPostgresUser()
	if err != nil {
		return nil, err
	}

	pgVersion, err := s.PgDataVersion(ctx)
	if err != nil {
		return nil, err
	}

	conf := s.ToPgConf(pgVersion)

	confArgs := make([]string, 0, len(conf))
	for _, k := range slices.Sorted(maps.Keys(conf)) {
		confArgs = append(confArgs, fmt.Sprintf("%s=%s", k, conf[k]))
	}

	if err := postgresUserChown(ctx, s.DataDir.PgDataPath()); err != nil {
		return nil, err
	}

	cmd := &exec.Command{
		Name:    "postgres",
		UID:     user.UID,
		GID:     user.GID,
		WorkDir: string(s.DataDir),
		Flags: exec.Flags{
			"-D": {s.DataDir.PgDataPath()},
			"-p": {fmt.Sprintf("%d", s.Port)},
			"-c": confArgs,
		},
	}

	return cmd, nil
}
