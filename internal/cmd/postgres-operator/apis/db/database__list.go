package db

import (
	"context"

	"github.com/octohelm/courier/pkg/courierhttp"

	"github.com/innoai-tech/postgres-operator/pkg/pgctl"
)

// +gengo:injectable
type ListDatabase struct {
	courierhttp.MethodGet `path:"/databases"`

	c *pgctl.Controller `inject:""`
}

func (req *ListDatabase) Output(ctx context.Context) (any, error) {
	return req.c.ListDatabase(ctx)
}
