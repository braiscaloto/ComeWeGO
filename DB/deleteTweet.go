package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*DeleteTweet borra un tweet*/
func DeleteTweet(ID string, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	db := MongoCN.Database("Twittor")
	col := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)
	condition := bson.M{
		"_id":    objID,
		"userid": userID,
	}

	_, err := col.DeleteOne(ctx, condition)

	return err
}
