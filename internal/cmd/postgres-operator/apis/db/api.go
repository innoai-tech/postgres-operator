// +gengo:operator:register=R
//
//go:generate go tool devtool gen .
package db

import (
	openidoperator "github.com/innoai-tech/postgres-operator/internal/cmd/postgres-operator/apis/openid/operator"
	"github.com/octohelm/courier/pkg/courier"
	"github.com/octohelm/courier/pkg/courierhttp"
)

var R = courier.NewRouter(
	courierhttp.Group("/db"),
	&openidoperator.ValidAccount{},
)
