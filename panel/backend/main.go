package main

import (
	"encoding/json"
	"net/http"
)

type Status struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Status  string `json:"status"`
}

func main() {

	http.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {

		response := Status{
			Name:    "PZ Velocity VPN",
			Version: "0.1.0",
			Status:  "running",
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(response)
	})


	http.ListenAndServe(":8080", nil)
}