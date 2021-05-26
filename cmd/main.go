package main

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/todanni/authentication/internal/server"
)

func main() {
	//// Read config
	//cfg, err := config.NewFromEnv()
	//if err != nil {
	//	log.Error(err)
	//	os.Exit(1)
	//}

	// Initialise router
	r := mux.NewRouter()

	// Create servers by passing DB connection and router
	server.NewTokenService(r)

	// Start the servers and listen
	log.Fatal(http.ListenAndServe(":8083", r))
}
