package auth

import (
	"context"
	"time"

	"github.com/innoai-tech/postgres-operator/pkg/pgctl"
	openidv1 "github.com/octohelm/jwx/pkg/apis/openid/v1"
	"github.com/octohelm/jwx/pkg/sign"
)

// +gengo:injectable:provider
type Service struct {
	pgctl  *pgctl.Controller `inject:""`
	signer sign.Signer       `inject:""`
}

func (a *Service) ValidateAccessToken(ctx context.Context, accessToken string) (*openidv1.UserInfo, error) {
	tok, err := a.signer.Validate(ctx, accessToken, sign.WithClaimExpect(ClaimTokenType, TokenTypeAccessToken))
	if err != nil {
		return nil, err
	}

	return &openidv1.UserInfo{
		Sub: tok.Subject(),
	}, nil
}

func (a *Service) TokenByRefreshToken(ctx context.Context, refreshToken string) (*openidv1.Token, error) {
	tok, err := a.signer.Validate(ctx, refreshToken, sign.WithClaimExpect(ClaimTokenType, TokenTypeRefreshToken))
	if err != nil {
		return nil, err
	}

	return a.SignNewToken(ctx, tok.Subject())
}

func (a *Service) TokenByUserPassword(ctx context.Context, username string, password string) (*openidv1.Token, error) {
	if a.pgctl.User == username && a.pgctl.Password == password {
		return a.SignNewToken(ctx, username)
	}
	return nil, &openidv1.ErrInvalidUserOrPassword{}
}

func (s *Service) SignNewToken(ctx context.Context, sub string, opts ...sign.Option) (*openidv1.Token, error) {
	expiresIn := s.sessionExpiresIn()

	tokenSet := &openidv1.Token{
		TokenType: openidv1.TokenTypeBearer,
		ExpiresIn: int(expiresIn / time.Second),
	}

	opts = append(opts, sign.WithSubject(sub))

	accessToken, _, err := s.signer.Sign(
		ctx,
		append(opts, sign.WithClaim(ClaimTokenType, TokenTypeAccessToken), sign.WithExpiresIn(expiresIn))...,
	)
	if err != nil {
		return nil, err
	}
	tokenSet.AccessToken = accessToken

	refreshToken, _, err := s.signer.Sign(
		ctx, append(opts, sign.WithClaim(ClaimTokenType, TokenTypeRefreshToken), sign.WithExpiresIn(s.refreshTokenExpiresIn()))...)
	if err != nil {
		return nil, err
	}
	tokenSet.RefreshToken = refreshToken

	return tokenSet, nil
}

func (s *Service) sessionExpiresIn() time.Duration {
	return time.Duration(1) * time.Minute
}

func (s *Service) refreshTokenExpiresIn() time.Duration {
	return time.Duration(5) * time.Minute * 5
}

const (
	ClaimTokenType = "tktyp"
)

const (
	TokenTypeAccessToken  = "access_token"
	TokenTypeRefreshToken = "refresh_token"
)
