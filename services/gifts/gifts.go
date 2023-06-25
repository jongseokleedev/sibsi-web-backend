package gifts

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jongseokleedev/sibsi-web-backend/server/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Gift struct {
	Name             string `json:"name"`
	Description      string `json:"description""`
	TargetAmount     int64  `json:"target_amount"`
	CurrentAmount    int64  `json:"current_amount"`
	AmountVisibility bool   `json:"amount_visibility"`
	TargetDate       string `json:"target_date"`
}

func GetGift(c *gin.Context) (*Gift, error) {
	index, err := strconv.ParseInt(c.Param("index"), 10, 64)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	collection := configs.GetCollection(configs.DB, "gifts")
	var result Gift

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

func GetAllGifts(c *gin.Context) ([]*Gift, error) {
	collection := configs.GetCollection(configs.DB, "gifts")

	// 새 context.Context 객체 생성
	mongoCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(mongoCtx, bson.D{})
	if err != nil {
		fmt.Printf("err is %v", err)
		return nil, err
	}

	var results []*Gift
	if err = cursor.All(mongoCtx, &results); err != nil {
		log.Fatal(err)
	}

	return results, nil
}

func CreateNewGift(c *gin.Context) (*mongo.InsertOneResult, error) {
	var newGift Gift
	if err := c.BindJSON(&newGift); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}

	collection := configs.GetCollection(configs.DB, "gifts")
	insertResult, err := collection.InsertOne(context.Background(), newGift)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, err
	}
	return insertResult, nil
}

func UpdateGift(c *gin.Context) (*mongo.UpdateResult, error) {
	index, err := strconv.ParseInt(c.Param("index"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}

	var updatedGift Gift
	if err := c.BindJSON(&updatedGift); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}

	collection := configs.GetCollection(configs.DB, "gifts")
	//Update gift fields which is in body json format.
	update := bson.M{
		"$set": updatedGift,
	}

	updateResult, err := collection.UpdateOne(context.Background(), bson.M{"index": index}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, err
	}

	return updateResult, nil
}
