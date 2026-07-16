package users

import (
	"pz-velocity-vpn/backend/database"
	"pz-velocity-vpn/backend/utils"
)


type User struct {

	ID int
	Username string
	UUID string
	Status string

}


func Create(username string) (string,error){

	uuid := utils.GenerateUUID()


	_,err := database.DB.Exec(
		`
		INSERT INTO users
		(
		username,
		uuid,
		volume_limit,
		status
		)
		VALUES(?,?,?,?)
		`,
		username,
		uuid,
		100,
		"active",
	)


	return uuid,err

}