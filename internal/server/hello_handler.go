package server

import "net/http"

func (s *Server) helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	s.WriteJson(w, map[string]string{
		"message": "Hello World!",
	})
}