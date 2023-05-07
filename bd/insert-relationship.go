package bd

import (
	"context"
	"time"

	"github.com/fepc18/twiter/models"
)

// InsertRelationship inserts a relationship in the database
func InsertRelationship(t models.Relationship) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twitter")
	col := db.Collection("relationship")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
