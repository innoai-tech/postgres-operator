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
			SubSystem: "stat_user_tables",
			Help:      "User table statistics from pg_stat_user_tables",
			ValueType: prometheus.CounterValue,
			Labels:    []string{"__name__", "schemaname", "relname"},
			Names: []string{
				"seq_scan_total",
				"idx_scan_total",
				"n_tup_ins_total",
				"n_tup_upd_total",
				"n_tup_del_total",
			},
			Gather: func(ctx context.Context, a session.Adapter, emit func(v float64, labelValues ...string)) error {
				rows, err := a.Query(ctx, sqlfrag.Pair(`
                SELECT schemaname,
                       relname,
                       seq_scan,
                       idx_scan,
                       n_tup_ins,
                       n_tup_upd,
                       n_tup_del
                FROM pg_stat_user_tables;
            `))
				if err != nil {
					return err
				}

				return session.Scan(ctx, rows, session.Recv(func(v *struct {
					Schemaname string  `db:"schemaname"`
					Relname    string  `db:"relname"`
					SeqScan    float64 `db:"seq_scan"`
					IdxScan    float64 `db:"idx_scan"`
					NTupIns    float64 `db:"n_tup_ins"`
					NTupUpd    float64 `db:"n_tup_upd"`
					NTupDel    float64 `db:"n_tup_del"`
				},
				) error {
					emit(v.SeqScan, "seq_scan_total", v.Schemaname, v.Relname)
					emit(v.IdxScan, "idx_scan_total", v.Schemaname, v.Relname)
					emit(v.NTupIns, "n_tup_ins_total", v.Schemaname, v.Relname)
					emit(v.NTupUpd, "n_tup_upd_total", v.Schemaname, v.Relname)
					emit(v.NTupDel, "n_tup_del_total", v.Schemaname, v.Relname)
					return nil
				}))
			},
		},
	)
}
