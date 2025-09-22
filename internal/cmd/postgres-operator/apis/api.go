package apis

import (
	"github.com/octohelm/courier/pkg/courierhttp"

	archiveapis "github.com/innoai-tech/postgres-operator/internal/cmd/postgres-operator/apis/archive"
	serviceapis "github.com/innoai-tech/postgres-operator/internal/cmd/postgres-operator/apis/service"
	statusapis "github.com/innoai-tech/postgres-operator/internal/cmd/postgres-operator/apis/status"
)

var R = courierhttp.GroupRouter("/api/postgres-operator").With(
	courierhttp.GroupRouter("/v1").With(
		statusapis.R,
		archiveapis.R,
		serviceapis.R,
	),
)
