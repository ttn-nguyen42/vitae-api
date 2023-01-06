package models

type Url struct {
	For string `bson:"for" binding:"required"`
	Url string `bson:"url" binding:"required,url"`
}

type Description struct {
	Description string `bson:"description" binding:"required"`
}
