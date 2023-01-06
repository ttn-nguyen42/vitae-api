package models

import "time"

type Qualification struct {
	Id          string        `bson:"_id"`
	UserId      string        `bson:"user_id"`
	Name        string        `bson:"name" binding:"required"`
	Issuer      string        `bson:"issuer" binding:"required"`
	Date        time.Time     `bson:"date"`
	Score       float32       `bson:"score"`
	Description []Description `bson:"description"`
	Url         []Url         `bson:"url"`
}
