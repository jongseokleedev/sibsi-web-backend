package configs

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func ConnectDB() *mongo.Client {

	EnvSetup()
	//fmt.Println(MongodbUser, " pw ", MongodbPassword)
	clientOptions := options.Client().ApplyURI("mongodb+srv://" + MongodbUser + ":" + MongodbPassword + "@sibsicluster0.vjrdgld.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	//[Warning] only admin allows can remove comments on this part
	//fmt.Println("Insert Documents into DB")
	//animals := services.CreateAnimals()
	//services.InsertDocuments(client, ctx, animals)

	return client
}

var DB = ConnectDB()

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(MongoDbDatabase).Collection(collectionName)
	return collection
}
