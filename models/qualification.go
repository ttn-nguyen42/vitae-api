package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Qualification struct {
	Id          primitive.ObjectID        `bson:"_id"`
	UserId      primitive.ObjectID        `bson:"user_id"`
	Name        string        `bson:"name" binding:"required"`
	Issuer      string        `bson:"issuer" binding:"required"`
	Date        time.Time     `bson:"date"`
	Score       float32       `bson:"score"`
	Description []Description `bson:"description"`
	Url         []Url         `bson:"url"`
}
