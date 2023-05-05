package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/fepc18/twiter/bd"
	"github.com/fepc18/twiter/models"
)

func SaveTweet(w http.ResponseWriter, r *http.Request) {
	var message models.SaveTweet //dto
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Data error "+err.Error(), http.StatusBadRequest)
		return
	}
	register := models.SaveTweet{
		UserID:  IDUser, // traerlo del token
		Message: message.Message,
		//Date:    message.Date,
		Date: time.Now(),
	}
	_, status, err := bd.InsertTweet(register)
	if err != nil {
		http.Error(w, "An error occurred while trying to insert the record. Try again "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "The record has not been inserted "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
