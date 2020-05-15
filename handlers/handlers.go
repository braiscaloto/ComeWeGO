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
	router.HandleFunc("/tweet", middlew.CheckDB(middlew.ValidateJWT(routers.PostTweet))).Methods("POST")
	router.HandleFunc("/readTweets", middlew.CheckDB(middlew.ValidateJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/deleteTweet", middlew.CheckDB(middlew.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/uploadAvatar", middlew.CheckDB(middlew.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/getAvatar", middlew.CheckDB(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/uploadBanner", middlew.CheckDB(middlew.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/getBanner", middlew.CheckDB(routers.GetBanner)).Methods("GET")

	/*router.HandleFunc("/altaRelacion", middlew.CheckDB(middlew.ValidateJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlew.CheckDB(middlew.ValidateJWT(routers.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middlew.CheckDB(middlew.ValidateJWT(routers.ConsultaRelacion))).Methods("GET")

	router.HandleFunc("/userList", middlew.CheckDB(middlew.ValidateJWT(routers.UserList))).Methods("GET")
	router.HandleFunc("/readTweetsFollowers", middlew.CheckDB(middlew.ValidateJWT(routers.ReadTweetsFollowers))).Methods("GET")*/
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
