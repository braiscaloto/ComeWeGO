package main

import (
	"log"

	db "github.com/braiscaloto/Twittor/DB"
	"github.com/braiscaloto/Twittor/handlers"
)

func main() {

	if db.CheckConnectionDB() == 0 {
		log.Fatal("No DB connection")
		return
	}

	handlers.Handlers()

}
