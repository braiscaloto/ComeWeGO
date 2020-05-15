package db

import (
	"context"
	"log"
	"time"

	"github.com/braiscaloto/Twittor-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*ReadTweet consulta los tweets de la base de datos*/
func ReadTweet(ID string, pagina int64) ([]*models.ReturnTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("Twittor")
	col := db.Collection("tweet")

	var resultados []*models.ReturnTweets

	condition := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "date", Value: -1}})
	opciones.SetSkip((pagina - 1) * 20)

	cursor, err := col.Find(ctx, condition, opciones)

	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}
	for cursor.Next(context.TODO()) {
		var registro models.ReturnTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}
	return resultados, true
}
