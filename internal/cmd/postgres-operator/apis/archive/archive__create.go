package archive

import (
	"context"

	"github.com/innoai-tech/postgres-operator/pkg/pgctl"
	"github.com/octohelm/courier/pkg/courierhttp"
)

// ImportArchive
// +gengo:injectable
type CreateArchive struct {
	courierhttp.MethodPost `path:"/archives"`

	c *pgctl.Controller `inject:""`
}

func (req *CreateArchive) Output(ctx context.Context) (any, error) {
	return req.c.CreateArchive(ctx)
}
