package main

import (
	"encoding/json"

	rabbitmq "github.com/rabbitmq/amqp091-go"
	"github.com/thibleroy/trop-tiede/back/src/shared/broker"
	"go.uber.org/zap"
)

var brokerConn *rabbitmq.Connection

type Device struct {
	Id    string
	Label string
}

func handleRPCRequest(delivery rabbitmq.Delivery) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("Received message",
		// Structured context as strongly typed Field values.
		zap.String("topic", delivery.RoutingKey),
		zap.String("message", string(delivery.Body)),
	)
	device := Device{
		Id:    "azv12gb12jdbkcpn",
		Label: "device label",
	}
	deviceStr, err := json.Marshal(device)
	if err != nil {
		logger.Error("Error marshall", zap.Error(err))
	}
	broker.PublishRPCResponse(*brokerConn, delivery.RoutingKey, string(deviceStr), delivery)
}

func main() {
	//username: default_user_UrXVpvKrxFg-4KXGxAq
	//password: ZKlaA4FqTUZP2_T1oTWrMWZw0SeQfSRG
	//BrokerUrl:      "host.docker.internal",
	client_options := broker.BrokerClientOptions{
		BrokerUrl:      "localhost",
		BrokerPort:     "12345",
		BrokerUsername: "default_user_UrXVpvKrxFg-4KXGxAq",
		BrokerPassword: "ZKlaA4FqTUZP2_T1oTWrMWZw0SeQfSRG",
	}
	brokerConn, _ = broker.Connect(client_options)
	var forever chan struct{}
	//broker.Consume(*brokerConn, "device", testMessageHandlerC)
	broker.Consume(*brokerConn, "device_rpc", handleRPCRequest)
	<-forever
}
