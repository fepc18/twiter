package middleware

import (
	"net/http"

	"github.com/fepc18/twiter/bd"
)

// CheckDB is the middleware that allows to know the state of the database
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "DataBase Connection has been lose", 500)
			return
		}
		next.ServeHTTP(w, r)
	}

}
