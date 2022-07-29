package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teguh11/hypefast/controllers"
)

type Route struct{}

func (r *Route) Init() *mux.Router {
	handler := mux.NewRouter()

	shortenControllers := controllers.InitShortenControllers()
	// Routes consist of a path and a handler function.
	handler.HandleFunc("/generate", shortenControllers.ShortenURL).Methods(http.MethodPost)
	handler.HandleFunc("/stat/{randomString}", shortenControllers.Statistic).Methods(http.MethodGet)
	handler.HandleFunc("/{randomString}", shortenControllers.RedirectURL).Methods(http.MethodGet)

	return handler
}
