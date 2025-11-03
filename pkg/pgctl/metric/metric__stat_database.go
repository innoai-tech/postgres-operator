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
			SubSystem: "stat_database",
			Help:      "Database-level statistics from pg_stat_database",
			ValueType: prometheus.CounterValue,
			Labels:    []string{"__name__", "datname"},
			Names: []string{
				"xact_commit_total",
				"xact_rollback_total",
				"blks_read_total",
				"blks_hit_total",
				"tup_returned_total",
				"tup_fetched_total",
				"tup_inserted_total",
				"tup_updated_total",
				"tup_deleted_total",
				"deadlocks_total",
				"temp_files_total",
				"temp_bytes_total",
				"blk_read_time_total",
				"blk_write_time_total",
			},
			Gather: func(ctx context.Context, a session.Adapter, emit func(v float64, labelValues ...string)) error {
				rows, err := a.Query(ctx, sqlfrag.Pair(`
                SELECT datname,
                       xact_commit,
                       xact_rollback,
                       blks_read,
                       blks_hit,
                       tup_returned,
                       tup_fetched,
                       tup_inserted,
                       tup_updated,
                       tup_deleted,
                       deadlocks,
                       temp_files,
                       temp_bytes,
                       blk_read_time,
                       blk_write_time
                FROM pg_stat_database
                WHERE datname NOT IN ('template0', 'template1');
            `))
				if err != nil {
					return err
				}

				return session.Scan(ctx, rows, session.Recv(func(v *struct {
					Datname      string  `db:"datname"`
					XactCommit   float64 `db:"xact_commit"`
					XactRollback float64 `db:"xact_rollback"`
					BlksRead     float64 `db:"blks_read"`
					BlksHit      float64 `db:"blks_hit"`
					TupReturned  float64 `db:"tup_returned"`
					TupFetched   float64 `db:"tup_fetched"`
					TupInserted  float64 `db:"tup_inserted"`
					TupUpdated   float64 `db:"tup_updated"`
					TupDeleted   float64 `db:"tup_deleted"`
					Deadlocks    float64 `db:"deadlocks"`
					TempFiles    float64 `db:"temp_files"`
					TempBytes    float64 `db:"temp_bytes"`
					BlkReadTime  float64 `db:"blk_read_time"`
					BlkWriteTime float64 `db:"blk_write_time"`
				},
				) error {
					emit(v.XactCommit, "xact_commit_total", v.Datname)
					emit(v.XactRollback, "xact_rollback_total", v.Datname)
					emit(v.BlksRead, "blks_read_total", v.Datname)
					emit(v.BlksHit, "blks_hit_total", v.Datname)
					emit(v.TupReturned, "tup_returned_total", v.Datname)
					emit(v.TupFetched, "tup_fetched_total", v.Datname)
					emit(v.TupInserted, "tup_inserted_total", v.Datname)
					emit(v.TupUpdated, "tup_updated_total", v.Datname)
					emit(v.TupDeleted, "tup_deleted_total", v.Datname)
					emit(v.Deadlocks, "deadlocks_total", v.Datname)
					emit(v.TempFiles, "temp_files_total", v.Datname)
					emit(v.TempBytes, "temp_bytes_total", v.Datname)
					emit(v.BlkReadTime, "blk_read_time_total", v.Datname)
					emit(v.BlkWriteTime, "blk_write_time_total", v.Datname)
					return nil
				}))
			},
		},
	)
}
