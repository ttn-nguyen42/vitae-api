package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Education struct {
	Id             primitive.ObjectID `bson:"_id,omitempty" copier:"-"`
	UserId         primitive.ObjectID `bson:"user_id"`
	School         string             `bson:"school" binding:"required"`
	Programme      string             `bson:"programme" binding:"required"`
	Specialization string             `bson:"specialization"`
	From           time.Time          `bson:"from" binding:"required"`
	To             time.Time          `bson:"to" binding:"required"`
	AverageScore   float32            `bson:"average_score" binding:"alphanum,gte=0,lte=10"`
	Description    []Description      `bson:"description"`
	Courses        []EducationCourses `bson:"courses"`
	Url            []Url              `bson:"url"`
	City           string             `bson:"address" binding:"required"`
	Country        string             `bson:"country" binding:"required"`
}

type EducationCourses struct {
	Name  string `bson:"name" binding:"required"`
	Grade string `bson:"grade"`
}
