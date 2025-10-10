package archive

import (
	"context"

	archivev1 "github.com/innoai-tech/postgres-operator/pkg/apis/archive/v1"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl"
	"github.com/octohelm/courier/pkg/courierhttp"
)

// DeleteArchive
// +gengo:injectable
type DeleteArchive struct {
	courierhttp.MethodDelete `path:"/archives/{archiveCode}"`

	ArchiveCode archivev1.ArchiveCode `name:"archiveCode" in:"path"`

	c *pgctl.Controller `inject:""`
}

func (req *DeleteArchive) Output(ctx context.Context) (any, error) {
	return nil, req.c.ArchiveController().DeleteArchive(ctx, req.ArchiveCode)
}
