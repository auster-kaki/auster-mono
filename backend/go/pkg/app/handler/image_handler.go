package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/auster-kaki/auster-mono/pkg/app/presenter/response"
)

func NewImageHandler() []Input {
	return []Input{
		{
			method:  http.MethodGet,
			path:    "/images/{path...}",
			handler: getImage,
		},
	}
}

func getImage(w http.ResponseWriter, r *http.Request) {
	// ファイルの存在確認
	filename := r.PathValue("path")
	fmt.Println(filename)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		response.HandleError(r.Context(), w, fmt.Errorf("image not found path:%s. %w", filename, err))
		return
	}

	// Content-Typeの設定
	ext := filepath.Ext(filename)
	switch ext {
	case ".jpg", ".jpeg":
		w.Header().Set("Content-Type", "image/jpeg")
	case ".png":
		w.Header().Set("Content-Type", "image/png")
	case ".gif":
		w.Header().Set("Content-Type", "image/gif")
	default:
		http.Error(w, "サポートされていない画像形式です", http.StatusBadRequest)
		return
	}

	// ファイルを送信
	http.ServeFile(w, r, filename)
}
