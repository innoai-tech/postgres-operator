package status

import (
	"context"

	"github.com/innoai-tech/postgres-operator/pkg/pgctl"
	"github.com/octohelm/courier/pkg/courierhttp"
)

// Readiness
// +gengo:injectable
type Readiness struct {
	courierhttp.MethodGet `path:"/readiness"`

	c *pgctl.Controller `inject:""`
}

func (req *Readiness) Output(ctx context.Context) (any, error) {
	err := req.c.IsReady(ctx)
	if err != nil {
		return nil, &pgctl.ErrPostgresNotReady{
			Reason: err,
		}
	}
	return map[string]any{"ready": true}, nil
}
