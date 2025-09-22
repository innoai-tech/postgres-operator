package internal

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/innoai-tech/postgres-operator/pkg/exec"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl/pgconf"
)

func InitDB(ctx context.Context, c pgconf.Conf) error {
	pgdata := c.DataDir.PgDataPath()

	if err := os.MkdirAll(pgdata, 0o777); err != nil {
		return err
	}

	pwFile, err := exec.WriteTempFile("pgpass", []byte(fmt.Sprintf("%s\n", c.Password)))
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(pwFile)
	}()

	user, err := lookupPostgresUser()
	if err != nil {
		return err
	}

	if err := os.Chown(pgdata, user.UID, user.GID); err != nil {
		return err
	}

	if err := os.Chown(pwFile, user.UID, user.GID); err != nil {
		return err
	}

	cmd := &exec.Command{
		Name: "initdb",
		UID:  user.UID,
		GID:  user.GID,
		Flags: exec.Flags{
			"-D":         {pgdata},
			"--username": {c.User},
			"--pwfile":   {pwFile},
		},
	}

	if err := cmd.Run(ctx); err != nil {
		return err
	}

	authMethod := &exec.Command{
		Name: "postgres",
		Flags: exec.Flags{
			"-D": {pgdata},
			"-C": {"password_encryption"},
		},
	}

	authMethod.UID = user.UID
	authMethod.GID = user.GID

	ret, err := authMethod.Output(ctx)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(filepath.Join(pgdata, "pg_hba.conf"), os.O_WRONLY|os.O_APPEND, 0o600)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, "host all all all %s\n", string(ret))
	return err
}
