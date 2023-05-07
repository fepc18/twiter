package routers

import (
	"net/http"

	"github.com/fepc18/twiter/bd"
	"github.com/fepc18/twiter/models"
)

// SaveRelationship saves the relationship between users
func SaveRelationship(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id") // traerlo de la url //id del usuario
	if len(ID) < 1 {
		http.Error(w, "Must send the parameter ID", http.StatusBadRequest)
		return
	}
	var t models.Relationship
	t.UserID = IDUser     // IDUser is the global variable from token that contains the ID of the user
	t.UserRelationID = ID // ID is the ID of the user that we want to follow

	status, err := bd.InsertRelationship(t)
	if err != nil {
		http.Error(w, "Error saving the relationship "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "Relationship was not saved "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
