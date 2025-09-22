package internal

import (
	"fmt"
	"os/user"
	"strconv"
	"sync"
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
