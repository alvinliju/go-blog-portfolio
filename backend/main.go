package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	//mongo drivers
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//dotenv
	"github.com/joho/godotenv"
)

var mongoClient *mongo.Client

// init function will run before the main function
// func init() {
// 	if err := connectDB(); err != nil {
// 		log.Fatal("Could not connect to mongo")
// 	}

// 	fmt.Println("Connected to DB")
// }
//

// connect db function
func connectDB() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	URI := os.Getenv("DB_URI")

	fmt.Println(URI)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(URI).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	fmt.Println("DB connected")
	mongoClient = client
	return err
}

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world",
	})
}

func main() {
	r := gin.Default()

	r.GET("/", Hello)

	connectDB()
	r.Run()
}
