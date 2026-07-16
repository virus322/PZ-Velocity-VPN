package xray


type XrayConfig struct {

	Port int

	Protocol string

	Tag string

}


func DefaultConfig() XrayConfig {


	return XrayConfig{

		Port:443,

		Protocol:"vless",

		Tag:"pzuplink",

	}

}