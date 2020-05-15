package routers

import (
	"net/http"

	db "github.com/braiscaloto/Twittor-backend/DB"
	"github.com/braiscaloto/Twittor-backend/models"
)

/*BajaRelacion elimina la relación entre usuarios*/
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relacion
	t.UsuarioID = IDUser
	t.UsuarioRelacionID = ID

	status, err := db.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Error al intentar borrar la relación "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha podido borrar la relación "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
