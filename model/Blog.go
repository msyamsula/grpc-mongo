package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	ID         primitive.ObjectID `bson:"_id"`
	Author_id  string             `bson:"author_id"`
	Title      string             `bson:"title"`
	Content    string             `bson:"content"`
	Created_at time.Time          `bson:"created_at"`
	Updated_at time.Time          `bson:"Updated_at"`
}
