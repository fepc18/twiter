package routers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/fepc18/twiter/bd"
	"github.com/fepc18/twiter/jwt"
	"github.com/fepc18/twiter/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Login")

	w.Header().Add("content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Email and/or password invalid "+err.Error(), http.StatusBadRequest)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}
	userExist, existe := bd.Login(t.Email, t.Password)
	if existe == false {
		http.Error(w, "Email and/or password invalid ", http.StatusBadRequest)
		return
	}

	jwtKey, err := jwt.GenerateJWT(userExist)
	if err != nil {
		http.Error(w, "An error occurred while trying to generate the token corresponding to the user "+err.Error(), http.StatusBadRequest)
		return
	}

	resp := models.ResponseLogin{
		Token: jwtKey,
	}
	//resp := userExist

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	//set cookie for 24 hours from backend
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
