package db

import "golang.org/x/crypto/bcrypt"

/*CodePassword es la rutina que me permite cifrar la contraseña*/
func CodePassword(pass string) (string, error) {

	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err

}
