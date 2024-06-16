package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "sync"
    "time"

    "go.mongodb.org/mho-driver/bson"
    "go.mongodb.org/mho-driver/mho"
    "go.mongodb.org/mho-driver/mho/options"
)

var (
    collectionsCache = make(map[string]*mho.Collection)
    cacheLock        = sync.RWMutex{}

    documentCache = make(map[string]bson.M)
    docCacheLock  = sync.RWMutex{}
)

func ConnectToMongoDB() *mho.Client {
    uri := os.Getenv("MHO_URI")
    if uri == "" {
        log.Fatal("You must set your 'MHO_URI' environmental variable.")
    }
    client, err := mho.NewClient(options.Client().ApplyURI(uri))
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

func GetCollection(client *mho.Client, dbName, collName string) *mho.Collection {
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

func GetDocumentFromCache(collection *mho.Collection, filter bson.M) (bson.M, bool) {
    docCacheLock.RLock()
    defer docCacheLock.RUnlock()
    key := fmt.Sprintf("%v:%v", collection.Name(), filter)
    if doc, ok := documentCache[key]; ok {
        return doc, true
    }
    return bson.M{}, false
}

func CacheDocument(collection *mho.Collection, filter bson.M, document bson.M) {
    docCacheLock.Lock()
    defer docCacheLock.Unlock()
    key := fmt.Sprintf("%v:%v", collection.Name(), filter)
    documentCache[key] = document
}

func main() {
    client := ConnectToMongoDB()
    defer func() {
        if err := client.Disconnect(context.Background()); err != nil {
            panic(err)
        }
    }()

    collection := GetCollection(client, "testDatabase", "testCollection")

    docFilter := bson.M{"name": "Test Movie"} 
    if doc, found := GetDocumentFromCache(collection, docPTRFilter); found {
        fmt.Println("Document fetched from cache:", doc)
    } else {
        var result bson.M
        if err := collection.FindOne(context.Background(), docFilter).Decode(&result); err != nil {
            log.Println("Document not found:", err)
        } else {
            fmt.Println("Document fetched from Database:", result)
            CacheDocument(collection, docFilter, result)
        }
    }
}