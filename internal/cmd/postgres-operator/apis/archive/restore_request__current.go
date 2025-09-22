package archive

import (
	"context"

	archivev1 "github.com/innoai-tech/postgres-operator/pkg/apis/archive/v1"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl"
	"github.com/octohelm/courier/pkg/courierhttp"
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

	return &archivev1.Archive{Code: code}, nil
}
