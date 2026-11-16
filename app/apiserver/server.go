package apiserver

import (
	"io"
	"net/http"
)

type server struct {
	router *http.ServeMux
}

func NewServer() *server {
	s := &server{
		router: http.NewServeMux(),
	}

	s.configurate()

	return s
}

func (s *server) configurate() {
	s.router.HandleFunc("/", s.handleStart())
}

func (s *server) handleStart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}
}
