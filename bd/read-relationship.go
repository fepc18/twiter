package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/fepc18/twiter/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ReadRelationship reads the relationship between two users
func ReadRelationship(t models.Relationship) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twitter")
	col := db.Collection("relationship")

	filter := bson.M{
		"userid":         t.UserID,
		"userrelationid": t.UserRelationID,
	}

	var result models.Relationship
	err := col.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		fmt.Println("relationship not found")
		return false, err
	}
	fmt.Println("relationship found")
	return true, nil

}
