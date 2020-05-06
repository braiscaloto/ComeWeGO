package middelwares

import (
	"net/http"

	db "github.com/braiscaloto/Twittor-backend/DB"
)

/*CheckDB es el middleware que me permite conocer el estado de la base de datos*/
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnectionDB() == 0 {
			http.Error(w, "Connection lost", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
