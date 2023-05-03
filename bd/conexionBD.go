package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN is the object to connect to the database
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb://fepc18:1pXPKmJnDuIGW9th@ac-iox01qu-shard-00-00.k96w0lu.mongodb.net:27017,ac-iox01qu-shard-00-01.k96w0lu.mongodb.net:27017,ac-iox01qu-shard-00-02.k96w0lu.mongodb.net:27017/?ssl=true&replicaSet=atlas-odx1d7-shard-0&authSource=admin&retryWrites=true&w=majority")

// Funcion para conectar a la BD
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión exitosa con la BD")
	return client
}

// CheckConnection is the ping to the database
func CheckConnection() int {
	err := MongoCN.Ping(context.Background(), nil)
	if err != nil {
		return 0
	}
	return 1
}
