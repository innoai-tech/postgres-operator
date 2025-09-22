package openid

import (
	"context"

	"github.com/octohelm/courier/pkg/courierhttp"
	"github.com/octohelm/jwx/pkg/jwk"
)

// JWKs 密钥列表
// +gengo:injectable
type JWKs struct {
	courierhttp.MethodGet `path:"/.well-known/jwks.json"`

	keySetProvider jwk.KeySetProvider `inject:""`
}

func (req *JWKs) Output(ctx context.Context) (any, error) {
	return req.keySetProvider.TypedPublicSet()
}
