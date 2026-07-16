package network


import (
	"os/exec"
)



func ApplyOptimization() error {


	cmd := exec.Command(
		"sysctl",
		"-p",
	)


	return cmd.Run()

}