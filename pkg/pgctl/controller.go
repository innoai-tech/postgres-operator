package pgctl

import (
	"context"
	"time"

	archivev1 "github.com/innoai-tech/postgres-operator/pkg/apis/archive/v1"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl/archive"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl/internal"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl/pgconf"
	"github.com/octohelm/exp/xchan"
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

func (c *Controller) IsReady(ctx context.Context) error {
	return internal.IsReady(ctx, c.Conf)
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

func (v *Controller) Restart(ctx context.Context) error {
	v.sub.Send(Event{Type: EventTypeShutdown})
	return nil
}

func (v *Controller) NotifyReady(ctx context.Context) error {
	v.sub.Send(Event{Type: EventTypeReady})
	return nil
}
