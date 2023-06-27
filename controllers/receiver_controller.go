package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jongseokleedev/sibsi-web-backend/server/responses"
	"github.com/jongseokleedev/sibsi-web-backend/server/services/receivers"
	"net/http"
)

func GetReceiver(c *gin.Context) {

	var receiver *receivers.Receiver
	receiver, err := receivers.GetReceiver(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
		})
	} else {

		c.JSON(http.StatusOK, responses.HttpResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": &receiver},
		})
	}
}

func CreateNewReceiver(c *gin.Context) {

	result, err := receivers.CreateNewReceiver(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
		})
	} else {
		c.JSON(http.StatusOK, responses.HttpResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": &result},
		})
	}
}
