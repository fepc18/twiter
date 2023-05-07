package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fepc18/twiter/bd"
)

// GetTweetsFollowers reads the tweets from the followers with pagination
func GetTweetsFollowers(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "You must send the page parameter", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page")) // Atoi converts a string to an integer
	if err != nil {
		http.Error(w, "You must send the page parameter as an integer greater than 0", http.StatusBadRequest)
		return
	}
	response, correct := bd.ReadTweetsFollowers(IDUser, int64(page))
	if !correct {
		http.Error(w, "Error reading tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
