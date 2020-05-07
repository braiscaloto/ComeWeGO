package middlewares

import (
	"net/http"

	"github.com/braiscaloto/Twittor-backend/routers"
)

/*ValidateJWT nos permite validar el JWT que viene en la petición*/
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Token error! "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)

	}
}
