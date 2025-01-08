package httputil

import (
	"encoding/json"
	"net/http"
)

func WriteJSONWithCode(w http.ResponseWriter, resp any, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(resp)
}
