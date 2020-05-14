package db

import (
	"context"
	"time"

	"github.com/braiscaloto/Twittor-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*UpdateRegister actualizamos nuestros datos*/
func UpdateRegister(u models.User, ID string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	db := MongoCN.Database("Twittor")
	col := db.Collection("usuarios")

	register := make(map[string]interface{})
	if len(u.Name) > 0 {
		register["name"] = u.Name
	}
	if len(u.Surnames) > 0 {
		register["surnames"] = u.Surnames
	}
	register["birthday"] = u.Birthday
	if len(u.Avatar) > 0 {
		register["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		register["banner"] = u.Banner
	}
	if len(u.Biography) > 0 {
		register["biography"] = u.Biography
	}
	if len(u.Location) > 0 {
		register["location"] = u.Location
	}
	if len(u.WebSite) > 0 {
		register["WebSite"] = u.WebSite
	}

	updtString := bson.M{
		"$Set": register,
	}
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	_, err := col.UpdateOne(ctx, filter, updtString)

	if err != nil {
		return false, err
	}
	return true, nil
}
