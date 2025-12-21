package students

import "net/http"

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Students Rest API"))
	}
}
