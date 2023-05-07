package routers

import (
	"encoding/json"
	"net/http"

	"github.com/fepc18/twiter/bd"
	"github.com/fepc18/twiter/models"
)

// TerminateRelationship deletes the relationship between users
func TerminateRelationship(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id") // traerlo de la url //id del usuario al que seguiamos
	if len(ID) < 1 {
		http.Error(w, "Must send the parameter ID", http.StatusBadRequest)
		return
	}
	var t models.Relationship
	t.UserID = IDUser     // IDUser is the global variable from token that contains the ID of the user
	t.UserRelationID = ID // ID is the ID of the user that we want to follow

	status, err := bd.TerminateRelationship(t)
	if err != nil {
		http.Error(w, "Error terminating the relationship "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "Relationship was not terminated "+err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode("relationship deleted")
	w.WriteHeader(http.StatusOK)
	//w.WriteHeader(http.StatusNoContent)
}
