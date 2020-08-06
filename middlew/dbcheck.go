package middlew

import (
	"net/http"

	"github.com/EduardoSantos7/twitter/db"
)

/* Middleware that allow me too know the db state */
func DBCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Lost conection with the DB", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
