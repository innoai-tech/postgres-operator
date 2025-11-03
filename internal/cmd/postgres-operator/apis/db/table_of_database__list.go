package db

import (
	"context"

	"github.com/octohelm/courier/pkg/courierhttp"

	databasev1 "github.com/innoai-tech/postgres-operator/pkg/apis/database/v1"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl"
)

// +gengo:injectable
type ListTableOfDatabase struct {
	courierhttp.MethodGet `path:"/databases/{databaseCode}/tables"`

	DatabaseName databasev1.DatabaseCode `name:"databaseCode" in:"path"`

	c *pgctl.Controller `inject:""`
}

func (req *ListTableOfDatabase) Output(ctx context.Context) (any, error) {
	return req.c.ListTableOfDatabase(ctx, req.DatabaseName)
}
