package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Activity struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" copier:"-"`
	UserId      primitive.ObjectID `bson:"user_id"`
	Role        string             `bson:"role" binding:"required"`
	Organizer   string             `bson:"organizer" binding:"required"`
	From        time.Time          `bson:"from" binding:"required"`
	To          time.Time          `bson:"to" binding:"required"`
	Description []Description      `bson:"description"`
	Url         []Url              `bson:"url"`
}
