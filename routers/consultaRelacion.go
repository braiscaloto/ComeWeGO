package routers

import (
	"encoding/json"
	"net/http"

	db "github.com/braiscaloto/Twittor-backend/DB"
	"github.com/braiscaloto/Twittor-backend/models"
)

/*ConsultaRelacion compruebo si existe relación entre dos usuarios*/
func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relacion
	t.UsuarioID = IDUser
	t.UsuarioRelacionID = ID

	var resp models.RespuestaConsultaRelacion

	status, err := db.ConsultoRelacion(t)
	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
