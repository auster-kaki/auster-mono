package http

import (
	"bytes"
	"cmp"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	appRPC "github.com/auster-kaki/auster-mono/pkg/app/rpc"
)

type Client struct {
	client *http.Client
	host   string
	port   string
}

func NewClient() *Client {
	return &Client{
		client: &http.Client{},
		host:   cmp.Or(os.Getenv("DIARY_HOST"), "localhost"),
		port:   cmp.Or(os.Getenv("DIARY_PORT"), "5050"),
	}
}

// URL
func (c *Client) URL() string {
	return fmt.Sprintf("http://%s:%s", c.host, c.port)
}

func (c *Client) Diary() appRPC.Diary { return c }

func (c *Client) CreateImage(ctx context.Context, input appRPC.CreateImageInput) (appRPC.CreateImageOutput, error) {

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	file, err := os.Open(input.SourcePath)
	if err != nil {
		return appRPC.CreateImageOutput{}, err
	}
	defer file.Close()

	// ファイルを追加
	part, err := writer.CreateFormFile("source_path", filepath.Base(input.SourcePath))
	if err != nil {
		return appRPC.CreateImageOutput{}, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return appRPC.CreateImageOutput{}, err
	}

	targetFile, err := os.Open(input.TargetPath)
	if err != nil {
		return appRPC.CreateImageOutput{}, err
	}
	defer targetFile.Close()

	targetPart, err := writer.CreateFormFile("target_path", filepath.Base(input.TargetPath))
	if err != nil {
		return appRPC.CreateImageOutput{}, err
	}
	_, err = io.Copy(targetPart, targetFile)
	if err != nil {
		return appRPC.CreateImageOutput{}, err
	}

	writer.Close()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.URL()+"/process", &buf)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	if err != nil {
		return appRPC.CreateImageOutput{}, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return appRPC.CreateImageOutput{}, err
	}
	defer res.Body.Close()

	var response struct {
		JobID  string `json:"job_id"`
		Status string `json:"status"`
	}

	if err != nil {
		return appRPC.CreateImageOutput{}, err
	}

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return appRPC.CreateImageOutput{}, err
	}

	return appRPC.CreateImageOutput{
		JobID: response.JobID,
	}, nil
}

func (c *Client) GetStatus(ctx context.Context, input appRPC.GetStatusInput) (appRPC.GetStatusOutput, error) {
	return appRPC.GetStatusOutput{}, nil
}

func (c *Client) GetImagePath(ctx context.Context, input appRPC.GetImagePathInput) (appRPC.GetImagePathOutput, error) {
	return appRPC.GetImagePathOutput{}, nil
}
