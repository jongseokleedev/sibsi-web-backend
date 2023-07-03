package gifts

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/jongseokleedev/sibsi-web-backend/server/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
)

type Provider struct {
	OwnerUserID      string `json:"owner_user_id"`
	Name             string `json:"name"`
	PhoneNumber      string `json:"phone_number"`
	NickName         string `json:"nick_name"`
	Message          string `json:"message"`
	Amount           int64  `json:"amount"`
	AmountVisibility bool   `json:"amount_visibility"`
}

type Gift struct {
	OwnerUserID      string   `json:"owner_user_id"`
	Name             string   `json:"name"`
	Image            string   `json:"image"`
	Description      string   `json:"description"`
	TargetAmount     int64    `json:"target_amount"`
	CurrentAmount    int64    `json:"current_amount"`
	AmountVisibility bool     `json:"amount_visibility"`
	TargetDate       string   `json:"target_date"`
	ProviderIds      []string `json:"providerids"`
}

func GetGift(c *gin.Context) (*Gift, error) {
	id := c.Param("id")                            // Get the id from the URL parameter
	objectId, err := primitive.ObjectIDFromHex(id) // Convert the id from string to ObjectID
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}
	collection := configs.GetCollection(configs.DB, "gifts")
	var result Gift

	// 새 context.Context 객체 생성
	mongoCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = collection.FindOne(mongoCtx, bson.M{"_id": objectId}).Decode(&result)
	if err != nil {
		fmt.Printf("err is %v", err)
		return nil, err
	}

	return &result, nil
}
func GetAllProviders(c *gin.Context) ([]*Provider, error) {
	id := c.Param("id")                                // Get the id from the URL parameter
	giftObjectId, err := primitive.ObjectIDFromHex(id) // Convert the id from string to ObjectID
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}

	// Find the gift
	giftCollection := configs.GetCollection(configs.DB, "gifts")
	var gift Gift
	err = giftCollection.FindOne(context.Background(), bson.M{"_id": giftObjectId}).Decode(&gift)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gift ID not found: " + id})
		return nil, err
	}

	// Find all providers for the gift
	providerCollection := configs.GetCollection(configs.DB, "providers")
	var providers []*Provider
	providerIds := make([]primitive.ObjectID, len(gift.ProviderIds))
	log.Printf("Gift ID: %s, Provider IDs: %v", id, gift.Name)
	for i, idStr := range gift.ProviderIds {
		providerId, err := primitive.ObjectIDFromHex(idStr)
		log.Printf("Gift ID: %s, Provider IDs: %v", id, providerId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return nil, err
		}
		providerIds[i] = providerId
	}

	log.Printf("Gift ID: %s, Provider IDs: %v", id, providerIds)
	cursor, err := providerCollection.Find(context.Background(), bson.M{"_id": bson.M{"$in": providerIds}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, err
	}
	if err = cursor.All(context.Background(), &providers); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, err
	}

	return providers, nil
}

func CreateNewGift(c *gin.Context) (*mongo.InsertOneResult, error) {
	//configs.Upload(c)

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
	var newGift Gift
	if err := c.BindJSON(&newGift); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}

	// userId 값 비교
	if newGift.OwnerUserID != userId {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid userId"})
		return nil, errors.New("invalid userId")
	}
	//Upload image to S3 aws server
	result, err := UploadImage(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, err
	}
	newGift.Image = result

	collection := configs.GetCollection(configs.DB, "gifts")
	insertResult, err := collection.InsertOne(context.Background(), newGift)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, err
	}
	newGiftId := insertResult.InsertedID.(primitive.ObjectID).Hex()
	receiverCollection := configs.GetCollection(configs.DB, "receivers")
	update := bson.M{
		"$push": bson.M{"giftid": newGiftId},
	}

	_, err = receiverCollection.UpdateOne(context.Background(), bson.M{"userid": userId}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, err
	}

	return insertResult, nil
}

func UpdateGift(c *gin.Context) (*mongo.UpdateResult, error) {
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

	var updatedGift Gift
	if err := c.BindJSON(&updatedGift); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}
	// userId 값 비교
	if updatedGift.OwnerUserID != userId {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid userId"})
		return nil, errors.New("invalid userId")
	}
	//Upload image to S3 aws server
	result, err := UploadImage(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, err
	}
	updatedGift.Image = result

	collection := configs.GetCollection(configs.DB, "gifts")
	//Update gift fields which is in body json format.
	update := bson.M{
		"$set": updatedGift,
	}

	updateResult, err := collection.UpdateOne(context.Background(), bson.M{"owner_user_id": userId}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, err
	}

	return updateResult, nil
}
func CreateProvider(c *gin.Context) (*mongo.InsertOneResult, error) {
	id := c.Param("id")                            // Get the id from the URL parameter
	objectID, err := primitive.ObjectIDFromHex(id) // Convert the id from string to ObjectID
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}

	// Parse the provider data from the request
	var newProvider Provider
	if err := c.BindJSON(&newProvider); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}

	// Insert the new provider
	providerCollection := configs.GetCollection(configs.DB, "providers")
	insertResult, err := providerCollection.InsertOne(context.Background(), newProvider)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, err
	}

	// Retrieve the ID of the newly inserted provider
	newProviderID := insertResult.InsertedID.(primitive.ObjectID).Hex()
	// Add the new provider ID to the gift
	giftCollection := configs.GetCollection(configs.DB, "gifts")
	update := bson.M{
		"$push": bson.M{"providerids": newProviderID},
	}

	_, err = giftCollection.UpdateOne(context.Background(), bson.M{"_id": objectID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, err
	}

	return insertResult, nil
}

func RemoveProviderFromGift(c *gin.Context) (*mongo.UpdateResult, error) {
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

	giftCollection := configs.GetCollection(configs.DB, "gifts")
	var result Gift

	// 새 context.Context 객체 생성
	mongoCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := giftCollection.FindOne(mongoCtx, bson.M{"owner_user_id": userId}).Decode(&result)
	if err != nil {
		fmt.Printf("cannot find gift for corresponding user Id %v", err)
		return nil, err
	}

	providerID := c.Param("providerID")

	// Check if provider ID exists in the providers collection
	providerCollection := configs.GetCollection(configs.DB, "providers")
	objectId, err := primitive.ObjectIDFromHex(providerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}
	var provider Provider
	err = providerCollection.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&provider)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Provider ID not found: " + providerID})
		return nil, err
	}

	// If provider ID is valid, remove it from the gift
	update := bson.M{
		"$pull": bson.M{"providerids": providerID},
	}
	updateResult, err := giftCollection.UpdateOne(context.Background(), bson.M{"owner_user_id": userId}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, err
	}

	// Delete the provider from the providers collection
	_, err = providerCollection.DeleteOne(context.Background(), bson.M{"_id": objectId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, err
	}

	return updateResult, nil
}

func UploadImage(c *gin.Context) (string, error) {
	file, _ := c.FormFile("file")
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// The session the S3 Uploader will use
	sess := configs.CreateSession()

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("sibsi-image-bucket"), // 버킷 이름
		Key:    aws.String(file.Filename),
		Body:   src,
	})
	if err != nil {
		return "", err
	}

	return result.Location, err
}
