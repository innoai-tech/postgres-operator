package internal

import (
	"context"
	"fmt"

	"github.com/innoai-tech/postgres-operator/pkg/pgctl/pgconf"
)

func PresetDB(ctx context.Context, c pgconf.Conf) error {
	if err := PSQL(ctx, c, fmt.Sprintf(`
ALTER DATABASE %s REFRESH COLLATION VERSION;
`, c.Database.Name)); err != nil {
		return err
	}

	if err := PSQL(ctx, c, fmt.Sprintf(`
ALTER USER %s WITH PASSWORD '%s'
`, c.Database.User, c.Database.Password)); err != nil {
		return err
	}
	return nil
}
