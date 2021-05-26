package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/todanni/authentication/pkg/token"
)

func NewTokenService(router *mux.Router) token.Service {
	server := &tokenService{
		router: router,
	}
	server.routes()
	return server
}

type tokenService struct {
	router *mux.Router
}

func (s *tokenService) Issue(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (s *tokenService) Verify(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
