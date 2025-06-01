package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) WriteJson(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}