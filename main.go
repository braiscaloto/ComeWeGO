package main

import (
	"log"

	db "github.com/braiscaloto/Twittor-backend/DB"
	"github.com/braiscaloto/Twittor-backend/handlers"
)

func main() {

	if db.CheckConnectionDB() == 0 {
		log.Fatal("No DB connection")
		return
	}

	handlers.Handlers()

}
