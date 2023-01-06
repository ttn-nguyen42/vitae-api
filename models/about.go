package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type About struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" copier:"-"`
	FirstName  string             `bson:"first_name" binding:"required"`
	LastName   string             `bson:"last_name" binding:"required"`
	Email      string             `bson:"email" binding:"required,email"`
	Birthday   time.Time          `bson:"birthday"`
	City       string             `bson:"city" binding:"required"`
	Country    string             `bson:"country" binding:"required"`
	Occupation string             `bson:"occupation" binding:"required"`
}
