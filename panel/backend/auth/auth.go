package auth


type Admin struct {

	Username string

	Password string

}


func DefaultAdmin() Admin {

	return Admin{

		Username:"admin",

		Password:"change-me",

	}

}