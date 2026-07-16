package telegram


type Config struct {

	Token string

	Admin string

}


func DefaultConfig() Config {


	return Config{

		Token:"BOT_TOKEN",

		Admin:"@pouyadev",

	}

}