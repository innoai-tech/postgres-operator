package sample

import (
	"archive/tar"
	"compress/gzip"
	"embed"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/innoai-tech/postgres-operator/pkg/pgctl/archive/internal"
	"golang.org/x/sync/errgroup"
)

//go:embed pgdata
var src embed.FS

func Backup(outputDir string) error {
	if err := os.MkdirAll(outputDir, 0o700); err != nil {
		return err
	}

	eg := &errgroup.Group{}

	pgData, err := fs.Sub(src, "pgdata")
	if err != nil {
		return err
	}

	eg.Go(func() error {
		f, err := os.OpenFile(filepath.Join(outputDir, "base.tar.gz"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o600)
		if err != nil {
			return err
		}
		defer f.Close()

		return writeAsTarGz(internal.Include(pgData, func(path string, entry fs.DirEntry) bool {
			return !(strings.HasPrefix(path, "pg_wal/") || path == "pg_wal")
		}), f)
	})

	eg.Go(func() error {
		f, err := os.OpenFile(filepath.Join(outputDir, "pg_wal.tar.gz"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o600)
		if err != nil {
			return err
		}
		defer f.Close()

		pgWal, err := fs.Sub(pgData, "pg_wal")
		if err != nil {
			return err
		}

		return writeAsTarGz(pgWal, f)
	})

	if err := eg.Wait(); err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(outputDir, "backup_manifest"), []byte("{}"), 0o600)
}

func writeAsTarGz(src fs.FS, w io.Writer) error {
	gw := gzip.NewWriter(w)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	if err := tw.AddFS(src); err != nil {
		return err
	}

	return nil
}
