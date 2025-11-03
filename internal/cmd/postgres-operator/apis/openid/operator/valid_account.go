package operator

import (
	"context"

	"github.com/octohelm/courier/pkg/courierhttp"
	openidv1 "github.com/octohelm/jwx/pkg/apis/openid/v1"
	"github.com/octohelm/jwx/pkg/openid"

	"github.com/innoai-tech/postgres-operator/internal/cmd/postgres-operator/domain/auth"
)

// +gengo:injectable:provider
type ValidAccount struct {
	Authorization openidv1.Authorization `name:"Authorization" in:"header"`

	httpRequest *courierhttp.Request `inject:""`

	as *auth.Service `inject:""`
}

func (c *ValidAccount) Output(ctx context.Context) (any, error) {
	accessToken := c.Authorization.Get(openidv1.TokenTypeBearer)

	_, err := c.as.ValidateAccessToken(ctx, accessToken)
	if err != nil {
		return nil, courierhttp.WrapError(err, openid.WithWwwAuthenticate(c.httpRequest))
	}

	return nil, nil
}
