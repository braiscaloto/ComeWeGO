package main

import (
	"log"

	"github.com/braiscaloto/Twittor-backend/bd"
	"github.com/braiscaloto/Twittor-backend/handlers"
)

func main() {
	if bd.CheckConnectionDB() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
