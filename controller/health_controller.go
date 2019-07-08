package controller

import (
	"encoding/json"
	"net/http"
)

type Health struct {
	Status string `json:"Status"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Health{Status: "healthy"})
}
