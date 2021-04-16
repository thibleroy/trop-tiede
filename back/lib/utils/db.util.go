package utils

import (
	"back/lib"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func NewResource() lib.IResource {
	creationTime := time.Now()
	return lib.IResource{
		ID:        primitive.NewObjectIDFromTimestamp(creationTime),
		CreatedAt: creationTime,
		UpdatedAt: creationTime,
	}
}

func initMongoConn(uri string) mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	return *client
}

func InitDB(url string, dbName string) mongo.Database {
	client := initMongoConn(url)
	ctx, c := context.WithTimeout(context.Background(), 2*time.Second)
	defer c()
	if client.Ping(ctx, nil) != nil {
		log.Fatal("Error pinging mongoDB")
	}
	fmt.Println("connected to mongo database", url)
	return *client.Database(dbName)
}
