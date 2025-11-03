package service

import (
	"context"

	"github.com/octohelm/courier/pkg/courierhttp"

	"github.com/innoai-tech/postgres-operator/pkg/pgctl"
)

// Restart
// +gengo:injectable
type Restart struct {
	courierhttp.MethodPost `path:"/restart"`

	c *pgctl.Controller `inject:""`
}

func (req *Restart) Output(ctx context.Context) (any, error) {
	return nil, req.c.Restart(ctx)
}
