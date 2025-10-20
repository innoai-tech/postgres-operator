package db

import (
	"context"

	databasev1 "github.com/innoai-tech/postgres-operator/pkg/apis/database/v1"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl"
	"github.com/octohelm/courier/pkg/courierhttp"
)

// +gengo:injectable
type QueryDatabase struct {
	courierhttp.MethodPost `path:"/databases/{databaseName}/query"`

	DatabaseName databasev1.DatabaseCode `name:"databaseName" in:"path"`

	Body DatabaseQueryRequest `in:"body"`

	c *pgctl.Controller `inject:""`
}

func (req *QueryDatabase) Output(ctx context.Context) (any, error) {
	return req.c.QueryDatabaseResult(ctx, req.DatabaseName, req.Body.Sql)
}

type DatabaseQueryRequest struct {
	Sql string `json:"sql"`
}
