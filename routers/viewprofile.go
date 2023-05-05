package routers

import (
	"encoding/json"
	"net/http"

	"github.com/fepc18/twiter/bd"
)

func ViewProfile(w http.ResponseWriter, r *http.Request) {
	/* vars := mux.Vars(r) // router-->r.HandleFunc("/provisions/{id}", Provisions)
	   id, ok := vars["id"]
	*/
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Must send the parameter ID", http.StatusBadRequest)
		return
	}
	profile, err := bd.FindProfile(ID)
	if err != nil {
		http.Error(w, "An error occurred while trying to find the record "+err.Error(), 400)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
