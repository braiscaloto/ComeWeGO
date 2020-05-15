package routers

import (
	"net/http"

	db "github.com/braiscaloto/Twittor-backend/DB"
	"github.com/braiscaloto/Twittor-backend/models"
)

/*BajaRelacion elimina la relación entre usuarios*/
func BajaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parámetro id es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UserID = IDUser
	t.UserRelacionID = ID

	status, err := db.BorroRelacion(t)

	if err != nil {
		http.Error(w, "Error al borrar relación"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado borrar la relación"+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
