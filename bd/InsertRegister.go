package bd

import (
	"context"
	"time"

	"github.com/fepc18/twiter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertRegister is the function that allows to insert the data in the database
func InsertRegister(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() // cancela el contexto cuando la función se termine de ejecutar

	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	return result.InsertedID.(primitive.ObjectID).String(), true, nil

}
