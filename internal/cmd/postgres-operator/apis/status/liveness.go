package status

import (
	"context"

	"github.com/octohelm/courier/pkg/courierhttp"

	"github.com/innoai-tech/postgres-operator/pkg/pgctl"
)

// Liveness
// +gengo:injectable
type Liveness struct {
	courierhttp.MethodGet `path:"/liveness"`

	c *pgctl.Controller `inject:""`
}

func (req *Liveness) Output(ctx context.Context) (any, error) {
	err := req.c.IsReady(ctx)
	if err != nil {
		return nil, &pgctl.ErrPostgresNotReady{
			Reason: err,
		}
	}
	return map[string]any{"ready": true}, nil
}
