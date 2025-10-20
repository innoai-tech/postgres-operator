package apis

import (
	openidoperator "github.com/innoai-tech/postgres-operator/internal/cmd/postgres-operator/apis/openid/operator"
	"github.com/octohelm/courier/pkg/courier"
	"github.com/octohelm/courier/pkg/courierhttp"

	archiveapis "github.com/innoai-tech/postgres-operator/internal/cmd/postgres-operator/apis/archive"
	dbapis "github.com/innoai-tech/postgres-operator/internal/cmd/postgres-operator/apis/db"
	openidapis "github.com/innoai-tech/postgres-operator/internal/cmd/postgres-operator/apis/openid"
	serviceapis "github.com/innoai-tech/postgres-operator/internal/cmd/postgres-operator/apis/service"
	statusapis "github.com/innoai-tech/postgres-operator/internal/cmd/postgres-operator/apis/status"
)

var R = courierhttp.GroupRouter("/api/postgres-operator").With(
	courierhttp.GroupRouter("/v1").With(
		statusapis.R,
		archiveapis.R,
		serviceapis.R,
		dbapis.R,
		openidapis.R,

		courier.NewRouter(&openidoperator.BaseURL{}),
	),
)
