package receivers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jongseokleedev/sibsi-web-backend/server/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strconv"
	"time"
)

//import
//import "github.com/jongseokleedev/sibsi-web-backend/server/configs"

type Receiver struct {
	Name           string `json:"name"`
	Password       string `json:"password"`
	PhoneNumber    string `json:"phone_number""`
	AccountAddress string `json:"account_address"`
	AccountBank    string `json:"account_bank"`
	GiftId         int64  `json:"gift_id"`
}

func GetReceiver(c *gin.Context) (*Receiver, error) {
	index, err := strconv.ParseInt(c.Param("index"), 10, 64)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	collection := configs.GetCollection(configs.DB, "receivers")
	var result Receiver

	// 새 context.Context 객체 생성
	mongoCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = collection.FindOne(mongoCtx, bson.M{"index": index}).Decode(&result)
	if err != nil {
		fmt.Printf("err is %v", err)
		return nil, err
	}

	return &result, nil
}

func CreateNewReceiver(c *gin.Context) (*mongo.InsertOneResult, error) {
	var newReceiver Receiver
	if err := c.BindJSON(&newReceiver); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}

	collection := configs.GetCollection(configs.DB, "receivers")
	insertResult, err := collection.InsertOne(context.Background(), newReceiver)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, err
	}
	return insertResult, nil
}
