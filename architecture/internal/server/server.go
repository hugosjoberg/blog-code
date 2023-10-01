package server

import (
	"net/http"

	"github.com/hugosjoberg/architecture/internal/admin"
	"github.com/hugosjoberg/architecture/internal/api"
	"github.com/hugosjoberg/architecture/internal/user"
)

type handlers struct {
	user.UserHandler
	admin.AdmimHandler
}

func newHandler() *handlers {
	return &handlers{
		UserHandler:  *user.NewUserHandler(),
		AdmimHandler: *admin.NewAdmimHandler(),
	}
}

func Start() {
	apiHandler := api.Handler(newHandler())
	http.ListenAndServe(":8080", apiHandler)
}
