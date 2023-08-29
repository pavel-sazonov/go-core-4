package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// API предоставляет интерфейс программного взаимодействия.
type API struct {
	router *mux.Router
}

// New создаёт объект API.
func New() *API {
	api := API{
		router: mux.NewRouter(),
	}

	api.endpoints()

	return &api
}

func (api *API) Router() *mux.Router {
	return api.router
}

// endpoints регистрирует конечные точки API.
func (api *API) endpoints() {
	api.router.Use(headersMiddleware)
	api.router.HandleFunc("/api/v1/docs/{search}", api.doc).Methods(http.MethodGet)
	api.router.HandleFunc("/api/v1/docs/", api.docs).Methods(http.MethodGet)
}
