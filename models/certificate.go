package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Certificate struct {
	Id          primitive.ObjectID        `bson:"_id"`
	UserId      primitive.ObjectID        `bson:"user_id"`
	Name        string        `bson:"name" binding:"required"`
	Issuer      string        `bson:"issuer" binding:"required"`
	Instructor  string        `bson:"instructor"`
	Date        time.Time     `bson:"date" binding:"required"`
	Description []Description `bson:"description"`
	Url         []Url         `bson:"url"`
}
