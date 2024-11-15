package handler

import (
	"net/http"

	"github.com/auster-kaki/auster-mono/pkg/app/usecase"
)

type VendorHandler struct {
	vendorUseCase usecase.VendorUseCase
}

func NewVendorHandler(u *usecase.VendorUseCase) map[string]http.HandlerFunc {
	//h := &VendorHandler{vendorUseCase: u}
	return map[string]http.HandlerFunc{}
}
