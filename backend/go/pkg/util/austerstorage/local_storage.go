package austerstorage

import (
	"fmt"
	"github.com/auster-kaki/auster-mono/pkg/util/austerid"
	"io"
	"os"
	"path/filepath"
)

type ContentType string

const (
	JPEG ContentType = "image/jpeg"
	PNG  ContentType = "image/png"
)

// Save はローカルにデータを保存しパスを返す
func Save(content ContentType, data []byte) (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	var path string
	switch content {
	case JPEG:
		path = filepath.Join(pwd, "assets", "images", fmt.Sprintf("%s.jpg", austerid.Generate[string]()))
	case PNG:
		path = filepath.Join(pwd, "assets", "images", fmt.Sprintf("%s.png", austerid.Generate[string]()))
	default:
		return "", fmt.Errorf("invalid content type: %s", content)
	}

	file, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return "", err
	}

	return path, nil
}

func Get(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return b, nil
}
