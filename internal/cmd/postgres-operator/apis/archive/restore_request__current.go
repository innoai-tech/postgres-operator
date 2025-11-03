package archive

import (
	"context"

	"github.com/octohelm/courier/pkg/courierhttp"
	"github.com/octohelm/objectkind/pkg/runtime"

	archivev1 "github.com/innoai-tech/postgres-operator/pkg/apis/archive/v1"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl"
)

// CurrentRestoreRequest
// +gengo:injectable
type CurrentRestoreRequest struct {
	courierhttp.MethodGet `path:"/request-restore"`

	c *pgctl.Controller `inject:""`
}

func (req *CurrentRestoreRequest) Output(ctx context.Context) (any, error) {
	code, err := req.c.ArchiveController().CurrentRestoreRequest(ctx)
	if err != nil {
		return nil, err
	}

	if code == "" {
		return nil, &archivev1.ErrArchiveNotFound{}
	}

	return runtime.Build(func(a *archivev1.Archive) {
		a.Code = code
	}), nil
}
