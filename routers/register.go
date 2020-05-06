package register

import (
	"net/http"

	"github.com/braiscaloto/Twittor-backend/models"
)

/*Register es la función para crear el registro de usuarios en la base de datos*/
func Register(w http.ResponseWriter, r *http.Request) {

	var t models.User

	/*err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Data error"+err.Error(), 400)
		return
	}*/

}
