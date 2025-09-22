package archive

import (
	"context"

	archivev1 "github.com/innoai-tech/postgres-operator/pkg/apis/archive/v1"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl"
	"github.com/octohelm/courier/pkg/courierhttp"
)

// RequestRestoreArchive
// +gengo:injectable
type RequestRestoreArchive struct {
	courierhttp.MethodPut `path:"/archives/{archiveCode}/restore-request"`

	ArchiveCode archivev1.ArchiveCode `name:"archiveCode" in:"path"`

	c *pgctl.Controller `inject:""`
}

func (req *RequestRestoreArchive) Output(ctx context.Context) (any, error) {
	return nil, req.c.ArchiveController().RequestRestore(ctx, req.ArchiveCode)
}
