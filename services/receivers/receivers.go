package receivers

import (
	"context"
	"fmt"
	"github.com/GCI-js/bangalzoo/server/configs"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
	"time"
)

//import
//import "github.com/GCI-js/bangalzoo/server/configs"

type Receiver struct {
	Index          int64
	Name           string
	AccountAddress string
	AccountBank    string
	giftId         int64
}

func GetReceivers(c *gin.Context) (*Receiver, error) {
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
