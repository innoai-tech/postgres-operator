//go:generate go tool devtool gen .
package postgresoperatorclient

import (
	"context"

	"github.com/innoai-tech/infra/pkg/http/middleware"
	"github.com/innoai-tech/postgres-operator/pkg/strfmt"
	"github.com/octohelm/courier/pkg/courier"
	"github.com/octohelm/courier/pkg/courierhttp/client"
	openidv1 "github.com/octohelm/jwx/pkg/apis/openid/v1"
	"github.com/octohelm/jwx/pkg/authn"
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
		checkEndpoint := c.Endpoint
		checkEndpoint.Username = ""
		checkEndpoint.Password = ""
		if checkEndpoint.Path == "" {
			checkEndpoint.Path = "/api/postgres-operator/v1"
		}

		a := &authn.Authn{
			ClientAuth: openidv1.ClientAuth{
				ClientID:     c.Endpoint.Username,
				ClientSecret: c.Endpoint.Password,
			},
			CheckEndpoint:       checkEndpoint.String(),
			ExchangeTokenByPost: true,
			HttpTransports: []client.HttpTransport{
				middleware.NewLogRoundTripper(),
			},
		}

		c.client = &client.Client{
			Endpoint: c.Endpoint.String(),
			HttpTransports: []client.HttpTransport{
				middleware.NewLogRoundTripper(),
				a.AsHttpTransport(),
			},
		}
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
