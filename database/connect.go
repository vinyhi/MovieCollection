package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	collectionsCache = make(map[string]*mongo.Collection)
	cacheLock        = sync.RWMutex{}
)

func ConnectToMongoDB() *mongo.Client {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGO_URI' environmental variable.")
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected and pinged.")
	return client
}

func GetCollection(client *mongo.Client, dbName, collName string) *mongo.Collection {
	cacheKey := dbName + ":" + collName
	cacheLock.RLock()
	if collection, ok := collectionsCache[cacheKey]; ok {
		cacheLock.RUnlock()
		return collection
	}
	cacheLock.RUnlock()

	collection := client.Database(dbName).Collection(collName)

	cacheLock.Lock()
	collectionsCache[cacheKey] = collection
	cacheLock.Unlock()

	return collection
}

func main() {
	client := ConnectToMongoDB()
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	collection := GetCollection(client, "testDatabase", "testCollection")
	fmt.Println(collection)
}