package network


import (
	"os"
)



func EnableBBR() error {


	config := `

net.core.default_qdisc=fq
net.ipv4.tcp_congestion_control=bbr

`


	return os.WriteFile(
		"/etc/sysctl.d/99-pz-network.conf",
		[]byte(config),
		0644,
	)

}