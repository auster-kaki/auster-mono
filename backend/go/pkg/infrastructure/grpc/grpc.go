package grpc

import (
	"context"

	appRPC "github.com/auster-kaki/auster-mono/pkg/app/rpc"
)

type Client struct {
}

func New() *Client {
	return &Client{}
}

func (c *Client) Diary() appRPC.Diary { return c }

func (c *Client) Create(ctx context.Context) error {
	return nil
}
