package api

import (
	"encoding/json"
	"net/http"

	"pz-velocity-vpn/backend/database"
)


func ListUsers(w http.ResponseWriter, r *http.Request){

	rows, err := database.DB.Query(
		"SELECT id, username, status FROM users",
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer rows.Close()


	var users []map[string]interface{}


	for rows.Next(){

		var id int
		var username string
		var status string


		rows.Scan(
			&id,
			&username,
			&status,
		)


		users = append(users,
			map[string]interface{}{
				"id": id,
				"username": username,
				"status": status,
			},
		)

	}


	w.Header().Set(
		"Content-Type",
		"application/json",
	)


	json.NewEncoder(w).Encode(users)

}