package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Health struct {
	Status string `json:"Status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Health{Status: "healthy"})
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/crossoff", makeHandler(crossoffHandler))
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
