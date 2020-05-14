package routers

import (
	"encoding/json"
	"net/http"

	db "github.com/braiscaloto/Twittor-backend/DB"
)

/*ViewProfile permite extraer los valores del Perfil*/
func ViewProfile(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar un parámetro ID", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar encontar el registro"+err.Error(), 400)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
