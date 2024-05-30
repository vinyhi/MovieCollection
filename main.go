package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func connectDB() *mongo.Client {
	mongoURI := os.Getenv("MONGO_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB.")
	return client
}

func movieRoutes(router *gin.Engine, client *mongo.Client) {
	moviesCollection := client.Database("movieDB").Collection("movies")
	router.GET("/movies", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "List of movies"})
	})
}

func main() {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	client := connectDB()

	movieRoutes(router, client)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}