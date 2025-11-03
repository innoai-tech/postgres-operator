package archive

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/octohelm/x/logr"

	archivev1 "github.com/innoai-tech/postgres-operator/pkg/apis/archive/v1"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl/archive/internal"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl/pgconf"
)

type Controller struct {
	DataDir pgconf.DataDir
}

func (c *Controller) ImportArchiveFromTar(ctx context.Context, code archivev1.ArchiveCode, r io.ReadCloser) error {
	defer r.Close()

	return internal.Untar(r, c.DataDir.PgArchivePath(string(code)))
}

func (c *Controller) DeleteArchive(ctx context.Context, code archivev1.ArchiveCode) error {
	root := c.DataDir.PgArchivePath(string(code))
	return os.RemoveAll(root)
}

const RESTORE_REQUEST_FILENAME = "restore_request"

func (c *Controller) RequestRestore(ctx context.Context, code archivev1.ArchiveCode) error {
	restoreRequestPath := c.DataDir.PgArchivePath(RESTORE_REQUEST_FILENAME)
	if err := os.WriteFile(restoreRequestPath, []byte(code), 0o644); err != nil {
		return err
	}
	return nil
}

func (c *Controller) CancelRestore(ctx context.Context) error {
	restoreRequestPath := c.DataDir.PgArchivePath(RESTORE_REQUEST_FILENAME)
	return os.RemoveAll(restoreRequestPath)
}

func (c *Controller) CurrentRestoreRequest(ctx context.Context) (archivev1.ArchiveCode, error) {
	restoreRequestPath := c.DataDir.PgArchivePath(RESTORE_REQUEST_FILENAME)
	restoreRequest, err := os.ReadFile(restoreRequestPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}
	return archivev1.ArchiveCode(string(bytes.TrimSpace(restoreRequest))), nil
}

func (c *Controller) CommitRestore(pctx context.Context) (finalErr error) {
	restoreRequest, err := c.CurrentRestoreRequest(pctx)
	if err != nil {
		return err
	}

	if restoreRequest == "" {
		return nil
	}

	_, l := logr.FromContext(pctx).Start(pctx, "CommitRestore", slog.String("archive", string(restoreRequest)))
	defer l.End()

	defer func() {
		if finalErr == nil {
			l.Info("done")

			_ = os.RemoveAll(c.DataDir.PgArchivePath(RESTORE_REQUEST_FILENAME))
		} else {
			l.Error(finalErr)
		}
	}()

	baseTarGz, err := os.Open(c.DataDir.PgArchivePath(string(restoreRequest), "base.tar.gz"))
	if err != nil {
		return err
	}
	defer baseTarGz.Close()

	baseTar, err := gzip.NewReader(baseTarGz)
	if err != nil {
		return err
	}

	pgWalTarGz, err := os.Open(c.DataDir.PgArchivePath(string(restoreRequest), "pg_wal.tar.gz"))
	if err != nil {
		return err
	}
	defer pgWalTarGz.Close()

	pgWalTar, err := gzip.NewReader(pgWalTarGz)
	if err != nil {
		return err
	}

	pgDataPath := c.DataDir.PgDataPath()
	pgDataExists := true

	if _, err := os.Stat(pgDataPath); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		pgDataExists = false
	}

	if pgDataExists {
		pgDataTempPath := pgDataPath + ".tmp"
		if err := os.Rename(pgDataPath, pgDataTempPath); err != nil {
			return err
		}

		defer func() {
			if finalErr == nil {
				// clean tmp pg data
				_ = os.RemoveAll(pgDataTempPath)
			} else {
				// rollback
				_ = os.RemoveAll(pgDataPath)
				_ = os.Rename(pgDataTempPath, pgDataPath)
			}
		}()

		if err = os.RemoveAll(pgDataPath); err != nil {
			return err
		}
	}

	if err := internal.Untar(baseTar, pgDataPath); err != nil {
		return err
	}

	if err := internal.Untar(pgWalTar, filepath.Join(pgDataPath, "pg_wal")); err != nil {
		return err
	}

	return nil
}
