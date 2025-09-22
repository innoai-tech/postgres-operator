package repository_test

import (
	"fmt"
	"testing"

	"github.com/innoai-tech/infra/pkg/configuration/testingutil"
	"github.com/innoai-tech/infra/pkg/otel"
	productv1 "github.com/innoai-tech/postgres-operator/internal/example/apis/product/v1"
	transactionv1 "github.com/innoai-tech/postgres-operator/internal/example/apis/transaction/v1"
	"github.com/innoai-tech/postgres-operator/internal/example/domain/product"
	productfilter "github.com/innoai-tech/postgres-operator/internal/example/domain/product/filter"
	productrepository "github.com/innoai-tech/postgres-operator/internal/example/domain/product/repository"
	"github.com/octohelm/objectkind/pkg/idgen"
	"github.com/octohelm/objectkind/pkg/runtime"
	"github.com/octohelm/storage/pkg/filter"
	sessiondb "github.com/octohelm/storage/pkg/session/db"
	"github.com/octohelm/x/testing/bdd"
)

func TestSkuRepository(t *testing.T) {
	ctx, d := testingutil.BuildContext(t, func(d *struct {
		otel.Otel
		idgen.IDGen

		sessiondb.Database

		productrepository.ProductRepository
		productrepository.SkuRepository
	},
	) {
		d.LogLevel = "debug"
		d.LogFormat = "text"
		d.EnableMigrate = true

		d.ApplyCatalog("test", product.T)
	})

	pdt := bdd.MustDo(func() (*productv1.Product, error) {
		pdt := runtime.Build(func(pdt *productv1.Product) {
			pdt.Code = "pdt"
			pdt.Name = "测试产品"
		})

		if err := d.PutProduct(ctx, pdt); err != nil {
			return nil, err
		}

		return pdt, nil
	})

	bdd.FromT(t).Given("skus", func(b bdd.T) {
		skus := make([]*productv1.Sku, 10)

		for i := range skus {
			skus[i] = runtime.Build(func(sku *productv1.Sku) {
				sku.Code = productv1.SkuCode(fmt.Sprintf("pdt-%03d", i))

				sku.Spec.Price = transactionv1.CurrencyValue(1 + i)
				sku.Spec.Currency = transactionv1.CurrencyCNY
			})
		}

		b.When("do put the skus", func(b bdd.T) {
			b.Then("success",
				bdd.NoError(
					d.PutSkuOfProduct(ctx, pdt, skus...),
				),
			)

			b.When("list of product", func(b bdd.T) {
				skuList, err := d.ListSku(ctx, &productfilter.SkuByProductID{
					ProductID: filter.Eq(pdt.ID),
				})

				b.Then("success",
					bdd.NoError(err),
					bdd.Equal(10, len(skuList.Items)),
				)
			})

			b.When("list of product with skus", func(b bdd.T) {
				productList, err := d.ListProduct(ctx, &productfilter.ProductByID{
					ID: filter.Eq(pdt.ID),
				})

				b.Then("success",
					bdd.NoError(err),
					bdd.Equal(10, len(productList.Items[0].Skus)),
				)
			})
		})
	})
}
