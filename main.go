package main

import (
	"log"

	bd "github.com/braiscaloto/Twittor/DB"
	"github.com/braiscaloto/Twittor/handlers"
)

func main() {

	if bd.CheckConnectionDB() == 0 {
		log.Fatal("No DB connection")
		return
	}

	handlers.Handlers()

}
