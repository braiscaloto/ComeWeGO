package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	db "github.com/braiscaloto/Twittor-backend/DB"
	"github.com/braiscaloto/Twittor-backend/models"
)

/*UploadAvatar subimos avatar a base de datos*/
func UploadAvatar(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]

	/*Establezco como nombrar el archivo que guardo*/
	var archivo string = "uploads/avatars/" + IDUser + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, "Error al subir la imagen"+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "Error al copiar la imagen"+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Avatar = IDUser + "." + extension
	status, err = db.UpdateRegister(user, IDUser)

	if err != nil || status == false {
		http.Error(w, "Error al guardar el avatar en la base de datos"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
