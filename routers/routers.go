package routers

import (
	"encoding/json"
	"net/http"

	"github.com/fepc18/twiter/bd"
	"github.com/fepc18/twiter/models"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error in the received data "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Password must be at least 6 characters", http.StatusBadRequest)
		return
	}
	_, encontrado, _ := bd.CheckUserExist(t.Email)
	if encontrado == true {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	_, status, err := bd.InsertRegister(t)
	if err != nil {
		http.Error(w, "An error occurred while trying to register the user "+err.Error(), http.StatusInternalServerError)
		return
	}
	if status == false {
		http.Error(w, "An error occurred while trying to register the user ", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
