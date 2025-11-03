package archive

import (
	"archive/tar"
	"context"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/octohelm/objectkind/pkg/runtime"
	sqltypetime "github.com/octohelm/storage/pkg/sqltype/time"

	archivev1 "github.com/innoai-tech/postgres-operator/pkg/apis/archive/v1"
)

func (c *Controller) ExportArchiveAsTar(ctx context.Context, code archivev1.ArchiveCode) (func(writer io.Writer) error, error) {
	root, err := os.OpenRoot(c.DataDir.PgArchivePath(string(code)))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, &archivev1.ErrArchiveNotFound{}
		}
		return nil, err
	}

	return func(w io.Writer) error {
		tw := tar.NewWriter(w)
		defer tw.Close()

		return tw.AddFS(root.FS())
	}, nil
}

func (c *Controller) ListArchive(ctx context.Context) (*archivev1.ArchiveList, error) {
	list := &archivev1.ArchiveList{}

	root := c.DataDir.PgArchivePath()

	if _, err := os.Stat(root); err != nil {
		if os.IsNotExist(err) {
			return list, nil
		}
		return nil, err
	}

	if err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if path == root {
			return nil
		}

		if d.IsDir() {
			a := runtime.Build(func(a *archivev1.Archive) {
				a.Code = archivev1.ArchiveCode(d.Name())
				a.SetCreationTimestamp(sqltypetime.Timestamp(a.Code.Time()))
			})

			entries, err := os.ReadDir(path)
			if err != nil {
				return err
			}

			omit := false

			a.Files = make([]*archivev1.File, 0, len(entries))

			for _, entry := range entries {
				name := entry.Name()

				if name == "backup_manifest" {
					omit = true
				}

				info, err := entry.Info()
				if err != nil {
					return err
				}

				a.Files = append(a.Files, &archivev1.File{
					Name:           entry.Name(),
					Size:           info.Size(),
					LastModifiedAt: sqltypetime.Timestamp(info.ModTime()),
				})
			}

			if omit {
				list.Add(a)
			}

			return filepath.SkipDir
		}

		return nil
	}); err != nil {
		return nil, err
	}

	list.Items = slices.SortedFunc(slices.Values(list.Items), func(a, b *archivev1.Archive) int {
		return strings.Compare(string(b.Code), string(a.Code))
	})

	return list, nil
}
