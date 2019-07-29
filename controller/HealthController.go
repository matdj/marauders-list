package controller

import (
	"encoding/json"
	"net/http"
)

type Health struct {
	Status string `json:"Status"`
}

type HealthController struct {

}

func (c *HealthController) HealthHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Health{Status: "healthy"})
}
