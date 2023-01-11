package models

type Url struct {
	For string `bson:"for" json:"for" binding:"required"`
	Url string `bson:"url" json:"url" binding:"required,url"`
}

type Description struct {
	Description string `bson:"description" json:"description" binding:"required"`
}
