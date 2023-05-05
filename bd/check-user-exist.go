package bd

import (
	"context"
	"time"

	"github.com/fepc18/twiter/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckUserExist is the function that allows to know if the user already exists in the database
func CheckUserExist(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() // cancela el contexto cuando la función se termine de ejecutar
	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	condition := bson.M{"email": email}
	var resultado models.User
	err := col.FindOne(ctx, condition).Decode(&resultado) // decode json the result in the variable resultado

	if err != nil {
		return resultado, false, ""
	}

	ID := resultado.ID.Hex()
	return resultado, true, ID
}
