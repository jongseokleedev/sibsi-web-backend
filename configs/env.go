package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	Port    uint64
	GinMode string

	MongodbUser     string
	MongodbPassword string
	MongoDbDatabase string
)

func EnvSetup() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("%v", err)
		log.Fatal("Error Loading .env file. Please put your .env file in root project")
	}

	log.Print("configs initialize...")
	Port, _ = strconv.ParseUint(os.Getenv("PORT"), 0, 64)

	MongodbUser = os.Getenv("MONGODB_USER")
	MongodbPassword = os.Getenv("MONGODB_PASSWORD")
	MongoDbDatabase = os.Getenv("MONGODB_DATABASE")
}
