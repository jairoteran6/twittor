package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var MongoCN = conectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://user-mongo-meli:challengemeli2021@sandbox.rgvhn.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

func conectarBD() *mongo.Client  {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil{
		log.Fatal(err)
		return client
	}

	err = client.Ping(context.TODO(),nil)
	if err != nil {
		log.Fatal(err)
		return client
	}
	log.Println("Conexion Exitosa con la BD")
	return client
}

func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(),nil)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return 1
}