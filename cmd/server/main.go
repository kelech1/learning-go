package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kelech1/learning-go/internal/handlers"
)

func main() {
	r := mux.NewRouter()

	// Serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// Routes
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/analyze", handlers.AnalyzeHandler).Methods("POST")

	// Start server
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
