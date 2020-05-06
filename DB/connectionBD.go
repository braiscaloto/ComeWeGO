package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexión a la base de datos*/
var MongoCN = ConnectBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://braiskalo:Otto3192@cluster0-lui73.mongodb.net/test?retryWrites=true&w=majority")

/*ConnectBD es la función que me permite conectar la BD.*/
func ConnectBD() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return client
	}

	log.Println("Connection made")
	return client

}

/*CheckConnectionDB es el ping a la BD*/
func CheckConnectionDB() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}

	return 1
}
