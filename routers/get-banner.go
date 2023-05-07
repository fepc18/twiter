package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/fepc18/twiter/bd"
)

// GetBanner sends the banner to the http response
func GetBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id") // traerlo de la url //id del usuario
	if len(ID) < 1 {
		http.Error(w, "Must send the parameter ID", http.StatusBadRequest)
		return
	}

	profile, err := bd.FindProfile(ID)
	if err != nil {
		http.Error(w, "User not found "+err.Error(), http.StatusBadRequest)
		return
	}
	openfile, err := os.Open("uploads/banners/" + profile.Banner)

	if err != nil {
		http.Error(w, "Image not found "+err.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(w, openfile) // copy the image to the http response
	if err != nil {
		http.Error(w, "Error copying the image "+err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-type", "image/jpeg")
	w.WriteHeader(http.StatusOK)

}
