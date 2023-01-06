package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SocialMedia struct {
	Id           primitive.ObjectID `bson:"_id"`
	UserId       primitive.ObjectID `bson:"user_id"`
	Platform     string `bson:"platform" binding:"required"`
	User         string `bson:"user" binding:"required"`
	Url          Url    `bson:"url" binding:"required"`
	ColoredIcon  Url    `bson:"colored_icon"`
	MonotoneIcon Url    `bson:"monotone_icon"`
}
