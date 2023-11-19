package admin

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/hugosjoberg/architecture/internal/api"
)

type AdminHandler struct{}

func (u *AdminHandler) PostAdminUser(w http.ResponseWriter, r *http.Request) *api.Response {
	var req api.PostAdminUserJSONRequestBody
	if err := render.Bind(r, &req); err != nil {
		errorCode := int32(http.StatusBadRequest)
		errorMessage := "Bad request"
		return api.PostAdminUserJSONDefaultResponse(api.Error{Code: &errorCode, Message: &errorMessage})
	}
	id := "123"
	return api.PostAdminUserJSON200Response(api.User{ID: &id, Name: req.Name})
}

func NewAdminHandler() *AdminHandler {
	return &AdminHandler{}
}
