package users

import (
	"pz-velocity-vpn/backend/database"
)


type User struct {

	ID int
	Username string
	UUID string
	VolumeLimit int
	VolumeUsed int
	ExpireDate string
	Status string

}


func Create(username string, uuid string) error {


	_, err := database.DB.Exec(
		`
		INSERT INTO users
		(username,uuid,volume_limit,status)
		VALUES(?,?,?,?)
		`,
		username,
		uuid,
		100,
		"active",
	)

	return err
}