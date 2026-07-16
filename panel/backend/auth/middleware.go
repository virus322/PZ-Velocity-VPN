package auth


import (
	"net/http"
)


func Protected(
	next http.HandlerFunc,
	token string,
) http.HandlerFunc {


	return func(
		w http.ResponseWriter,
		r *http.Request,
	){


		auth := r.Header.Get("Authorization")


		if auth != token {

			http.Error(
				w,
				"Unauthorized",
				401,
			)

			return

		}


		next(w,r)

	}

}