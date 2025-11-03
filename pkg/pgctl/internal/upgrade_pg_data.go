package internal

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/innoai-tech/postgres-operator/pkg/exec"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl/pgconf"
)

type ToolchainDir string

func (dir ToolchainDir) UsrBinDir(pgVersion string) string {
	return filepath.Join(string(dir), "usr/lib/postgresql", pgVersion, "bin")
}

func (dir ToolchainDir) UsrLibDir(pgVersion string) string {
	return filepath.Join(string(dir), "usr/lib/postgresql", pgVersion, "lib")
}

func (dir ToolchainDir) UsrShareDir(pgVersion string) string {
	return filepath.Join(string(dir), "usr/share/postgresql", pgVersion, "lib")
}

func (dir ToolchainDir) CreateLinks(dirs ...string) error {
	base := string(dir)

	for _, d := range dirs {
		if strings.HasPrefix(d, string(base)) {
			target := d[len(base):]
			if err := os.MkdirAll(filepath.Dir(target), 0777); err != nil {
				return err
			}
			if err := os.Symlink(d, d[len(base):]); err != nil {
				return err
			}
		}
	}

	return nil
}

var ErrMissingToolchain = errors.New("missing toolchain")

func UpgradePgData(ctx context.Context, c pgconf.Conf, oldPgVersion string) (finalErr error) {
	t := ToolchainDir("/postgres-toolchain")

	if err := t.CreateLinks(
		t.UsrBinDir(oldPgVersion),
		t.UsrLibDir(oldPgVersion),
		t.UsrShareDir(oldPgVersion),
	); err != nil {
		if os.IsNotExist(err) {
			return errors.Join(ErrMissingToolchain, err)
		}
		return err
	}

	pgData := c.DataDir.PgDataPath()
	pgDataOld := pgData + ".old"

	err := os.Rename(pgData, pgDataOld)
	if err != nil {
		return err
	}

	defer func() {
		if finalErr != nil {
			// rollback to old pg data
			_ = os.RemoveAll(pgData)
			_ = os.Rename(pgDataOld, pgData)
		}
	}()

	if err := InitDB(ctx, c); err != nil {
		return err
	}

	if err := postgresUserChown(ctx, pgData, pgDataOld); err != nil {
		return err
	}

	user, err := lookupPostgresUser()
	if err != nil {
		return err
	}

	cmd := &exec.Command{
		WorkDir: pgData,
		Name:    filepath.Join("/usr/lib/postgresql", c.PgVersion, "bin/pg_upgrade"),
		UID:     user.UID,
		GID:     user.GID,
		Flags: exec.Flags{
			"-U": {c.User},
			"-b": {filepath.Join("/usr/lib/postgresql", oldPgVersion, "bin")},
			"-d": {pgDataOld},
			"-B": {filepath.Join("/usr/lib/postgresql", c.PgVersion, "bin")},
			"-D": {pgData},
		},
	}

	return cmd.Run(ctx)
}

func CompleteUpgradeIfNeed(ctx context.Context, c pgconf.Conf) error {
	pgData := c.DataDir.PgDataPath()

	deleteOldClusterScript := filepath.Join(pgData, "delete_old_cluster.sh")

	_, err := os.Stat(deleteOldClusterScript)
	if err != nil {
		// if not exists, should skip
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	user, err := lookupPostgresUser()
	if err != nil {
		return err
	}

	for _, args := range [][]string{
		{"--all", "--analyze-in-stages", "--missing-stats-only"},
		{"--all", "--analyze-only"},
	} {
		cmd := &exec.Command{
			WorkDir: pgData,
			Name:    filepath.Join("/usr/lib/postgresql", c.PgVersion, "bin/vacuumdb"),
			UID:     user.UID,
			GID:     user.GID,
			Args:    slices.Concat([]string{"-U", c.User}, args),
		}

		if err := cmd.Run(ctx); err != nil {
			return err
		}
	}

	// do final old data cleanup
	cmd := &exec.Command{
		WorkDir: pgData,
		Name:    "sh",
		UID:     user.UID,
		GID:     user.GID,
		Args:    []string{deleteOldClusterScript},
	}
	if err := cmd.Run(ctx); err != nil {
		return err
	}

	return nil
}
