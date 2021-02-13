package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
	"time"
)

func initMongoConn(url string, port int) (mongo.Client, context.Context) {
	uri := url + ":" + strconv.Itoa(port)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://" + uri))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to mongo database", uri)
	return *client, ctx
}

func InitDB(url string, port int, dbName string) (mongo.Database, context.Context) {
	client, c := initMongoConn(url, port)
	return *client.Database(dbName), c
}



