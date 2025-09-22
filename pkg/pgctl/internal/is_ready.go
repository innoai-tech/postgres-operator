package internal

import (
	"context"
	"errors"
	"fmt"

	"github.com/innoai-tech/postgres-operator/pkg/pgctl/pgconf"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/octohelm/storage/pkg/sqlfrag"
)

func IsReady(ctx context.Context, c pgconf.Conf) error {
	pgConn, err := pgx.Connect(ctx, c.ToDSN().String())
	if err != nil {
		if isErrorUnknownDatabase(err) {
			return createDatabase(ctx, c)
		}
		return err
	}
	defer pgConn.Close(ctx)

	if err := pgConn.Ping(ctx); err != nil {
		if isErrorUnknownDatabase(err) {
			return createDatabase(ctx, c)
		}
		return err
	}
	return nil
}

func createDatabase(ctx context.Context, c pgconf.Conf) error {
	dsn := c.ToDSN()
	dsn.Path = "/postgres"

	pgConn, err := pgx.Connect(ctx, dsn.String())
	if err != nil {
		return fmt.Errorf("connect to database failed: %w", err)
	}
	defer pgConn.Close(ctx)

	sql, args := sqlfrag.Collect(ctx, sqlfrag.Pair("CREATE DATABASE ?;", sqlfrag.Const(c.Name)))
	if _, err := pgConn.Exec(ctx, sql, args...); err != nil {
		return fmt.Errorf("database creation failed: %w", err)
	}

	return nil
}

func isErrorUnknownDatabase(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "3D000" {
			return true
		}
	}
	return false
}
