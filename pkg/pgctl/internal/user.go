package internal

import (
	"context"
	"fmt"
	"os/user"
	"slices"
	"strconv"
	"sync"

	"github.com/innoai-tech/postgres-operator/pkg/exec"
)

type User struct {
	UID int
	GID int
}

var lookupPostgresUser = sync.OnceValues(func() (*User, error) {
	u, err := user.Lookup("postgres")
	if err != nil {
		return nil, fmt.Errorf("lookup user failed: %w", err)
	}

	uid, err := strconv.Atoi(u.Uid)
	if err != nil {
		return nil, fmt.Errorf("invalid uid %s: %w", u.Uid, err)
	}
	gid, err := strconv.Atoi(u.Gid)
	if err != nil {
		return nil, fmt.Errorf("invalid gid %s: %w", u.Gid, err)
	}

	return &User{UID: uid, GID: gid}, nil
})

func postgresUserChown(ctx context.Context, dirs ...string) error {
	if len(dirs) == 0 {
		return nil
	}

	chown := &exec.Command{
		Name: "chown",
		Args: slices.Concat([]string{
			"-R",
			"postgres:postgres",
		}, dirs),
	}

	return chown.Run(ctx)
}
