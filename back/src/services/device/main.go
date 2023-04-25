package main

import (
	"context"
	"encoding/json"
	"time"

	rabbitmq "github.com/rabbitmq/amqp091-go"
	"github.com/thibleroy/trop-tiede/back/src/shared/broker"
	"github.com/thibleroy/trop-tiede/back/src/shared/env"
	"github.com/thibleroy/trop-tiede/back/src/shared/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var brokerConn *rabbitmq.Connection
var dbConn *mongo.Client

type Device struct {
	Id    string
	Label string
}

type User struct {
	Id        string
	Firstname string
}

type Response struct {
	Device Device
	User   User
}

func handleDeviceRPCRequest(delivery rabbitmq.Delivery, ctx context.Context) error {
	device := Device{Id: "123", Label: "hello !"}
	deviceStr, err := json.Marshal(device)
	corrId := utils.RandomString(32)
	res, err := broker.RPC(*brokerConn, "hey, message from http gateway for user", "user_rpc", "user_rpc_cb", corrId)
response := Response{Device: device, res}
	broker.PublishRPCResponse(*brokerConn, delivery, string(deviceStr), ctx)
	return err
}

func main() {
	env := env.GetServerEnv()
	client_options := broker.BrokerClientOptions{
		BrokerUrl:      env.RabbitMQBrokerUrl,
		BrokerPort:     env.RabbitMQBrokerPort,
		BrokerUsername: env.RabbitMQBrokerUsername,
		BrokerPassword: env.RabbitMQBrokerPassword,
	}
	db_client_options := options.Client().ApplyURI("mongodb://localhost:55550/")
	db_client_credentials := options.Credential{Username: "my-user", Password: "vivititi"}
	db_client_options.SetAuth(db_client_credentials)
	db_client_options.SetDirect(true)
	brokerConn, _ = broker.Connect(client_options)
	// dbConn, _ = db.RegisterDatabase(context.Background(), *db_client_options)
	// db := dbConn.Database("tesg")
	// coll := db.Collection("hey test")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// _, err := coll.InsertOne(ctx, Device{Id: "123", Label: "hello !"})
	// utils.HandleError(err, "Error inserting data")
	// ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	broker.ServeRPC(*brokerConn, "device_rpc", ctx, handleDeviceRPCRequest)
}
