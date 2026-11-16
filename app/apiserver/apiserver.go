package apiserver

import "net/http"

func StartServer(cfg Config) error {

	srv := NewServer()

	return http.ListenAndServe(cfg.Addr, srv.router)
}
