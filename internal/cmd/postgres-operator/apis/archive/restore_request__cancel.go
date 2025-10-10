package archive

import (
	"context"

	"github.com/innoai-tech/postgres-operator/pkg/pgctl"
	"github.com/octohelm/courier/pkg/courierhttp"
)

// CancelRestoreRequest
// +gengo:injectable
type CancelRestoreRequest struct {
	courierhttp.MethodDelete `path:"/request-restore"`

	c *pgctl.Controller `inject:""`
}

func (req *CancelRestoreRequest) Output(ctx context.Context) (any, error) {
	return nil, req.c.ArchiveController().CancelRestore(ctx)
}
