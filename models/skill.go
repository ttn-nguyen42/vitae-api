package models

import "time"

type Skill struct {
	Type        string        `json:"type" binding:"required"`
	Name        string        `json:"name" binding:"required"`
	Proficiency float32       `json:"proficiency"`
	Description []Description `json:"description"`
	Since       time.Time     `json:"since" binding:"required"`
}
