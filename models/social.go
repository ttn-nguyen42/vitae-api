package models

type SocialMedia struct {
	Id           string `bson:"_id"`
	UserId       string `bson:"user_id"`
	Platform     string `bson:"platform" binding:"required"`
	User         string `bson:"user" binding:"required"`
	Url          Url    `bson:"url" binding:"required"`
	ColoredIcon  Url    `bson:"colored_icon"`
	MonotoneIcon Url    `bson:"monotone_icon"`
}
