package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct{}

func (r *Route) Init() *mux.Router {
	handler := mux.NewRouter()
	// Routes consist of a path and a handler function.
	handler.HandleFunc("/", YourHandler)

	return handler
}
func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}
