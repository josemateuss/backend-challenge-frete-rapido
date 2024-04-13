package internal

import (
	"encoding/json"
	"net/http"
)

type Response struct{}

func (r Response) Write(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (r Response) WriteError(w http.ResponseWriter, status int, message string) {
	r.Write(w, status, map[string]string{"error": message})
}
