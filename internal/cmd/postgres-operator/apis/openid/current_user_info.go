package openid

import (
	"context"

	"github.com/innoai-tech/postgres-operator/internal/cmd/postgres-operator/domain/auth"
	"github.com/octohelm/courier/pkg/courierhttp"
	openidv1 "github.com/octohelm/jwx/pkg/apis/openid/v1"
	"github.com/octohelm/jwx/pkg/openid"
)

// CurrentUserInfo
// +gengo:injectable
type CurrentUserInfo struct {
	courierhttp.MethodGet `path:"/user/info"`

	Authorization openidv1.Authorization `name:"Authorization" in:"header"`

	httpRequest *courierhttp.Request `inject:""`

	as *auth.Service `inject:""`
}

func (r *CurrentUserInfo) Output(ctx context.Context) (any, error) {
	accessToken := r.Authorization.Get(openidv1.TokenTypeBearer)

	user, err := r.as.ValidateAccessToken(ctx, accessToken)
	if err != nil {
		return nil, courierhttp.WrapError(err, openid.WithWwwAuthenticate(r.httpRequest))
	}

	return user, nil
}
