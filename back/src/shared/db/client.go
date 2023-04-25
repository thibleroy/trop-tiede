package db

import (
	"context"

	"github.com/thibleroy/trop-tiede/back/src/shared/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func RegisterDatabase(context context.Context, config options.ClientOptions) (*mongo.Client, error) {
	client, err := mongo.Connect(context, &config)
	utils.HandleError(err, "Error connection db"+config.GetURI())
	return client, err
}
