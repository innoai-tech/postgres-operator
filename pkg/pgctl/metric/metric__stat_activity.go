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
			SubSystem: "stat_activity",
			Help:      "Number of connections by state in pg_stat_activity",
			ValueType: prometheus.GaugeValue,
			Labels:    []string{"datname", "state"},
			Gather: func(ctx context.Context, a session.Adapter, emit func(v float64, labelValues ...string)) error {
				rows, err := a.Query(ctx, sqlfrag.Pair(`
		SELECT datname, state, COUNT(*) AS count
		FROM pg_stat_activity
		WHERE datname IS NOT NULL
		GROUP BY datname, state;
	`))
				if err != nil {
					return err
				}

				return session.Scan(ctx, rows, session.Recv(func(v *struct {
					Database string `db:"datname"`
					State    string `db:"state"`
					Count    int    `db:"count"`
				},
				) error {
					emit(float64(v.Count), v.Database, v.State)
					return nil
				}))
			},
		},
	)
}
