package http

import (
	"cmp"
	"context"
	"net/http"
	"os"

	appRPC "github.com/auster-kaki/auster-mono/pkg/app/rpc"
)

type Client struct {
	client *http.Client
	host   string
}

func NewClient() *Client {
	return &Client{
		client: http.DefaultClient,
		host:   cmp.Or(os.Getenv(""), "localhost"),
	}
}

func (c *Client) Diary() appRPC.Diary { return c }

func (c *Client) Create(ctx context.Context) error {
	return nil
}
