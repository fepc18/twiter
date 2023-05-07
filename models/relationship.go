package models

// Relationship is the model for the relationship between users

type Relationship struct {
	UserID         string `bson:"userid" json:"userId"`
	UserRelationID string `bson:"userrelationid" json:"userRelationId"`
}
