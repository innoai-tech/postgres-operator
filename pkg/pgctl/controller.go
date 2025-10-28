package pgctl

import (
	"context"
	"time"

	archivev1 "github.com/innoai-tech/postgres-operator/pkg/apis/archive/v1"
	databasev1 "github.com/innoai-tech/postgres-operator/pkg/apis/database/v1"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl/archive"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl/internal/db"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl/pgconf"
	"github.com/octohelm/exp/xchan"
	metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"
	"github.com/octohelm/objectkind/pkg/runtime"
)

type EventType int

const (
	EventTypeShutdown EventType = iota + 1
	EventTypeBackup
	EventTypeReady
)

type Event struct {
	Type EventType
	Data any
}

// +gengo:injectable:provider
type Controller struct {
	pgconf.Conf

	sub xchan.Subject[Event]
}

func (c *Controller) DBController(ctx context.Context) *db.Controller {
	return db.New(c.Conf.ToDSN())
}

func (c *Controller) IsReady(ctx context.Context) error {
	return c.DBController(ctx).IsReady(ctx)
}

func (c *Controller) ListDatabase(ctx context.Context) (*metav1.List[databasev1.Database], error) {
	return c.DBController(ctx).ListDatabase(ctx)
}

func (c *Controller) ListTableOfDatabase(ctx context.Context, databaseName databasev1.DatabaseCode) (*metav1.List[databasev1.Table], error) {
	return c.DBController(ctx).ListTableOfDatabase(ctx, databaseName)
}

func (c *Controller) QueryDatabaseResult(ctx context.Context, databaseCode databasev1.DatabaseCode, sql string) (*databasev1.Result, error) {
	return c.DBController(ctx).WithName(databaseCode).QueryResult(ctx, sql)
}

func (c *Controller) ArchiveController() *archive.Controller {
	return &archive.Controller{DataDir: c.DataDir}
}

func (c *Controller) Observe() xchan.Observer[Event] {
	return c.sub.Observe()
}

func (c *Controller) CreateArchive(ctx context.Context) (*archivev1.Archive, error) {
	pgVersion, err := c.PgVersion(ctx)
	if err != nil {
		return nil, err
	}

	a := runtime.Build(func(a *archivev1.Archive) {
		a.Code = archivev1.NewArchiveCode(time.Now(), "pg"+pgVersion)
	})

	c.sub.Send(Event{Type: EventTypeBackup, Data: a})

	return a, nil
}

func (c *Controller) Restart(ctx context.Context) error {
	c.sub.Send(Event{Type: EventTypeShutdown})
	return nil
}

func (c *Controller) NotifyReady(ctx context.Context) error {
	c.sub.Send(Event{Type: EventTypeReady})
	return nil
}
