package helpers

import (
	"encoding/json"
	"net/http"
)

// Response is
func Response(rw http.ResponseWriter, httpStatus int, data interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(httpStatus)
	json.NewEncoder(rw).Encode(data)
}
