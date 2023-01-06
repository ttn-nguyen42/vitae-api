package models

import "time"

type Certificate struct {
	Id          string        `bson:"_id"`
	UserId      string        `bson:"user_id"`
	Name        string        `bson:"name" binding:"required"`
	Issuer      string        `bson:"issuer" binding:"required"`
	Instructor  string        `bson:"instructor"`
	Date        time.Time     `bson:"date" binding:"required"`
	Description []Description `bson:"description"`
	Url         []Url         `bson:"url"`
}
