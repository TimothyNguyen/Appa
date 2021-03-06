package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     redisURL,
	PoolSize: 10,
	DB:       0, // use default DB
})

// Set client options
var clientOptions = options.Client().ApplyURI(mongodbURL)

// Connect to MongoDB
var mongodb, err = mongo.Connect(context.Background(), clientOptions)

// connect to collection
var userCollection = mongodb.Database("appa-test").Collection("users")

func main() {
	fmt.Println("MongoDB: " + mongodbURL)
	fmt.Println("ACCESS_SECRET: " + string(accessSecret))
	fmt.Println("REFRESH_SECRET: " + string(refreshSecret))
	r := gin.Default()
	// Check the connection
	err = mongodb.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	pong, err := rdb.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong, err)
	r.Use(cors.Default())
	r.POST("/login", login)
	r.POST("/register", register)
	r.POST("/refresh", refresh)
	r.POST("/logout", logout)
	r.Run(":8000")
}
