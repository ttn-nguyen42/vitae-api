package models

import (
	"time"
)

type About struct {
	Id         string    `bson:"_id"`
	FirstName  string    `bson:"first_name" binding:"required"`
	LastName   string    `bson:"last_name" binding:"required"`
	Email      string    `bson:"email" binding:"required,email"`
	Birthday   time.Time `bson:"birthday"`
	City       string    `bson:"city" binding:"required"`
	Country    string    `bson:"country" binding:"required"`
	Occupation string    `bson:"occupation" binding:"required"`
}
