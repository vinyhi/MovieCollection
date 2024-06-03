package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %s", err)
	}
}

func connectDB() *mongo.Client {
	mongoURI := os.Getenv("MONGO_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to create new MongoDB client: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %s", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %s", err)
	}

	fmt.Println("Connected to MongoDB.")
	return client
}

func movieRoutes(router *gin.Engine, client *mongo.Client) {
	moviesCollection := client.Database("movieDB").Collection("movies")

	router.GET("/movies", func(c *gin.Context) {
		var movies []interface{}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		cur, err := moviesCollection.Find(ctx, "{}")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies from database"})
			log.Printf("Failed to fetch movies: %s", err)
			return
		}
		defer cur.Close(ctx)

		for cur.Next(ctx) {
			var movie interface{}
			err := cur.Decode(&movie)
			if err != nil {
				log.Printf("Error decoding movie: %s", err)
				continue
			}
			movies = append(movies, movie)
		}

		c.JSON(http.StatusOK, gin.H{"movies": movies})
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

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %s", err)
	}
}