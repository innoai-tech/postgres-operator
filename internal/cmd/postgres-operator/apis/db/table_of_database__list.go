package db

import (
	"context"

	databasev1 "github.com/innoai-tech/postgres-operator/pkg/apis/database/v1"
	"github.com/innoai-tech/postgres-operator/pkg/pgctl"
	"github.com/octohelm/courier/pkg/courierhttp"
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
