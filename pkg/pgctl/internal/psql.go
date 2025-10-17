package internal

import (
	"context"
	"os"
	"strings"

	"github.com/innoai-tech/postgres-operator/pkg/exec"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl/pgconf"
)

func PSQL(ctx context.Context, c pgconf.Conf, sql string) error {
	user, err := lookupPostgresUser()
	if err != nil {
		return err
	}

	pwFile, err := exec.WriteTempFile("psql", []byte(strings.TrimSpace(sql)))
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(pwFile)
	}()

	if err := os.Chown(pwFile, user.UID, user.GID); err != nil {
		return err
	}

	cmd := &exec.Command{
		Name: "psql",
		UID:  user.UID,
		GID:  user.GID,
		Flags: exec.Flags{
			"-U": {c.User},
			"-f": {pwFile},
		},
	}

	if err := cmd.Run(ctx); err != nil {
		return err
	}

	return err
}
