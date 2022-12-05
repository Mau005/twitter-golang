package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://mau:nEE5toKtjGINleHm@twitter.biy9zqc.mongodb.net/?retryWrites=true&w=majority")

/*ConectarBD Permite conectar a la BD*/
func ConectarBD() *mongo.Client {
	/*
		los Context es un espacio en memoria que se puede compartir cosas.
		contextos nos sirven para intercambiar la informacion entre funcion y funcion
		el contentext.TODO indica que tiene permiso de ingresar a la BD sin restricciones
	*/
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatalln(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Conexión exitosa a BD")
	return client
}

/*ChequeoConnection hace un ping a la BD*/
func ChequeoConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}
