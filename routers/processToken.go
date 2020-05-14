package routers

import (
	"errors"
	"strings"

	db "github.com/braiscaloto/Twittor-backend/DB"
	"github.com/braiscaloto/Twittor-backend/models"
	"github.com/dgrijalva/jwt-go"
)

/*Email valor del email usado en los EndPoints*/
var Email string

/*IDUser valor del id usado en los EndPoints*/
var IDUser string

/*ProcessToken proceso para extraer valores del token*/
func ProcessToken(tk string) (*models.Claim, bool, string, error) {

	myKey := []byte("braiscaloto")
	claims := &models.Claim{}

	//Separamos el Bearer del resto del token
	splitToken := strings.Split(tk, "Bearer")
	//Comprobamos si nos llega el token en el formato que pedimos, el Bearer y el resto separados
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid format token")
	}

	//Nos quedamos con el token sin la parte del Bearer
	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		_, locatedUser, _ := db.UserAlreadyExists(claims.Email)
		if locatedUser == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, locatedUser, IDUser, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}
	return claims, false, string(""), err

}
