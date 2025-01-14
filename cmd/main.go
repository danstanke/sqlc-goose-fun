package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/danstanke/sqlc-goose-fun/api"
)

// HealthCheck represents the structure of our health check response
type HealthCheck struct {
	Status string `json:"status"`
}

// serverImpl implements the generated ServerInterface
type serverImpl struct{}

// Health implements the /health endpoint
func (s *serverImpl) Health(w http.ResponseWriter, r *http.Request) {
	healthResponse(w, r)
}

// SystemHealth implements the /system/health endpoint
func (s *serverImpl) SystemHealth(w http.ResponseWriter, r *http.Request) {
	healthResponse(w, r)
}

// ApplicationHealth implements the /application/health endpoint
func (s *serverImpl) ApplicationHealth(w http.ResponseWriter, r *http.Request) {
	healthResponse(w, r)
}

func healthResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := HealthCheck{Status: "healthy"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	si := &serverImpl{}

	// Setup the server with the generated router
	router := api.NewServer(si)

	// Start the server
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
