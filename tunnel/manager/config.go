package manager


type TunnelConfig struct {


	Protocol string


	MTU int


	KeepAlive int


}


func DefaultTunnel() TunnelConfig {


	return TunnelConfig{

		Protocol:"tcp",

		MTU:1500,

		KeepAlive:30,

	}

}