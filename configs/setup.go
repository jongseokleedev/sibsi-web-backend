package configs

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
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

var client *redis.Client

func RedisInit() {
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}

	client = redis.NewClient(&redis.Options{
		Addr: dsn,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func GetClient() *redis.Client {
	return client
}
