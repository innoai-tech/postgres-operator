package db

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func isErrorUnknownDatabase(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "3D000" {
			return true
		}
	}
	return false
}
