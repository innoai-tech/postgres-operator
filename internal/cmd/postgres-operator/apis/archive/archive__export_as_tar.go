package archive

import (
	"context"
	"fmt"
	"io"
	"mime"
	"net/http"

	archivev1 "github.com/innoai-tech/postgres-operator/pkg/apis/archive/v1"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl"
	"github.com/octohelm/courier/pkg/courierhttp"
	"github.com/octohelm/x/logr"
)

// ExportArchiveAsTar
// +gengo:injectable
type ExportArchiveAsTar struct {
	courierhttp.MethodGet `path:"/archives/{archiveCode}/as-tar"`

	ArchiveCode archivev1.ArchiveCode `name:"archiveCode" in:"path"`

	c *pgctl.Controller `inject:""`
}

func (req *ExportArchiveAsTar) Output(ctx context.Context) (any, error) {
	writeTo, err := req.c.ArchiveController().ExportArchiveAsTar(ctx, req.ArchiveCode)
	if err != nil {
		return nil, err
	}

	return any(&upgrader{
		ArchiveCode: req.ArchiveCode,
		WriteTo:     writeTo,
	}), nil
}

type upgrader struct {
	ArchiveCode archivev1.ArchiveCode
	WriteTo     func(w io.Writer) error
}

func (u *upgrader) Upgrade(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/x-tar")
	w.Header().Set("Content-Disposition", mime.FormatMediaType("attachment", map[string]string{
		"filename": fmt.Sprintf("%s.tar", u.ArchiveCode),
	}))

	w.WriteHeader(http.StatusOK)

	if err := u.WriteTo(w); err != nil {
		logr.FromContext(r.Context()).Error(err)
	}

	return nil
}
