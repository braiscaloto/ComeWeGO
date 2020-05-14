package db

import (
	"context"
	"fmt"
	"time"

	"github.com/braiscaloto/Twittor-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*SearchProfile busca un perfil en la base de datos*/
func SearchProfile(ID string) (models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("Twittor")
	col := db.Collection("usuarios")

	var profile models.User
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)
	//omitimos que nos devuelva la contraseña
	profile.Password = ""
	if err != nil {
		fmt.Println("Register not found" + err.Error())
		return profile, err
	}
	return profile, nil

}
