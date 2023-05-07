package routers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/fepc18/twiter/bd"
	"github.com/fepc18/twiter/models"
)

// UploadBanner allows to upload the banner to the server
func UploadBanner(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Uploading banner...")
	file, handler, err := r.FormFile("banner") // banner es el nombre del campo que se envia desde el front
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/banners/" + IDUser + "." + extension // IDUser es una variable global
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error uploading the image! "+err.Error(), http.StatusBadRequest)
		return
	}

	//buf := bytes.NewBuffer(nil)
	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "Error writing the image! "+err.Error(), http.StatusBadRequest)
		return
	}
	var user models.User
	var status bool
	user.Banner = IDUser + "." + extension
	status, err = bd.ModifyRegister(user, IDUser)
	if err != nil || status == false {
		http.Error(w, "Error saving the banner in the database! "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
