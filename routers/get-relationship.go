package routers

import (
	"encoding/json"
	"net/http"

	"github.com/fepc18/twiter/bd"
	"github.com/fepc18/twiter/models"
)

// GetRelationship reads the relationship between two users
func GetRelationship(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id") // traerlo de la url //id del usuario al que seguiamos
	if len(ID) < 1 {
		http.Error(w, "Must send the parameter ID", http.StatusBadRequest)
		return
	}
	var t models.Relationship
	t.UserID = IDUser     // IDUser is the global variable from token that contains the ID of the user
	t.UserRelationID = ID // ID is the ID of the user that we want to follow

	var resp models.ResponseReadRelationship
	status, err := bd.ReadRelationship(t)
	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
