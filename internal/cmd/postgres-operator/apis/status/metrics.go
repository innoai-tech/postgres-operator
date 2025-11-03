package status

import (
	"context"

	"github.com/octohelm/courier/pkg/courierhttp"

	"github.com/innoai-tech/postgres-operator/pkg/pgctl"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl/metric"
)

// Metrics
// +gengo:injectable
type Metrics struct {
	courierhttp.MethodGet `path:"/metrics"`

	c *pgctl.Controller `inject:""`
	e *metric.Exporter  `inject:""`
}

func (req *Metrics) Output(ctx context.Context) (any, error) {
	err := req.c.IsReady(ctx)
	if err != nil {
		return nil, &pgctl.ErrPostgresNotReady{
			Reason: err,
		}
	}
	return req.e.MetricFamilySet(ctx)
}
