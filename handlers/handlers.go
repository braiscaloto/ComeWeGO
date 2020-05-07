package handlers

import (
	"log"
	"net/http"
	"os"

	middelwares "github.com/braiscaloto/Twittor-backend/middlewares"
	"github.com/braiscaloto/Twittor-backend/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Handlers seteo mi puerto, el handler y pongo al servidor escuchando*/
func Handlers() {

	router := mux.NewRouter()

	router.HandleFunc("/register", middelwares.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middelwares.CheckDB(routers.Login)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
