package telegram


type Bot struct {

	Token string

}


func New(token string) *Bot {


	return &Bot{

		Token:token,

	}

}


func (b *Bot) Start(){


	// Telegram listener will be added here

}