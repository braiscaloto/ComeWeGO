package db

import (
	"context"
	"time"

	"github.com/braiscaloto/Twittor-backend/models"
)

/*InsertoRelacion inserta la relación entre usuarios en la base de datos*/
func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	db := MongoCN.Database("Twittor")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)

	if err != nil {
		return false, err
	}
	return true, nil
}
