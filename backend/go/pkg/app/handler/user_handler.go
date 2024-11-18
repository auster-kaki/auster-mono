package handler

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/auster-kaki/auster-mono/pkg/app/presenter/request"
	"github.com/auster-kaki/auster-mono/pkg/app/presenter/response"
	"github.com/auster-kaki/auster-mono/pkg/app/usecase"
	"github.com/auster-kaki/auster-mono/pkg/entity"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(u *usecase.UserUseCase) []Input {
	h := &UserHandler{userUseCase: u}
	return []Input{
		{method: http.MethodGet, path: "/users", handler: h.GetUsers},
		{method: http.MethodGet, path: "/users/{id}", handler: h.GetUser},
		{method: http.MethodPost, path: "/users", handler: h.CreateUser},
		{method: http.MethodPatch, path: "/users/{id}", handler: h.UpdateUser},
		{method: http.MethodGet, path: "/hobbies", handler: h.GetHobbies},
	}
}

func (h *UserHandler) GetHobbies(w http.ResponseWriter, r *http.Request) {
	out, err := h.userUseCase.GetHobbies(r.Context())
	if err != nil {
		response.InternalError(w, err)
		return
	}
	response.OK(w, out)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	out, err := h.userUseCase.GetUsers(r.Context())
	if err != nil {
		response.InternalError(w, err)
		return
	}
	response.OK(w, out)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	out, err := h.userUseCase.GetUser(r.Context(), entity.UserID(id))
	if err != nil {
		response.InternalError(w, err)
		return
	}

	// 画像ファイルとjsonを返す処理
	w.Header().Set("Content-Type", "multipart/mixed; boundary=boundary")

	// マルチパートレスポンスを生成
	multipartWriter := multipart.NewWriter(w)
	if err := multipartWriter.SetBoundary("boundary123"); err != nil {
		response.InternalError(w, err)
		return
	}

	// ユーザー情報をjsonで返す
	userJSON, err := json.Marshal(out.User)
	if err != nil {
		response.InternalError(w, err)
		return
	}

	// ユーザー情報をマルチパートレスポンスに書き込む
	userPart, err := multipartWriter.CreatePart(map[string][]string{
		"Content-Type": {"application/json"},
	})
	if err != nil {
		response.InternalError(w, err)
		return
	}
	if _, err := userPart.Write(userJSON); err != nil {
		response.InternalError(w, err)
		return
	}

	// 画像をマルチパートレスポンスに書き込む
	imagePart, err := multipartWriter.CreatePart(map[string][]string{
		"Content-Type": {"image/jpeg"},
	})
	if err != nil {
		response.InternalError(w, err)
		return
	}

	if _, err := imagePart.Write(out.Photo); err != nil {
		response.InternalError(w, err)
		return
	}

	// マルチパートレスポンスを閉じる
	if err := multipartWriter.Close(); err != nil {
		response.InternalError(w, err)
		return
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req request.User
	if err := request.Decode(r, &req); err != nil {
		response.BadRequest(w, err)
		return
	}

	// 画像ファイルを []byte で受け取る処理
	file, _, err := r.FormFile("image")
	if err != nil {
		response.InternalError(w, err)
		return
	}
	defer file.Close()

	photo, err := io.ReadAll(file)
	if err != nil {
		response.InternalError(w, err)
		return
	}

	if err := h.userUseCase.CreateUser(r.Context(), &usecase.UserInput{
		Name:   req.Name,
		Gender: req.Gender,
		Age:    req.Age,
		Photo:  photo,
	}); err != nil {
		response.InternalError(w, err)
		return
	}
	response.Created(w, nil)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var (
		id  = r.PathValue("id")
		req request.User
	)
	if err := request.Decode(r, &req); err != nil {
		response.InternalError(w, err)
		return
	}

	// 画像ファイルを []byte で受け取る処理
	file, _, err := r.FormFile("image")
	if err != nil {
		response.InternalError(w, err)
		return
	}
	defer file.Close()

	photo, err := io.ReadAll(file)
	if err != nil {
		response.InternalError(w, err)
		return
	}

	if err := h.userUseCase.UpdateUser(r.Context(), &usecase.UserInput{
		ID:     entity.UserID(id),
		Name:   req.Name,
		Gender: req.Gender,
		Age:    req.Age,
		Photo:  photo,
	}); err != nil {
		response.InternalError(w, err)
		return
	}
	response.OK(w, nil)
}
