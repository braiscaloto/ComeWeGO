package main

import (
	"log"
	"net/http"

	"github.com/braiscaloto/Twittor-backend/bd"
	"github.com/braiscaloto/Twittor-backend/handlers"
)

func cors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.Write([]byte("Hello, World!"))
}
func main() {
	if bd.CheckConnectionDB() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
