package register

import (
	"encoding/json"
	"net/http"

	db "github.com/braiscaloto/Twittor-backend/DB"
	"github.com/braiscaloto/Twittor-backend/models"
)

/*Register es la función para crear el registro de usuarios en la base de datos*/
func Register(w http.ResponseWriter, r *http.Request) {

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Data error"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Password is too short, it needs to have at least 6 characacters", 400)
		return
	}
	_, encontrado, _ := db.UserAlreadyExists(t.Email)

	if encontrado == true {
		http.Error(w, "Email already registered on Twittor", 400)
		return
	}

	_, status, err := db.InsertRegister(t)

	if err != nil {
		http.Error(w, "Register error"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "User registration failed to insert", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
