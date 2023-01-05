package models

import "time"

type Certificate struct {
	Name        string        `json:"name" binding:"required"`
	Issuer      string        `json:"issuer" binding:"required"`
	Instructor  string        `json:"instructor"`
	Date        time.Time     `json:"date" binding:"required"`
	Description []Description `json:"description"`
	Url         []Url         `json:"url"`
}
