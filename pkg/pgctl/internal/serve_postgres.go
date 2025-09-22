package internal

import (
	"context"
	"fmt"
	"maps"
	"slices"

	"github.com/innoai-tech/postgres-operator/pkg/exec"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl/pgconf"
	"github.com/innoai-tech/postgres-operator/pkg/units"
)

func PostgresServeCommand(ctx context.Context, s pgconf.Conf) (*exec.Command, error) {
	user, err := lookupPostgresUser()
	if err != nil {
		return nil, err
	}

	conf := map[string]string{
		// https://sourcegraph.com/docs/admin/config/postgres-conf#resource-dependent-configuration
		"max_connections":                  fmt.Sprintf("%d", s.MaxConnections),
		"max_parallel_maintenance_workers": fmt.Sprintf("%d", s.CPU),
		"max_parallel_workers":             fmt.Sprintf("%d", s.CPU),
		"max_worker_processes":             fmt.Sprintf("%d", s.CPU),
		"max_parallel_workers_per_gather":  fmt.Sprintf("%d", s.CPU/2),

		"effective_cache_size": fmt.Sprintf("%dMB", int(s.MEM/units.MiB)*3/4),
		"shared_buffers":       fmt.Sprintf("%dMB", int(s.MEM/units.MiB)/4),
		"maintenance_work_mem": fmt.Sprintf("%dMB", int(s.MEM/units.MiB)/16),
		"work_mem":             fmt.Sprintf("%dMB", int(s.MEM/units.MiB)/(4*s.MaxConnections*(s.CPU/2))),
	}

	confArgs := make([]string, 0, len(conf))
	for _, k := range slices.Sorted(maps.Keys(conf)) {
		confArgs = append(confArgs, fmt.Sprintf("%s=%s", k, conf[k]))
	}

	chown := &exec.Command{
		Name:    "chown",
		WorkDir: string(s.DataDir),
		Args: []string{
			"-R", "postgres:postgres", s.DataDir.PgDataPath(),
		},
	}

	if err := chown.Run(ctx); err != nil {
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
