package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/braiscaloto/Twittor-backend/middlew"
	"github.com/braiscaloto/Twittor-backend/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Handlers seteo mi puerto, el handler y pongo al servidor escuchando*/
func Handlers() {

	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/viewprofile", middlew.CheckDB(middlew.ValidateJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/modifyProfile", middlew.CheckDB(middlew.ValidateJWT(routers.ModifyProfile))).Methods("PUT")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
