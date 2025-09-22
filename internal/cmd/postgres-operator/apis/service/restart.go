package service

import (
	"context"

	"github.com/innoai-tech/postgres-operator/pkg/pgctl"
	"github.com/octohelm/courier/pkg/courierhttp"
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
