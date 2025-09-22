//go:generate go tool devtool gen .
package postgresoperatorclient

import (
	"context"

	"github.com/innoai-tech/postgres-operator/pkg/strfmt"
	"github.com/octohelm/courier/pkg/courier"
	"github.com/octohelm/courier/pkg/courierhttp/client"
)

// Client
// +gengo:client:openapi=http://0.0.0.0:8080/api/postgres-operator
// +gengo:client:typegen-policy=GoVendorAll
// +gengo:injectable:provider
type Client struct {
	Endpoint strfmt.Endpoint `flag:",omitzero"`

	client *client.Client
}

func (c *Client) Disabled(ctx context.Context) bool {
	return c.Endpoint.IsZero()
}

func (c *Client) beforeInit(ctx context.Context) error {
	if c.client == nil {
		cc := &client.Client{
			Endpoint: c.Endpoint.String(),
		}
		c.client = cc
	}

	return nil
}

func (c *Client) Do(ctx context.Context, req any, metas ...courier.Metadata) courier.Result {
	return c.client.Do(ctx, req, metas...)
}

type Response[Data any] interface {
	ResponseData() *Data
}

type Resp[Data any] struct{}

func (Resp[Data]) ResponseData() *Data {
	return new(Data)
}

func DoWith[Data any, Op Response[Data]](
	ctx context.Context,
	c courier.Client,
	build func(req *Op),
) (*Data, error) {
	req := new(Op)
	build(req)

	resp := new(Data)
	if _, ok := any(resp).(*courier.NoContent); ok {
		_, err := c.Do(ctx, req).Into(nil)
		return resp, err
	}

	_, err := c.Do(ctx, req).Into(resp)
	return resp, err
}
