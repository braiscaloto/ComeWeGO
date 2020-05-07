package jwt

import (
	"time"

	"github.com/braiscaloto/Twittor-backend/models"
	"github.com/dgrijalva/jwt-go"
)

/*GenerateJTW genera el cifrado con JWT*/
func GenerateJTW(t models.User) (string, error) {

	myKey := []byte("braiscaloto")

	payload := jwt.MapClaims{
		"Email":     t.Email,
		"Name":      t.Name,
		"Surnames":  t.Surnames,
		"Birthday":  t.Birthday,
		"Biography": t.Biography,
		"Location":  t.Location,
		"WebSite":   t.WebSite,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil

}
