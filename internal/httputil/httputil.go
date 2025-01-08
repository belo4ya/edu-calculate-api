package httputil

import (
	"encoding/json"
	"net/http"
)

// WriteJSONWithCode sends a JSON response with the specified HTTP status code.
func WriteJSONWithCode(w http.ResponseWriter, resp any, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(resp)
}
