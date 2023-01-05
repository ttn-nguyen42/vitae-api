package education

import (
	"Vitae/config/database"
	"Vitae/repositories/education"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	_ = education.New(database.Client, database.Context)
	c.JSON(http.StatusOK, nil)
}