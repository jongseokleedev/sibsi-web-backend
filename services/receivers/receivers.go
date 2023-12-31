package receivers

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jongseokleedev/sibsi-web-backend/server/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

//import
//import "github.com/jongseokleedev/sibsi-web-backend/server/configs"

type Receiver struct {
	UserId         string `json:"user_id"`
	Name           string `json:"name"`
	PhoneNumber    string `json:"phone_number""`
	AccountAddress string `json:"account_address"`
	AccountBank    string `json:"account_bank"`
	GiftId         string `json:"gift_id"`
}

func GetReceiver(c *gin.Context) (*Receiver, error) {
	value, ok := c.Get("user_id")
	if !ok {
		fmt.Println("userID not found")
		return nil, errors.New("id not found")
	}
	userId, ok := value.(string)
	if !ok {
		fmt.Println("user ID type casting error")
		return nil, errors.New("type casting error")
	}

	collection := configs.GetCollection(configs.DB, "receivers")
	var result Receiver

	// 새 context.Context 객체 생성
	mongoCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := collection.FindOne(mongoCtx, bson.M{"user_id": userId}).Decode(&result)
	if err != nil {
		fmt.Printf("err is %v", err)
		return nil, err
	}

	return &result, nil
}

func CreateNewReceiver(c *gin.Context) (*mongo.InsertOneResult, error) {
	userId, ok := c.Get("user_id")
	if !ok {
		fmt.Println("userID not found")
		return nil, errors.New("id not found")
	}
	var newReceiver Receiver
	if err := c.BindJSON(&newReceiver); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}
	newReceiver.UserId = userId.(string)

	collection := configs.GetCollection(configs.DB, "receivers")
	insertResult, err := collection.InsertOne(context.Background(), newReceiver)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, err
	}
	return insertResult, nil
}
