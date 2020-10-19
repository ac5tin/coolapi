package api

import (
	"api/tools"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler - routing
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling URL: ", r.URL.Path)
	router := mux.NewRouter()
	apirouter := router.PathPrefix("/api").Subrouter()

	// API Routes
	apirouter.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("pong")) }).Methods("GET")

	apirouter.HandleFunc("/tools/ip", tools.IPAddr).Methods("GET", "OPTIONS")
	// SERVE HTTP
	router.ServeHTTP(w, r)
}
