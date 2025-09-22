package openid

import (
	"context"

	"github.com/innoai-tech/postgres-operator/internal/cmd/postgres-operator/domain/auth"
	"github.com/octohelm/courier/pkg/courierhttp"
	openidv1 "github.com/octohelm/jwx/pkg/apis/openid/v1"
)

// ExchangeToken
// +gengo:injectable
type ExchangeToken struct {
	courierhttp.MethodPost `path:"/auth/token"`

	Authorization openidv1.Authorization `name:"Authorization,omitzero" in:"header"`
	GrantPayload  openidv1.GrantPayload  `in:"body" mime:"urlencoded"`

	as *auth.Service `inject:""`
}

func (r *ExchangeToken) Output(ctx context.Context) (any, error) {
	switch grant := r.GrantPayload.Grant.(type) {
	case *openidv1.PasswordGrant:
		return r.as.TokenByUserPassword(ctx, grant.Username, grant.Password)
	case *openidv1.ClientCredentialsGrant:
		return r.as.TokenByUserPassword(ctx, grant.ClientID, grant.ClientSecret)
	case *openidv1.RefreshTokenGrant:
		return r.as.TokenByRefreshToken(ctx, grant.RefreshToken)
	case *openidv1.AuthorizationCodeGrant:
		return nil, &openidv1.ErrUnsupportedGrantType{
			GrantType: grant.Type(),
		}
	}

	return nil, &openidv1.ErrUnsupportedGrantType{
		GrantType: "",
	}
}
