package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores seteos mi puerto el handler y pongo a escuchar el servidor*/
func Manejadores() {
	router := mux.NewRouter()
	PORT := os.Getenv("PORT") //Buscamos un purto en nuestro entorno global si no esta le asignamos
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router) //hor son
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
