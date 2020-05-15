package routers

import (
	"net/http"

	db "github.com/braiscaloto/Twittor-backend/DB"
	"github.com/braiscaloto/Twittor-backend/models"
)

/*AltaRelacion realiza el registro de la relación en tre usuarios*/
func AltaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parámetro id es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UserID = IDUser
	t.UserRelacionID = ID

	status, err := db.InsertoRelacion(t)

	if err != nil {
		http.Error(w, "Error al insertar relación"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar la relación"+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
