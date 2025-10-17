package archive_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/innoai-tech/infra/pkg/configuration/testingutil"
	"github.com/innoai-tech/infra/pkg/otel"
	archivev1 "github.com/innoai-tech/postgres-operator/pkg/apis/archive/v1"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl/archive"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl/archive/sample"
	slicesx "github.com/octohelm/x/slices"
	"github.com/octohelm/x/testing/bdd"
)

func TestController(t *testing.T) {
	bdd.FromT(t).Given("pg data dir", func(b bdd.T) {
		ctx, d := testingutil.BuildContext(t, func(d *struct {
			otel.Otel

			archive.Controller
		},
		) {
			d.LogLevel = "debug"
			d.LogFormat = "text"

			d.DataDir = "./testdata/target"
		})

		archiveCodes := bdd.MustDo(func() ([]archivev1.ArchiveCode, error) {
			base, err := time.Parse(time.RFC3339, "2025-11-10T00:00:00+08:00")
			if err != nil {
				return nil, err
			}

			backup0 := archivev1.NewArchiveCode(base.Add(-24*time.Hour), "pg16")

			if err := sample.Backup(d.DataDir.PgArchivePath(string(backup0))); err != nil {
				return nil, err
			}

			backup1 := archivev1.NewArchiveCode(base, "pg16")
			if err := sample.Backup(d.DataDir.PgArchivePath(string(backup1))); err != nil {
				return nil, err
			}

			return []archivev1.ArchiveCode{
				backup1,
				backup0,
			}, nil
		})

		b.When("list archive", func(b bdd.T) {
			list, err := d.ListArchive(ctx)

			b.Then("success",
				bdd.NoError(err),
				bdd.Equal(
					archiveCodes,
					slicesx.Map(list.Items, func(a *archivev1.Archive) archivev1.ArchiveCode {
						return a.Code
					}),
				),
			)

			b.When("export", func(b bdd.T) {
				archiveCode := archiveCodes[1]

				exportedFile := filepath.Join(os.TempDir(), fmt.Sprintf("%s.tar", archiveCode))

				bdd.MustDo(func() (any, error) {
					f := bdd.Must(os.OpenFile(exportedFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm))
					defer f.Close()

					writeTo, err := d.ExportArchiveAsTar(ctx, archiveCodes[1])
					if err != nil {
						return nil, err
					}

					if err := writeTo(f); err != nil {
						return nil, err
					}
					return nil, err
				})

				b.When("remove the exported", func(b bdd.T) {
					b.Then("success",
						bdd.NoError(d.DeleteArchive(ctx, archiveCode)),
					)

					b.Then("only remain one archive",
						bdd.Equal(archiveCodes[0], bdd.Must(d.ListArchive(ctx)).Items[0].Code),
					)

					b.When("import", func(b bdd.T) {
						f := bdd.Must(os.Open(exportedFile))
						defer f.Close()

						b.Then("success",
							bdd.NoError(d.ImportArchiveFromTar(ctx, archiveCode, f)),
						)

						list, err := d.ListArchive(ctx)
						b.Then("list all",
							bdd.NoError(err),
							bdd.Equal(
								archiveCodes,
								slicesx.Map(list.Items, func(a *archivev1.Archive) archivev1.ArchiveCode {
									return a.Code
								}),
							),
						)
					})
				})
			})
		})

		b.When("request restore", func(b bdd.T) {
			b.Then("success",
				bdd.NoError(d.RequestRestore(ctx, archiveCodes[0])),
			)

			b.When("commit restore", func(b bdd.T) {
				b.Then("success",
					bdd.NoError(d.CommitRestore(ctx)),
				)
			})
		})
	})
}
