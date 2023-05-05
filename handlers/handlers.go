package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/fepc18/twiter/middleware"
	"github.com/fepc18/twiter/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middleware.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middleware.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/viewprofile", middleware.CheckDB(middleware.ValidateJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/modifyprofile", middleware.CheckDB(middleware.ValidateJWT(routers.ModifyProfile))).Methods("PATCH")
	router.HandleFunc("/savetweet", middleware.CheckDB(middleware.ValidateJWT(routers.SaveTweet))).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
