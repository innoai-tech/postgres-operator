package archive

import (
	"context"
	"io"

	archivev1 "github.com/innoai-tech/postgres-operator/pkg/apis/archive/v1"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl"
	"github.com/octohelm/courier/pkg/courierhttp"
)

// ImportArchiveFromTar
// +gengo:injectable
type ImportArchiveFromTar struct {
	courierhttp.MethodPut `path:"/archives/{archiveCode}/from-tar"`

	ArchiveCode archivev1.ArchiveCode `name:"archiveCode" in:"path"`

	Body io.ReadCloser `in:"body"`

	c *pgctl.Controller `inject:""`
}

func (req *ImportArchiveFromTar) Output(ctx context.Context) (any, error) {
	err := req.c.ArchiveController().ImportArchiveFromTar(ctx, req.ArchiveCode, req.Body)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
