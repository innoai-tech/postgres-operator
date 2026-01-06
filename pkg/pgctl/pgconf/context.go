package pgconf

import (
	"cmp"
	"context"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// +gengo:injectable:provider
type Provider interface {
	GetPgConf() Conf
}

type Conf struct {
	// DataDir db data-dir
	DataDir DataDir `flag:""`
	// ArchiveDataDir archive data-dir
	ArchiveDataDir ArchiveDataDir `flag:",omitzero"`
	// PgVersion pg bin version, don't set this unless you know what will be happen
	PgVersion string `flag:",omitzero"`

	Database

	Setting
}

func (c *Conf) SetDefaults() {
	c.Setting.SetDefaults()

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

func (d *Conf) PgDataVersion(ctx context.Context) (string, error) {
	pgVersion, err := os.ReadFile(filepath.Join(d.DataDir.PgDataPath(), "PG_VERSION"))
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}
	return strings.TrimSpace(string(pgVersion)), nil
}

func (c *Conf) GetDataDir() DataDir {
	return c.DataDir
}

func (c *Conf) GetArchiveDataDir() ArchiveDataDir {
	return cmp.Or(c.ArchiveDataDir, ArchiveDataDir(c.DataDir))
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
