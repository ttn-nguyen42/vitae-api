package models

import "time"

type Skill struct {
	Id          string        `bson:"_id"`
	UserId      string        `bson:"user_id"`
	Type        string        `bson:"type" binding:"required"`
	Name        string        `bson:"name" binding:"required"`
	Proficiency float32       `bson:"proficiency"`
	Description []Description `bson:"description"`
	Since       time.Time     `bson:"since" binding:"required"`
}
