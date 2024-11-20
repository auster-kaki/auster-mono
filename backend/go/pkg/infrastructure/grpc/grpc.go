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

func (c *Client) CreateImage(ctx context.Context, input appRPC.CreateImageInput) (appRPC.CreateImageOutput, error) {
	return appRPC.CreateImageOutput{}, nil
}

func (c *Client) GetStatus(ctx context.Context, input appRPC.GetStatusInput) (appRPC.GetStatusOutput, error) {
	return appRPC.GetStatusOutput{}, nil
}

func (c *Client) GetImagePath(ctx context.Context, input appRPC.GetImagePathInput) (appRPC.GetImagePathOutput, error) {
	return appRPC.GetImagePathOutput{}, nil
}
