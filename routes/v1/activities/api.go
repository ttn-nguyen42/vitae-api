package activities

import (
	"Vitae/repositories"
	v1 "Vitae/routes/v1"
	"Vitae/tools/logging"
	"Vitae/tools/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetOne(activitiesService IReader) gin.HandlerFunc {
	return func(c *gin.Context) {
		activityId, ok := c.Params.Get("activityId")
		if !ok || len(activityId) == 0 {
			c.JSON(http.StatusBadRequest, v1.MessageResponse{
				Message: "Missing activityId as parameter",
			})
			return
		}
		var dto GetResponse
		err := activitiesService.GetOne(&dto, activityId)
		if err != nil {
			logging.Debug(err.Error())
			c.JSON(http.StatusInternalServerError, v1.MessageResponse{
				Message: http.StatusText(http.StatusInternalServerError),
			})
			return
		}
		c.JSON(http.StatusOK, dto)
	}
}

func GetAll(activitiesService IReader) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, ok := c.Params.Get("id")
		if !ok || len(userId) == 0 {
			c.JSON(http.StatusBadRequest, v1.MessageResponse{
				Message: "Missing user ID as parameter",
			})
			return
		}
		query, ok := c.GetQuery("amount")
		var amount int
		var err error
		if !ok || amount <= 0 {
			amount = repositories.Query10
		}
		if ok {
			amount, err = strconv.Atoi(query)
		}
		logging.Trace("Amount in query", map[string]interface{}{"amount": amount})
		if err != nil {
			c.JSON(http.StatusBadRequest, v1.MessageResponse{
				Message: "Incorrect query type for 'amount'",
			})
			return
		}
		dtos, err := activitiesService.GetAll(userId, amount)
		logging.Trace("DTOs in controller layer", map[string]interface{}{"length": len(dtos)})
		if _, ok := err.(*repositories.NotFoundError); ok {
			logging.Debug(err.Error())
			c.JSON(http.StatusBadRequest, v1.MessageResponse{
				Message: "Invalid user ID",
			})
			return
		}
		if err != nil {
			logging.Debug(err.Error())
			c.JSON(http.StatusInternalServerError, v1.MessageResponse{
				Message: http.StatusText(http.StatusInternalServerError),
			})
			return
		}
		c.JSON(http.StatusOK, dtos)
	}
}

func Post(activitiesService IWriter) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, ok := c.Params.Get("id")
		if !ok || len(userId) == 0 {
			c.JSON(http.StatusBadRequest, v1.MessageResponse{
				Message: "Missing user ID as parameter",
			})
			return
		}
		var dto PostRequest
		err := c.BindJSON(&dto)
		if err != nil {
			if !utils.IsProduction() {
				c.JSON(http.StatusBadRequest, v1.MessageResponse{
					Message: err.Error(),
				})
				return
			}
			c.JSON(http.StatusBadRequest, v1.MessageResponse{
				Message: http.StatusText(http.StatusBadRequest),
			})
			return
		}
		logging.Trace("Post body", map[string]interface{}{"dto": dto})
		id, err := activitiesService.AddOne(userId, dto)
		logging.Trace("Database result at handler level", map[string]interface{}{
			"id": id,
		})
		if _, ok = err.(*repositories.NotFoundError); ok {
			logging.Debug(err.Error())
			c.JSON(http.StatusBadRequest, v1.MessageResponse{
				Message: "User not found",
			})
			return
		}
		if _, ok = err.(*repositories.InvalidIdError); ok {
			logging.Debug(err.Error())
			c.JSON(http.StatusBadRequest, v1.MessageResponse{
				Message: "Invalid user ID",
			})
			return
		}
		if err != nil {
			logging.Debug(err.Error())
			c.JSON(http.StatusInternalServerError, v1.MessageResponse{
				Message: http.StatusText(http.StatusInternalServerError),
			})
			return
		}
		c.JSON(http.StatusCreated, v1.IdResponse{
			Id: id,
		})
	}
}
