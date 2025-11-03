package sidecar

import (
	"context"
	"time"

	"github.com/innoai-tech/infra/pkg/agent"
	"github.com/innoai-tech/infra/pkg/cron"

	archivev1 "github.com/innoai-tech/postgres-operator/pkg/apis/archive/v1"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl"
)

// AutoArchiver
// +gengo:injectable
type AutoArchiver struct {
	agent.Agent

	Period      cron.Spec `flag:",omitzero"`
	CleanPeriod cron.Spec `flag:",omitzero"`

	MaxArchivesInSameDay int `flag:",omitzero"`
	KeepUntilDays        int `flag:",omitzero"`

	c *pgctl.Controller `inject:""`
}

func (aa *AutoArchiver) SetDefaults() {
	if aa.Period == "" {
		aa.Period = "@never"
	}

	if aa.CleanPeriod == "" {
		aa.CleanPeriod = "@midnight"
	}

	if aa.MaxArchivesInSameDay == 0 {
		aa.MaxArchivesInSameDay = 1
	}

	if aa.KeepUntilDays == 0 {
		aa.KeepUntilDays = 7
	}
}

func (aa *AutoArchiver) Disabled(ctx context.Context) bool {
	return aa.Period.Schedule() == nil
}

func (aa *AutoArchiver) afterInit(ctx context.Context) error {
	if aa.Disabled(ctx) {
		return nil
	}

	aa.Host("create archive", func(ctx context.Context) error {
		for range aa.Period.Times(ctx) {
			aa.Go(ctx, func(ctx context.Context) error {
				_, err := aa.c.CreateArchive(ctx)
				return err
			})
		}
		return nil
	})

	aa.Host("clean old archives", func(ctx context.Context) error {
		for range aa.CleanPeriod.Times(ctx) {
			aa.Go(ctx, func(ctx context.Context) error {
				return aa.cleanOldArchives(ctx)
			})
		}
		return nil
	})

	return nil
}

func (aa *AutoArchiver) cleanOldArchives(ctx context.Context) error {
	ac := aa.c.ArchiveController()

	list, err := ac.ListArchive(ctx)
	if err != nil {
		return err
	}

	toDeletes := map[archivev1.ArchiveCode]struct{}{}
	yearDayCounts := map[int]int{}

	now := time.Now()

	minTime := now.Add(-time.Duration(aa.KeepUntilDays) * 24 * time.Hour)

	for _, a := range list.Items {
		archiveTime := a.Code.Time()

		if now.Sub(archiveTime) < 24*time.Hour {
			continue
		}

		if archiveTime.Before(minTime) {
			toDeletes[a.Code] = struct{}{}
			continue
		}

		yearDay := archiveTime.YearDay()

		yearDayCounts[yearDay]++

		if yearDayCounts[yearDay] > aa.MaxArchivesInSameDay {
			toDeletes[a.Code] = struct{}{}
		}
	}

	for a := range toDeletes {
		if err := ac.DeleteArchive(ctx, a); err != nil {
			return err
		}
	}

	return nil
}
