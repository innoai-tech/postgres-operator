package main_testing

import (
	"fmt"
	"slices"
	"testing"
	"time"

	"github.com/innoai-tech/infra/pkg/configuration/testingutil"
	"github.com/innoai-tech/infra/pkg/otel"
	"github.com/innoai-tech/postgres-operator/integration-testing/internal/postgresoperatorclient"
	productv1 "github.com/innoai-tech/postgres-operator/internal/example/apis/product/v1"
	transactionv1 "github.com/innoai-tech/postgres-operator/internal/example/apis/transaction/v1"
	"github.com/innoai-tech/postgres-operator/internal/example/domain/product"
	productfilter "github.com/innoai-tech/postgres-operator/internal/example/domain/product/filter"
	productrepository "github.com/innoai-tech/postgres-operator/internal/example/domain/product/repository"
	archivev1 "github.com/innoai-tech/postgres-operator/pkg/apis/archive/v1"
	"github.com/octohelm/objectkind/pkg/idgen"
	"github.com/octohelm/objectkind/pkg/runtime"
	"github.com/octohelm/storage/pkg/filter"
	sessiondb "github.com/octohelm/storage/pkg/session/db"
	"github.com/octohelm/x/testing/bdd"
)

func TestBackupAndRestore(t *testing.T) {
	ctx, d := testingutil.BuildContext(t, func(d *struct {
		otel.Otel
		idgen.IDGen

		sessiondb.Database

		PostgresOperatorClient postgresoperatorclient.Client

		productrepository.ProductRepository
		productrepository.SkuRepository
	},
	) {
		d.LogLevel = "debug"
		d.LogFormat = "text"
		d.EnableMigrate = true

		d.Endpoint.Scheme = "postgres"
		d.Endpoint.Hostname = "0.0.0.0"
		d.Endpoint.Port = 35432

		d.Endpoint.Username = "test"
		d.Endpoint.Password = "test"
		d.Endpoint.Path = "/test"

		d.PostgresOperatorClient.Endpoint.Scheme = "http"
		d.PostgresOperatorClient.Endpoint.Hostname = "0.0.0.0"
		d.PostgresOperatorClient.Endpoint.Port = 8080

		d.PostgresOperatorClient.Endpoint.Username = d.Endpoint.Username
		d.PostgresOperatorClient.Endpoint.Password = d.Endpoint.Password

		d.ApplyCatalog("test", product.T)
	})

	bdd.FromT(t).Given("create initial data", func(b bdd.T) {
		pdt := bdd.MustDo(func() (*productv1.Product, error) {
			pdt := runtime.Build(func(pdt *productv1.Product) {
				pdt.Code = "pdt"
				pdt.Name = "测试产品"
			})

			if err := d.PutProduct(ctx, pdt); err != nil {
				return nil, err
			}

			skus := make([]*productv1.Sku, 1)

			for i := range skus {
				skus[i] = runtime.Build(func(sku *productv1.Sku) {
					sku.Code = productv1.SkuCode(fmt.Sprintf("pdt-%03d", i))

					sku.Spec.Price = 1
					sku.Spec.Currency = transactionv1.CurrencyCNY
				})
			}

			if err := d.PutSkuOfProduct(ctx, pdt, skus...); err != nil {
				return nil, err
			}

			return pdt, nil
		})

		b.When("create archive", func(b bdd.T) {
			x, err := postgresoperatorclient.DoWith(ctx, &d.PostgresOperatorClient, func(req *postgresoperatorclient.CreateArchive) {
			})

			b.Then("success",
				bdd.NoError(err),
			)

			for {
				list := bdd.Must(postgresoperatorclient.DoWith(ctx, &d.PostgresOperatorClient, func(req *postgresoperatorclient.ListArchive) {
				}))

				if slices.ContainsFunc(list.Items, func(a *archivev1.Archive) bool {
					return a.Code == x.Code
				}) {
					break
				}

				time.Sleep(1 * time.Second)
				continue
			}

			b.When("delete pdt", func(b bdd.T) {
				b.Then("success",
					bdd.NoError(d.DeleteAllProduct(ctx, &productfilter.ProductByID{
						ID: filter.Eq(pdt.ID),
					})),
				)

				b.When("restore from archive", func(b bdd.T) {
					_, err := postgresoperatorclient.DoWith(ctx, &d.PostgresOperatorClient, func(req *postgresoperatorclient.RequestRestoreArchive) {
						req.ArchiveCode = bdd.Must(postgresoperatorclient.DoWith(ctx, &d.PostgresOperatorClient, func(req *postgresoperatorclient.ListArchive) {
						})).Items[0].Code
					})

					b.Then("success",
						bdd.NoError(err),
					)

					b.When("restart postgres for storing archive", func(b bdd.T) {
						_, err := postgresoperatorclient.DoWith(ctx, &d.PostgresOperatorClient, func(req *postgresoperatorclient.Restart) {
						})

						b.Then("success",
							bdd.NoError(err),
						)

						// wait db ready
						for {
							_, err := postgresoperatorclient.DoWith(ctx, &d.PostgresOperatorClient, func(req *postgresoperatorclient.Liveness) {
							})
							if err == nil {
								break
							}

							time.Sleep(1 * time.Second)
							continue
						}

						// try to remove old connects
						_, _ = d.ListProduct(ctx)

						list, err := d.ListProduct(ctx)
						b.Then("data restored success",
							bdd.NoError(err),
							bdd.Equal(pdt.ID, list.Items[0].ID),
						)
					})
				})
			})
		})
	})
}
