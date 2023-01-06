package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Skill struct {
	Id          primitive.ObjectID        `bson:"_id"`
	UserId      primitive.ObjectID        `bson:"user_id"`
	Type        string        `bson:"type" binding:"required"`
	Name        string        `bson:"name" binding:"required"`
	Proficiency float32       `bson:"proficiency"`
	Description []Description `bson:"description"`
	Since       time.Time     `bson:"since" binding:"required"`
}
