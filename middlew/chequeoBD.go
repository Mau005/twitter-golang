package middlew

import (
	"net/http"

	"github.com/Mau005/twitter-golang/bd"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if !bd.ChequeoConnection() {
			http.Error(w, "Conexion perdida con la BD", 500)
		}
	}
}
