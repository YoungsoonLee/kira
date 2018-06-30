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
	json.NewEncoder(w).Encode(map[string]string{"error": message})
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// ResponseJSON ...
// return response success
func ResponseJSON(w http.ResponseWriter, message interface{}, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if message != nil {
		json.NewEncoder(w).Encode(map[string]string{"success": message.(string)})
	}
	json.NewEncoder(w).Encode(data)
}
