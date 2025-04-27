package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Structure with server info
type Server struct {
	Router *mux.Router
}

// Creating a new server instance
func NewServer() *Server {
	router := mux.NewRouter()
	return &Server{
		Router: router,
	}
}

// Start server
func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(":8080", s.Router)
}
