package internal

import (
	"context"
	"os"

	archivev1 "github.com/innoai-tech/postgres-operator/pkg/apis/archive/v1"
	"github.com/innoai-tech/postgres-operator/pkg/exec"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl/pgconf"
	"github.com/octohelm/x/logr"
)

func BaseBackup(ctx context.Context, c pgconf.Conf, code archivev1.ArchiveCode) error {
	user, err := lookupPostgresUser()
	if err != nil {
		return err
	}

	backupPath := c.DataDir.PgBackupPath()

	if err := os.RemoveAll(backupPath); err != nil {
		return err
	}

	cmd := &exec.Command{
		Name:    "pg_basebackup",
		UID:     user.UID,
		GID:     user.GID,
		WorkDir: string(c.DataDir),
		Args: []string{
			"-U", c.User,
			"-D", backupPath,
			"-F", "tar", "-z",
			"-X", "stream",
			"--checkpoint", "fast",
			"--verbose",
			"-P",
		},
	}

	cctx, l := logr.FromContext(ctx).Start(ctx, "BaseBackup")
	defer l.End()

	l.Info("backing up")

	if err := cmd.Run(cctx); err != nil {
		return err
	}

	if err := os.MkdirAll(c.DataDir.PgArchivePath(), os.ModePerm); err != nil {
		return err
	}

	return os.Rename(backupPath, c.DataDir.PgArchivePath(string(code)))
}
