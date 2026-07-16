package xray

import (
	"os"
)

const ConfigPath = "core/xray/config.json"


func SaveConfig(data string) error {

	err := os.WriteFile(
		ConfigPath,
		[]byte(data),
		0644,
	)

	return err
}


func ReadConfig() (string,error){

	data,err := os.ReadFile(ConfigPath)

	if err != nil {
		return "",err
	}

	return string(data),nil

}