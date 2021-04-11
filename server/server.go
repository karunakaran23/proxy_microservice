package server

import (
	"net/http"
	"proxy_microservice/handler"
	"proxy_microservice/middleware"
)

type Server struct {
	Router *http.ServeMux
}

func (s *Server) InitRoute(h *handler.Handler) {
	s.Router.HandleFunc("/ProxyUser", middleware.JSONandCORS(h.ProxyUser))

}
