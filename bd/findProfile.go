package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/fepc18/twiter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindProfile finds a profile in the database
func FindProfile(Id string) (models.User, error) {
	fmt.Println(Id)
	db := MongoCN.Database("twitter")
	col := db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var profile models.User
	objID, _ := primitive.ObjectIDFromHex(Id)
	condition := bson.M{
		"_id": objID,
	}
	err := col.FindOne(ctx, condition).Decode(&profile)

	profile.Password = ""
	if err != nil {
		fmt.Println("Profile not found " + err.Error())
		return profile, err
	}
	return profile, nil

}
