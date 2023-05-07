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
	router.HandleFunc("/readtweets", middleware.CheckDB(middleware.ValidateJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/deletetweet", middleware.CheckDB(middleware.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/uploadavatar", middleware.CheckDB(middleware.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/uploadbanner", middleware.CheckDB(middleware.ValidateJWT(routers.UploadBanner))).Methods("POST")

	router.HandleFunc("/getavatar", middleware.CheckDB(routers.GetAvatar)).Methods("GET") // no necesita validar el token
	router.HandleFunc("/getbanner", middleware.CheckDB(routers.GetBanner)).Methods("GET")

	router.HandleFunc("/saverelationship", middleware.CheckDB(middleware.ValidateJWT(routers.SaveRelationship))).Methods("POST")
	router.HandleFunc("/terminaterelationship", middleware.CheckDB(middleware.ValidateJWT(routers.TerminateRelationship))).Methods("DELETE")

	router.HandleFunc("/readrelationship", middleware.CheckDB(middleware.ValidateJWT(routers.GetRelationship))).Methods("GET")
	router.HandleFunc("/listusers", middleware.CheckDB(middleware.ValidateJWT(routers.ReadUsers))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
