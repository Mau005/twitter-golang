package main

import (
	"log"

	"github.com/Mau005/twitter-golang/bd"
	"github.com/Mau005/twitter-golang/handlers"
)

func main() {
	if !bd.ChequeoConnection() {
		log.Fatal("Error sin conexion a la BD")
		return
	}

	handlers.Manejadores()
}
