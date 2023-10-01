package user

import (
	"net/http"

	"github.com/hugosjoberg/architecture/internal/api"
)

type UserHandler struct{}

func (u *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) *api.Response {
	id := "123"
	name := "John Doe"
	return api.GetUserJSON200Response(api.User{ID: &id, Name: &name})
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}
