package users

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jongseokleedev/sibsi-web-backend/server/configs"
	"github.com/jongseokleedev/sibsi-web-backend/server/services/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	"time"
)

//import
//import "github.com/jongseokleedev/sibsi-web-backend/server/configs"

type UserDto struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}
type User struct {
	ID       int    `json:"id"`
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

var secret = os.Getenv("SECRET")

func SignUp(c *gin.Context) (*mongo.InsertOneResult, error) {
	var newUser UserDto
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}

	collection := configs.GetCollection(configs.DB, "users")
	var result User

	// 새 context.Context 객체 생성
	mongoCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := collection.FindOne(mongoCtx, bson.M{"userid": newUser.UserID}).Decode(&result)
	if err == nil {
		fmt.Printf("user already exist, err is %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}

	insertResult, err := collection.InsertOne(context.Background(), newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, err
	}

	return insertResult, nil
}

func SignIn(c *gin.Context) (*string, error) {
	var user UserDto
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}

	collection := configs.GetCollection(configs.DB, "users")
	var result User

	// 새 context.Context 객체 생성
	mongoCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := collection.FindOne(mongoCtx, bson.M{"userid": user.UserID}).Decode(&result)
	if err != nil {
		fmt.Printf("user does not exist, err is %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}

	//@TODO pw hash verification
	if result.Password != user.Password {
		fmt.Printf("invalid password")

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, err
	}

	token, err := auth.GenerateToken(auth.NewClaim(user.UserID), secret)
	if err != nil {
		fmt.Printf("err is %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, err
	}

	return &token, nil

}
