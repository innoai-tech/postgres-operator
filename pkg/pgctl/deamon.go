package pgctl

import (
	"context"
	"errors"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/innoai-tech/infra/pkg/configuration"
	"github.com/octohelm/exp/xchan"
	"github.com/octohelm/x/logr"
	syncx "github.com/octohelm/x/sync"
	"github.com/octohelm/x/sync/singleflight"

	archivev1 "github.com/innoai-tech/postgres-operator/pkg/apis/archive/v1"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl/internal"
)

// +gengo:injectable
type Daemon struct {
	c *Controller `inject:""`

	// WithoutPg serve without postgres
	Off bool `flag:",omitzero"`
	// ExitOnError
	ExitOnError bool `flag:",omitzero"`

	sfg singleflight.Group[string]

	done         atomic.Bool
	wg           sync.WaitGroup
	processQueue chan configuration.Server
	processes    syncx.Map[configuration.Server, bool]
}

func (d *Daemon) afterInit(ctx context.Context) error {
	pgDataVersion, err := d.c.Conf.PgDataVersion(ctx)
	if err != nil {
		return err
	}

	if pgDataVersion == "" {
		if err := internal.InitDB(ctx, d.c.Conf); err != nil {
			return err
		}
	} else if d.c.Conf.PgVersion != "" {
		if d.c.Conf.PgVersion != pgDataVersion {
			if maxVersion(d.c.Conf.PgVersion, pgDataVersion) == d.c.Conf.PgVersion {
				if err := internal.UpgradePgData(ctx, d.c.Conf, pgDataVersion); err != nil {
					return err
				}
			} else {
				return errors.New("downgrade is not allowed")
			}
		}
	}

	d.processQueue = make(chan configuration.Server)

	return nil
}

func maxVersion(v1 string, v2 string) string {
	ver1, _ := strconv.ParseInt(v1, 10, 64)
	ver2, _ := strconv.ParseInt(v2, 10, 64)
	if ver1 > ver2 {
		return v1
	}
	return v2
}

var _ configuration.CanDisabled = &Daemon{}

func (d *Daemon) Disabled(ctx context.Context) bool {
	return d.Off
}

var _ configuration.Server = &Daemon{}

func (d *Daemon) Serve(gctx context.Context) error {
	injector := configuration.ContextInjectorFromContext(gctx)

	l := logr.FromContext(gctx)

	c, cancel := context.WithCancel(gctx)
	defer cancel()

	go func() {
		for ct := range xchan.Values(c, d.c.Observe()) {
			switch ct.Type {
			case EventTypeReady:
				d.wg.Go(func() {
					err, _ := d.sfg.Do("ready", func() error {
						defer d.sfg.Forget("ready")

						ctx := injector.InjectContext(context.Background())

						if err := internal.CompleteUpgradeIfNeed(ctx, d.c.Conf); err != nil {
							return err
						}

						if err := internal.PresetDB(ctx, d.c.Conf); err != nil {
							return err
						}

						return nil
					})
					if err != nil {
						l.Error(err)
					}
				})
			case EventTypeBackup:
				d.wg.Go(func() {
					err, _ := d.sfg.Do("backup", func() error {
						defer d.sfg.Forget("backup")

						ctx := injector.InjectContext(context.Background())

						return internal.BaseBackup(ctx, d.c.Conf, ct.Data.(*archivev1.Archive).Code)
					})

					if err != nil {
						l.Error(err)
					}
				})
			case EventTypeShutdown:
				d.wg.Go(func() {
					err, _ := d.sfg.Do("shutdown", func() error {
						defer d.sfg.Forget("shutdown")

						ctx := injector.InjectContext(context.Background())
						return d.shutdownAllRunning(ctx)
					})

					if err != nil {
						l.Error(err)
					}
				})
			}
		}
	}()

	go func() {
		for server := range d.processQueue {
			d.wg.Go(func() {
				ctx := injector.InjectContext(context.Background())

				t := time.NewTicker(1 * time.Second)
				defer t.Stop()

				for range t.C {
					if d.c.IsReady(ctx) == nil {
						_ = d.c.NotifyReady(ctx)
						break
					}
				}
			})

			d.wg.Go(func() {
				ctx := injector.InjectContext(context.Background())

				err := server.Serve(ctx)
				if err != nil {
					l.Error(err)

					if d.ExitOnError {
						os.Exit(1)
						return
					}
				}

				d.serving(&Postgres{Conf: d.c.Conf}, server)
			})
		}
	}()

	d.serving(&Postgres{
		Conf: d.c.Conf,
	})

	// for serve until shutdown
	d.wg.Add(1)
	d.wg.Wait()

	return nil
}

func (d *Daemon) shutdownAllRunning(ctx context.Context) error {
	er := &errgroup.Group{}

	for server := range d.processes.Range {
		er.Go(func() error {
			return server.Shutdown(ctx)
		})
	}

	return er.Wait()
}

func (d *Daemon) ShutdownTimeout(ctx context.Context) time.Duration {
	return 60 * time.Second
}

func (d *Daemon) Shutdown(ctx context.Context) error {
	if !d.done.Swap(true) {
		d.wg.Done()
		close(d.processQueue)

		return d.shutdownAllRunning(ctx)
	}

	return nil
}

func (r *Daemon) serving(server configuration.Server, deadServers ...configuration.Server) {
	if !r.done.Load() {
		r.processes.Store(server, true)

		r.processQueue <- server

		for _, e := range deadServers {
			r.processes.Delete(e)
		}
	}
}
