// Code generated by goa v3.21.1, DO NOT EDIT.
//
// hello client
//
// Command:
// $ goa gen hello/design

package hello

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "hello" service client.
type Client struct {
	SayHelloEndpoint goa.Endpoint
}

// NewClient initializes a "hello" service client given the endpoints.
func NewClient(sayHello goa.Endpoint) *Client {
	return &Client{
		SayHelloEndpoint: sayHello,
	}
}

// SayHello calls the "sayHello" endpoint of the "hello" service.
func (c *Client) SayHello(ctx context.Context, p string) (res string, err error) {
	var ires any
	ires, err = c.SayHelloEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}
