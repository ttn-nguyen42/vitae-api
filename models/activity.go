package models

import "time"

type Activity struct {
	Id          string        `bson:"_id"`
	UserId      string        `bson:"user_id"`
	Role        string        `bson:"role" binding:"required"`
	Organizer   string        `bson:"organizer" binding:"required"`
	From        time.Time     `bson:"from" binding:"required"`
	To          time.Time     `bson:"to" binding:"required"`
	Description []Description `bson:"description"`
	Url         []Url         `bson:"url"`
}
