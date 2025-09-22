package operator

import (
	"context"

	"github.com/innoai-tech/postgres-operator/internal/cmd/postgres-operator/domain/auth"
	"github.com/octohelm/courier/pkg/courierhttp"
	openidv1 "github.com/octohelm/jwx/pkg/apis/openid/v1"
	"github.com/octohelm/jwx/pkg/openid"
)

// +gengo:injectable:provider
type BaseURL struct {
	courierhttp.MethodGet

	Authorization openidv1.Authorization `name:"Authorization,omitzero" in:"header"`

	httpRequest *courierhttp.Request `inject:""`

	as *auth.Service `inject:""`
}

func (r *BaseURL) Output(ctx context.Context) (any, error) {
	accessToken := r.Authorization.Get(openidv1.TokenTypeBearer)

	_, err := r.as.ValidateAccessToken(ctx, accessToken)
	if err != nil {
		return nil, courierhttp.WrapError(err, openid.WithWwwAuthenticate(r.httpRequest))
	}

	return map[string]string{}, nil
}
