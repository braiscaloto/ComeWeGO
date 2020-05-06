package middelwares

import (
	"net/http"

	db "github.com/braiscaloto/Twittor/DB"
)

func checkDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnectionDB() == 0 {
			http.Error(w, "Connection lost", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
