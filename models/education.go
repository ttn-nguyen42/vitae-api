package models

import "time"

type Education struct {
	School         string             `json:"school" binding:"required"`
	Programme      string             `json:"programme" binding:"required"`
	Specialization string             `json:"specialization"`
	From           time.Time          `json:"from" binding:"required"`
	To             time.Time          `json:"to" binding:"required"`
	AverageScore   float32            `json:"averageScore" binding:"alphanum,gte=0,lte=10"`
	Description    []Description      `json:"description"`
	Courses        []EducationCourses `json:"courses"`
	Url            []Url              `json:"url"`
	City           string             `json:"address" binding:"required"`
	Country        string             `json:"country" binding:"required"`
}

type EducationCourses struct {
	Name  string `json:"name" binding:"required"`
	Grade string `json:"grade"`
}
