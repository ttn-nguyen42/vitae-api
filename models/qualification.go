package models

import "time"

type Qualification struct {
	Name        string        `json:"name" binding:"required"`
	Issuer      string        `json:"issuer" binding:"required"`
	Date        time.Time     `json:"date"`
	Score       float32       `json:"score"`
	Description []Description `json:"description"`
	Url         []Url         `json:"url"`
}
