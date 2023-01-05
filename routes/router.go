package routes

import (
	about "Vitae/routes/v1/about"
	"github.com/gin-gonic/gin"
)

func Create(engine *gin.Engine) {
	engine.Use(gin.Logger())

	// Returns a 500 on panic
	engine.Use(gin.Recovery())

	v1 := engine.Group("/api/v1")

	v1.GET("/about", about.Get)
}
