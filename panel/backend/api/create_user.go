package api

import (
	"encoding/json"
	"net/http"

	"pz-velocity-vpn/backend/users"
)


type CreateRequest struct {

	Username string `json:"username"`

}


func CreateUser(w http.ResponseWriter,r *http.Request){


	var req CreateRequest


	json.NewDecoder(r.Body).Decode(&req)


	uuid,err:=users.Create(req.Username)


	if err != nil {

		http.Error(w,err.Error(),500)

		return
	}


	json.NewEncoder(w).Encode(
		map[string]string{

			"username":req.Username,
			"uuid":uuid,
			"status":"created",

		},
	)

}