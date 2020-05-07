package db

import (
	"github.com/braiscaloto/Twittor-backend/models"
	"golang.org/x/crypto/bcrypt"
)

/*TryLogin consulta si los datos introducidos nos permiten iniciar sesión*/
func TryLogin(email string, password string) (models.User, bool) {

	user, found, _ := UserAlreadyExists(email)
	if found == false {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		return user, false
	}
	return user, true

}
