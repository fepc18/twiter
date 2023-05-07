package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/fepc18/twiter/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ReadTweetsFollowers reads the tweets from the followers
func ReadTweetsFollowers(ID string, page int64) ([]models.ResponseTweetsFollowers, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("relationship")

	skip := (page - 1) * 20 // 20 is the number of tweets per page

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"userid": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "userrelationid",
			"foreignField": "userid",
			"as":           "tweet",
		}})
	conditions = append(conditions, bson.M{"$unwind": "$tweet"}) // Unwind is used to get the tweet as a single document and not as an array
	conditions = append(conditions, bson.M{"$sort": bson.M{"tweet.date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, conditions)
	var result []models.ResponseTweetsFollowers
	err = cursor.All(ctx, &result)
	if err != nil {
		fmt.Println(err.Error())
		return result, false
	}
	return result, true
}
