package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/braiscaloto/Twittor-backend/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Handlers seteo mi puerto, el handler y pongo al servidor escuchando*/
func Handlers() {

	router := mux.NewRouter()

	router.HandleFunc("/register", middlewares.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlewares.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/viewprofile", middlewares.CheckDB()).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
