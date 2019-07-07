package main

import (
	"encoding/json"
	"net/http"
)

type Health struct {
	Status string `json:"Status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Health{Status: "healthy"})
}