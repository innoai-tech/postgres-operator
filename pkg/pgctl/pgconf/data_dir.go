package pgconf

import (
	"path/filepath"
	"slices"
)

type DataDir string

func (dir DataDir) PgDataPath() string {
	return filepath.Join(string(dir), "pgdata")
}

func (dir DataDir) PgBackupPath() string {
	return filepath.Join(string(dir), "pgbackup")
}

type ArchiveDataDir string

func (dir ArchiveDataDir) PgArchivePath(subPaths ...string) string {
	return filepath.Join(slices.Concat(
		[]string{string(dir), "pgarchive"},
		subPaths,
	)...)
}
