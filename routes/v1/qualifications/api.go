package qualifications

import (
	"Vitae/config/database"
	"Vitae/repositories/qualifications"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	_ = qualifications.New(database.Client)
	c.JSON(http.StatusOK, nil)
}