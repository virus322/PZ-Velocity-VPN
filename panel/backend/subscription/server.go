package subscription

import (
	"net/http"
)


func Handler(w http.ResponseWriter,r *http.Request){


	configs := []string{

		"vless://example",

	}


	result := Encode(configs)


	w.Header().Set(
		"Content-Type",
		"text/plain",
	)


	w.Write([]byte(result))

}