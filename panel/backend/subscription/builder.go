package subscription

import (
	"encoding/base64"
)


func Encode(configs []string) string {


	data := ""

	for _,c := range configs {

		data += c + "\n"

	}


	return base64.StdEncoding.EncodeToString(
		[]byte(data),
	)

}