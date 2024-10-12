package server

import "net/http"

func (s *Server) RegisterRoutes() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("GET /", s.HelloWorldHandler)
	return r
}

func (s Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
