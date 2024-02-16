package configs

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://mint:mint16@todolist.pw4xtxg.mongodb.net/?retryWrites=true&w=majority"))
    if err != nil {
        panic(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        panic(err)
    }

    err = client.Ping(context.TODO(), nil)
    if err != nil {
        panic(err)
    }
    fmt.Println("Connected to MongoDB")

    return client
}

var DB *mongo.Client = ConnectDB()
