package convert

import (
	runtimeconverter "github.com/octohelm/objectkind/pkg/runtime/converter"

	productv1 "github.com/innoai-tech/postgres-operator/internal/example/apis/product/v1"
	"github.com/innoai-tech/postgres-operator/internal/example/domain/product"
)

var Product = runtimeconverter.ForCodableObject(
	func(o *productv1.Product, m *product.Product) error {
		o.Status.State = m.State
		return nil
	},
	func(m *product.Product, o *productv1.Product) error {
		m.State = o.Status.State
		return nil
	},
)
