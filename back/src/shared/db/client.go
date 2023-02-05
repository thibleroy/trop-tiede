package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connect(param string) (*mongo.Client, error) {
	return mongo.Connect(context.TODO(), options.Client().ApplyURI(""))
}
