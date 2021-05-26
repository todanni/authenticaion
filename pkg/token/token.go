package token

import "net/http"

type Token struct {
}

type Service interface {
	// IssueToken will return a valid JWT
	Issue(w http.ResponseWriter, r *http.Request)

	// VerifyToken will verify the token is
	Verify(w http.ResponseWriter, r *http.Request)
}
