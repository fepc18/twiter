package routers

import (
	"encoding/json"
	"net/http"

	"github.com/fepc18/twiter/bd"
	"github.com/fepc18/twiter/models"
)

func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Data error "+err.Error(), http.StatusBadRequest)
		return
	}
	//ID := r.URL.Query().Get("id")// traerlo de la url
	ID := IDUser // traerlo del token
	if len(ID) < 1 {
		http.Error(w, "Must send the parameter ID", http.StatusBadRequest)
		return
	}

	var status bool
	status, err = bd.ModifyRegister(t, ID)
	if err != nil {
		http.Error(w, "An error occurred while trying to modify the record. Try again "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "The record has not been modified "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
