package routers

import (
	"encoding/json"
	"net/http"

	"github.com/fepc18/twiter/bd"
)

// DeleteTweet borra un tweet determinado
func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id") // traerlo de la url //id del tweet
	if len(ID) < 1 {
		http.Error(w, "Must send the parameter ID", http.StatusBadRequest)
		return
	}

	err := bd.DeleteTweet(ID, IDUser)
	if err != nil {
		http.Error(w, "An error occurred while trying to delete the record. Try again "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Tweet deleted successfully")
}
