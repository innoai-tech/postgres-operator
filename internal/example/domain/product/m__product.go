package product

import (
	sqltypecompose "github.com/octohelm/objectkind/pkg/sqltype/compose"

	productv1 "github.com/innoai-tech/postgres-operator/internal/example/apis/product/v1"
)

// Product
// +gengo:table:register=T
// +gengo:table:group=product
// @def primary ID
// @def unique_index i_code Code
// @def index i_state State
type Product struct {
	sqltypecompose.CodableResource[productv1.ProductID, productv1.ProductCode]

	State productv1.ProductState `db:"state,default=1"`
}
