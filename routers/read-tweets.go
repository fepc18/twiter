package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fepc18/twiter/bd"
)

// ReadTweets reads the tweets of a user
func ReadTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "The parameter ID is mandatory", http.StatusBadRequest)
		return
	}
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "The parameter page is mandatory", http.StatusBadRequest)
		return
	}
	page := r.URL.Query().Get("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "The page parameter must be greater than zero", http.StatusBadRequest)
		return
	}
	pag := int64(pageInt)
	response, correct := bd.ReadTweets(ID, pag)
	if correct == false {
		http.Error(w, "Error reading tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
