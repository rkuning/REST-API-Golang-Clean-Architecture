package util

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, code int, data any) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
