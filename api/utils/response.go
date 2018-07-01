package utils

import (
	"encoding/json"
	"net/http"
)

// ResponseError ...
// return response error
func ResponseError(w http.ResponseWriter, message string, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	} else {
		json.NewEncoder(w).Encode(map[string]string{"error": message})
	}
}

// ResponseJSON ...
// return response success
func ResponseJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
