package routes

import (
	"Vitae/config/database"
	aboutRepository "Vitae/repositories/about"
	"Vitae/routes/v1/about"
	"Vitae/routes/v1/activities"
	"Vitae/routes/v1/certificates"
	"Vitae/routes/v1/education"
	"Vitae/routes/v1/experiences"
	"Vitae/routes/v1/qualifications"
	"Vitae/routes/v1/skills"
	"Vitae/routes/v1/social"

	"github.com/gin-gonic/gin"
)

func Create(engine *gin.Engine) {
	engine.Use(gin.Logger())

	// Returns a 500 on panic
	engine.Use(gin.Recovery())

	v1 := engine.Group("/api/v1")

	// MongoDB client
	client := database.Client

	// Repositories
	aboutRepo := aboutRepository.New(client)

	// Services
	aboutService := about.NewService(aboutRepo)

	v1.GET("/about", about.GetAll(aboutService))
	v1.GET("/about/:id", about.GetOne(aboutService))
	v1.POST("/about", about.Post(aboutService))
	
	v1.GET("/activities", activities.Get)
	v1.GET("/certificates", certificates.Get)
	v1.GET("/education", education.Get)
	v1.GET("/experiences", experiences.Get)
	v1.GET("/qualifications", qualifications.Get)
	v1.GET("/skills", skills.Get)
	v1.GET("/social", social.Get)
}