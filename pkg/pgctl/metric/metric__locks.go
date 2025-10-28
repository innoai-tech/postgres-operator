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
			SubSystem: "locks",
			Help:      "Number of locks granted and waiting, grouped by mode",
			ValueType: prometheus.GaugeValue,
			Labels:    []string{"mode", "granted"},
			Gather: func(ctx context.Context, a session.Adapter, emit func(v float64, labelValues ...string)) error {
				rows, err := a.Query(ctx, sqlfrag.Pair(`
				SELECT mode,
				       granted,
				       COUNT(*) AS count
				FROM pg_locks
				GROUP BY mode, granted;
			`))
				if err != nil {
					return err
				}

				return session.Scan(ctx, rows, session.Recv(func(v *struct {
					Mode    string `db:"mode"`
					Granted bool   `db:"granted"`
					Count   int64  `db:"count"`
				},
				) error {
					granted := "false"
					if v.Granted {
						granted = "true"
					}
					emit(float64(v.Count), v.Mode, granted)
					return nil
				}))
			},
		},
	)
}
