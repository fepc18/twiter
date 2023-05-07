package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fepc18/twiter/bd"
)

// ReadUsers reads the users from the database with pagination
func ReadUsers(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	typeUser := r.URL.Query().Get("type")

	page := r.URL.Query().Get("page")
	if len(page) < 1 {
		http.Error(w, "Must send the parameter page", http.StatusBadRequest)
		return
	}
	pageInt := 0
	var err error
	if len(page) > 0 {
		pageInt, err = strconv.Atoi(page)
		if err != nil {
			http.Error(w, "Must send the parameter page with a value greater than zero", http.StatusBadRequest)
			return
		}
	}
	pag := int64(pageInt)
	result, status := bd.ReadUsers(IDUser, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error reading users", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
