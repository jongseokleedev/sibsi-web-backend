package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jongseokleedev/sibsi-web-backend/server/responses"
	"github.com/jongseokleedev/sibsi-web-backend/server/services/users"
	"net/http"
)

func SignUp(c *gin.Context) {
	result, err := users.SignUp(c)
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

func SignIn(c *gin.Context) {
	jwtToken, err := users.SignIn(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.HttpResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
		})
	} else {
		c.SetCookie("access-token", *jwtToken, 1800, "", "", false, false)
		c.JSON(http.StatusOK, responses.HttpResponse{
			Status:  http.StatusOK,
			Message: "success",
		})
	}

}
