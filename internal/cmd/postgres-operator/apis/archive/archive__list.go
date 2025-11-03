package archive

import (
	"context"

	"github.com/octohelm/courier/pkg/courierhttp"

	"github.com/innoai-tech/postgres-operator/pkg/pgctl"
)

// ListArchive
// +gengo:injectable
type ListArchive struct {
	courierhttp.MethodGet `path:"/archives"`

	c *pgctl.Controller `inject:""`
}

func (req *ListArchive) Output(ctx context.Context) (any, error) {
	return req.c.ArchiveController().ListArchive(ctx)
}
