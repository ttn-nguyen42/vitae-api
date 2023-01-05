package activities

import (
	"Vitae/config/database"
	"Vitae/repositories/activities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	_ = activities.New(database.Client, database.Context)
	c.JSON(http.StatusOK, nil)
}