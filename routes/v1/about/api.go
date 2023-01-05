package about

import (
	"Vitae/config/database"
	"Vitae/repositories/about"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get(c *gin.Context) {
	_ = about.New(database.Client, database.Context)
	c.JSON(http.StatusOK, nil)
}
