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
	log.Fatal(http.ListenAndServe(":8080", nil))
}
