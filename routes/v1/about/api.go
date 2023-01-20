package about

import (
	"Vitae/repositories"
	v1 "Vitae/routes/v1"
	"Vitae/tools/logging"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get one 'about', equivalent to an user
// Requires QUERY id=string
func GetOne(service IReader) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, ok := c.Params.Get("id")
        if !ok || len(id) == 0 {
            c.JSON(http.StatusBadRequest, v1.MessageResponse{
                Message: "Missing user ID as parameter",
            })
            return
        }
        var dto GetResponse
        err := service.GetOne(&dto, id)
        if _, ok := err.(*repositories.InvalidIdError); ok {
            logging.Debug(err.Error())
            c.JSON(http.StatusNotFound, v1.MessageResponse{
                Message: http.StatusText(http.StatusNotFound),
            })
            return
        } 
        if _, ok := err.(*repositories.NotFoundError); ok {
            logging.Debug(err.Error())
            c.JSON(http.StatusNotFound, v1.MessageResponse{
                Message: http.StatusText(http.StatusNotFound),
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
        c.JSON(http.StatusOK, dto)
    }
}

// Get all 'about's, equivalent to getting all users
// Requires QUERY amount=int
// ADMIN feature
func GetAll(service IReader) gin.HandlerFunc {
    return func(c *gin.Context) {
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
        dtos, err := service.GetAll(amount)
        logging.Trace("DTOs in controller layer", map[string]interface{}{"length": len(dtos)})
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

// Add an 'about', equivalent to add an user profile
// Requires a DTO
// Returns an ID of the newly created user
// ADMIN feature
func Post(service IWriter) gin.HandlerFunc {
    return func(c *gin.Context) {
        var dto PostRequest
        err := c.BindJSON(&dto)
        if err != nil {
            c.JSON(http.StatusBadRequest, v1.MessageResponse{
                Message: http.StatusText(http.StatusBadRequest),
            })
            return
        }
        logging.Trace("Post body", map[string]interface{}{"dto": dto})
        id, err := service.AddOne(dto)
        logging.Trace("Database result at handler level", map[string]interface{}{
            "id": id,
        })
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
