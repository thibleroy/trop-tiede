package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func RegisterDatabase(context context.Context, config options.ClientOptions) (*mongo.Client, error) {
	client, err := mongo.Connect(context, options.Client().ApplyURI(""))
	if err != nil {
		zap.Error(err)
		return nil, err
	}
	return client, err
}
