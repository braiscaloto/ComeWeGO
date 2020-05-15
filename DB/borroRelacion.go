package db

import (
	"context"
	"time"

	"github.com/braiscaloto/Twittor-backend/models"
)

/*BorroRelacion elimina la relación entre usuarios de la base de datos*/
func BorroRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	db := MongoCN.Database("Twittor")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)

	if err != nil {
		return false, err
	}
	return true, nil
}
