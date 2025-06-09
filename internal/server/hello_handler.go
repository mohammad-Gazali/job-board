package server

import "net/http"

// @Summary Get Hello
// @Description say hello
// @Tags examples
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /hello [get]
func (s *Server) helloWorldHandler(w http.ResponseWriter, _ *http.Request) {
	s.WriteJson(w, map[string]string{
		"message": "Hello World!",
	})
}