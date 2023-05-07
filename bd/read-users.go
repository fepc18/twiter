package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/fepc18/twiter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReadUsers reads the users from the database
// typeUser: if empty, all users are read, if typeUser = "new", only new users are read
func ReadUsers(ID string, page int64, search string, typeUser string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}
	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println("error reading users " + err.Error())
		return results, false
	}
	var found, include bool
	for cur.Next(ctx) {
		var s models.User
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println("error decoding user " + err.Error())
			return results, false
		}
		var r models.Relationship
		r.UserID = ID
		r.UserRelationID = s.ID.Hex()
		include = false
		found, err = ReadRelationship(r)
		if typeUser == "new" && found == false {
			include = true // include the user in the results, new are the users that we don't follow
		}
		if typeUser == "follow" && found == true {
			include = true // include the user in the results, follow are the users that we follow
		}
		if r.UserRelationID == ID {
			include = false
		}
		if include == true {
			results = append(results, &s)
		}
	}
	err = cur.Err()
	if err != nil {
		fmt.Println("error reading users " + err.Error())
		return results, false
	}
	cur.Close(ctx)
	return results, true
}
