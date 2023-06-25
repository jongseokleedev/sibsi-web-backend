package controllers

import (
	"github.com/GCI-js/bangalzoo/server/responses"
	"github.com/GCI-js/bangalzoo/server/services/receivers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetReceiver(c *gin.Context) {

	var receiver *receivers.Receiver
	receiver, err := receivers.GetReceivers(c)
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
