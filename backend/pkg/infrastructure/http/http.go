package http

import (
	"context"
	"net/http"

	appRPC "github.com/auster-kaki/auster-mono/pkg/app/rpc"
)

type Client struct {
	client *http.Client
}

func NewClient() *Client {
	return &Client{client: http.DefaultClient}
}

func (c *Client) Diary() appRPC.Diary { return c }

func (c *Client) Create(ctx context.Context) error {
	return nil
}
