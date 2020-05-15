package db

import (
	"context"
	"fmt"
	"time"

	"github.com/braiscaloto/Twittor-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ConsultaRelacion nos muestra las relaciones existentes entre usuarios*/
func ConsultaRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	db := MongoCN.Database("Twittor")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UserID,
		"usuariorelacionid": t.UserRelacionID,
	}

	var resultado models.Relacion

	fmt.Println(resultado)
	err := col.FindOne(ctx, condicion).Decode(&resultado)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}
