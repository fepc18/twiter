package bd

import (
	"context"
	"time"

	"github.com/fepc18/twiter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReadTweets reads the tweets of a user
func ReadTweets(ID string, page int64) ([]*models.ReadTweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")
	var results []*models.ReadTweet
	condition := bson.M{
		"userid": ID,
	}
	options := options.Find()
	options.SetLimit(20)
	options.SetSort(bson.D{{Key: "date", Value: -1}})
	options.SetSkip((page - 1) * 20)
	cursor, err := col.Find(ctx, condition, options)
	if err != nil {
		return results, false
	}
	for cursor.Next(context.Background()) {
		var item models.ReadTweet
		err := cursor.Decode(&item)
		if err != nil {
			return results, false
		}
		results = append(results, &item)
	}
	return results, true
}
