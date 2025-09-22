package pgconf

import (
	"context"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/innoai-tech/postgres-operator/pkg/units"
)

// +gengo:injectable:provider
type Provider interface {
	GetPgConf() Conf
}

type Conf struct {
	// DataDir db data-dir
	DataDir DataDir `flag:""`

	Database

	ResourceRequests
}

func (c *Conf) SetDefaults() {
	if c.CPU == 0 {
		c.CPU = 2
	}

	if c.MEM == 0 {
		c.MEM = 4 * units.GiB
	}

	if c.MaxConnections == 0 {
		c.MaxConnections = 100 * c.CPU
	}

	if c.Port == 0 {
		c.Port = 5432
	}
}

func (c *Conf) ToDSN() *url.URL {
	db := &url.URL{}

	db.Scheme = "postgres"
	db.Host = net.JoinHostPort("0.0.0.0", strconv.Itoa(int(c.Port)))
	db.Path = "/" + c.Name
	db.User = url.UserPassword(c.User, c.Password)

	values := &url.Values{}
	values.Set("sslmode", "disable")

	db.RawQuery = values.Encode()

	return db
}

func (d *Conf) PgVersion(ctx context.Context) (string, error) {
	pgVersion, err := os.ReadFile(filepath.Join(d.DataDir.PgDataPath(), "PG_VERSION"))
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}
	return strings.TrimSpace(string(pgVersion)), nil
}

type Database struct {
	// Name db name
	Name string `flag:""`
	// User db user
	User string `flag:""`
	// Password db password
	Password string `flag:",secret"`
	// Port db listen port
	Port uint16 `flag:",omitzero"`
}

type ResourceRequests struct {
	// CPU db cpu requests
	CPU int `flag:",omitzero"`
	// MEM db mem requests
	MEM units.BinarySize `flag:",omitzero"`
	// MaxConnections db max connections
	MaxConnections int `flag:",omitzero"`
}
