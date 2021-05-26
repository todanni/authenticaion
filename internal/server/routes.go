package server

import "net/http"

func (s *tokenService) routes() {
	// Issue a token
	s.router.HandleFunc("/issue", s.Issue).Methods(http.MethodGet)

	// Verify a token
	s.router.HandleFunc("/verify", s.Verify).Methods(http.MethodPost)
}
