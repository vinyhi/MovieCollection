package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Movie struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Director    string             `bson:"director"`
	Genre       []string           `bson:"genre"`
	Review      string             `bson:"review,omitempty"`
	Rating      float64            `bson:"rating,omitempty"`
	ReleaseDate time.Time          `bson:"releaseDate"`
	UserID      primitive.ObjectID `bson:"userId"`
	CreatedAt   time.Time          `bson:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt"`
}

func main() {
	if err := loadEnvironment(); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	mongoURI := getMongoURI()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := connectMongoDB(ctx, mongoURI)
	defer disconnectMongoDB(ctx, client)

	collection := getMongoCollection(client, "moviesDB", "movies")
	insertExampleMovie(ctx, collection)
}

func loadEnvironment() error {
	return godotenv.Load()
}

func getMongoURI() string {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("You must set your 'MONGO_URI' environmental variable.")
	}
	return mongoURI
}

func connectMongoDB(ctx context.Context, mongoURI string) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	return client
}

func disconnectMongoDB(ctx context.Context, client *mongo.Client) {
	if err := client.Disconnect(ctx); err != nil {
		log.Fatal("Failed to disconnect MongoDB:", err)
	}
}

func getMongoCollection(client *mongo.Client, dbName, collectionName string) *mongo.Collection {
	return client.Database(dbName).Collection(collectionName)
}

func insertExampleMovie(ctx context.Context, collection *mongo.Collection) {
	exampleMovie := Movie{
		Title:       "Example Movie",
		Director:    "John Doe",
		Genre:       []string{"Drama", "Thriller"},
		Review:      "An example review.",
		Rating:      4.5,
		ReleaseDate: time.Now(),
		UserID:      primitive.NewObjectID(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	result, err := collection.InsertOne(ctx, exampleMovie)
	if err != nil {
		log.Fatal("Failed to insert document:", err)
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
}