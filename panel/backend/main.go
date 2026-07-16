package main

import (
	"encoding/json"
	"log"
	"net/http"

	"pz-velocity-vpn/backend/api"
	"pz-velocity-vpn/backend/database"
)

type Status struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Status  string `json:"status"`
}

func main() {

	// Initialize Database
	err := database.Init()

	if err != nil {
		log.Fatal("Database error:", err)
	}


	// API Status
	http.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {

		response := Status{
			Name:    "PZ Velocity VPN",
			Version: "0.1.0",
			Status:  "running",
		}


		w.Header().Set(
			"Content-Type",
			"application/json",
		)


		json.NewEncoder(w).Encode(response)
	})


	// Users API
	http.HandleFunc(
		"/api/users",
		api.ListUsers,
	)


	log.Println("PZ Velocity VPN API running on :1196")


	err = http.ListenAndServe(
		":1196",
		nil,
	)


	if err != nil {
		log.Fatal(err)
	}
}