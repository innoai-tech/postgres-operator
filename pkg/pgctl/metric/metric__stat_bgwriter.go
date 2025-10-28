package metric

import (
	"context"

	"github.com/octohelm/storage/pkg/session"
	"github.com/octohelm/storage/pkg/sqlfrag"
	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	register(
		&Metric{
			Namespace: namespace,
			SubSystem: "stat_bgwriter",
			Help:      "Background writer statistics from pg_stat_bgwriter",
			ValueType: prometheus.CounterValue,
			Labels:    []string{"__name__"},
			Names: []string{
				"buffers_alloc_total",
				"buffers_backend_total",
				"buffers_checkpoint_total",
				"buffers_clean_total",
				"checkpoint_sync_time_total",
				"checkpoint_write_time_total",
				"checkpoints_req_total",
				"checkpoints_timed_total",
				"maxwritten_clean_total",
			},
			Gather: func(ctx context.Context, a session.Adapter, emit func(v float64, labelValues ...string)) error {
				rows, err := a.Query(ctx, sqlfrag.Pair(`
                SELECT buffers_alloc,
                       buffers_backend,
                       buffers_checkpoint,
                       buffers_clean,
                       checkpoint_sync_time,
                       checkpoint_write_time,
                       checkpoints_req,
                       checkpoints_timed,
                       maxwritten_clean
                FROM pg_stat_bgwriter;
            `))
				if err != nil {
					return err
				}

				return session.Scan(ctx, rows, session.Recv(func(v *struct {
					BuffersAlloc      float64 `db:"buffers_alloc"`
					BuffersBackend    float64 `db:"buffers_backend"`
					BuffersCheckpoint float64 `db:"buffers_checkpoint"`
					BuffersClean      float64 `db:"buffers_clean"`
					CheckpointSync    float64 `db:"checkpoint_sync_time"`
					CheckpointWrite   float64 `db:"checkpoint_write_time"`
					CheckpointsReq    float64 `db:"checkpoints_req"`
					CheckpointsTimed  float64 `db:"checkpoints_timed"`
					MaxwrittenClean   float64 `db:"maxwritten_clean"`
				},
				) error {
					emit(v.BuffersAlloc, "buffers_alloc_total")
					emit(v.BuffersBackend, "buffers_backend_total")
					emit(v.BuffersCheckpoint, "buffers_checkpoint_total")
					emit(v.BuffersClean, "buffers_clean_total")
					emit(v.CheckpointSync, "checkpoint_sync_time_total")
					emit(v.CheckpointWrite, "checkpoint_write_time_total")
					emit(v.CheckpointsReq, "checkpoints_req_total")
					emit(v.CheckpointsTimed, "checkpoints_timed_total")
					emit(v.MaxwrittenClean, "maxwritten_clean_total")
					return nil
				}))
			},
		},
	)
}
