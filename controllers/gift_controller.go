package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jongseokleedev/sibsi-web-backend/server/responses"
	"github.com/jongseokleedev/sibsi-web-backend/server/services/gifts"
	"net/http"
)

func GetGift(c *gin.Context) {
	var gift *gifts.Gift
	gift, err := gifts.GetGift(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
		})
	} else {
		c.JSON(http.StatusOK, responses.HttpResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": &gift},
		})
	}
}

func GetAllProviders(c *gin.Context) {
	var providers []*gifts.Provider
	providers, err := gifts.GetAllProviders(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
		})
	} else {
		c.JSON(http.StatusOK, responses.HttpResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": &providers},
		})
	}
}

func CreateNewGift(c *gin.Context) {

	result, err := gifts.CreateNewGift(c)
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

func UpdateGift(c *gin.Context) {

	result, err := gifts.UpdateGift(c)
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

func CreateNewProvider(c *gin.Context) {

	result, err := gifts.CreateProvider(c)
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

func RemoveProvider(c *gin.Context) {

	result, err := gifts.RemoveProviderFromGift(c)
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
