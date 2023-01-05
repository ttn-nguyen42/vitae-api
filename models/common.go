package models

type Url struct {
	For string `json:"for" binding:"required"`
	Url string `json:"url" binding:"required,url"`
}

type Description struct {
	Description string `json:"description" binding:"required"`
}
