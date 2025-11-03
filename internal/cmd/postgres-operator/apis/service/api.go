// +gengo:operator:register=R
//
//go:generate go tool devtool gen .
package service

import (
	"github.com/octohelm/courier/pkg/courier"
	"github.com/octohelm/courier/pkg/courierhttp"

	openidoperator "github.com/innoai-tech/postgres-operator/internal/cmd/postgres-operator/apis/openid/operator"
)

var R = courier.NewRouter(
	courierhttp.Group("/service"),
	&openidoperator.ValidAccount{},
)
