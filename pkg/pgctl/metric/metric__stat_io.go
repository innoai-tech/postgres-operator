package metric

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/octohelm/storage/pkg/session"
	"github.com/octohelm/storage/pkg/sqlfrag"
)

func init() {
	register(
		&Metric{
			Namespace: namespace,
			SubSystem: "stat_io",
			Help:      "I/O statistics by backend type and object from pg_stat_io (PostgreSQL 15+)",
			ValueType: prometheus.CounterValue,
			Labels:    []string{"__name__", "backend_type", "context", "object"},
			Names: []string{
				"extends_total",
				"fsync_time_total",
				"fsyncs_total",
				"read_time_total",
				"reads_total",
				"write_time_total",
				"writes_total",
			},
			Gather: func(ctx context.Context, a session.Adapter, emit func(v float64, labelValues ...string)) error {
				rows, err := a.Query(ctx, sqlfrag.Pair(`
                SELECT backend_type,
                       context,
                       object,
                       extends,
                       fsync_time,
                       fsyncs,
                       read_time,
                       reads,
                       write_time,
                       writes
                FROM pg_stat_io;
            `))
				if err != nil {
					return err
				}

				return session.Scan(ctx, rows, session.Recv(func(v *struct {
					BackendType string  `db:"backend_type"`
					Context     string  `db:"context"`
					Object      string  `db:"object"`
					Extends     float64 `db:"extends"`
					FsyncTime   float64 `db:"fsync_time"`
					Fsyncs      float64 `db:"fsyncs"`
					ReadTime    float64 `db:"read_time"`
					Reads       float64 `db:"reads"`
					WriteTime   float64 `db:"write_time"`
					Writes      float64 `db:"writes"`
				},
				) error {
					emit(v.Extends, "extends_total", v.BackendType, v.Context, v.Object)
					emit(v.FsyncTime, "fsync_time_total", v.BackendType, v.Context, v.Object)
					emit(v.Fsyncs, "fsyncs_total", v.BackendType, v.Context, v.Object)
					emit(v.ReadTime, "read_time_total", v.BackendType, v.Context, v.Object)
					emit(v.Reads, "reads_total", v.BackendType, v.Context, v.Object)
					emit(v.WriteTime, "write_time_total", v.BackendType, v.Context, v.Object)
					emit(v.Writes, "writes_total", v.BackendType, v.Context, v.Object)
					return nil
				}))
			},
		},
	)
}
