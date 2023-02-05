package main

import (
	rabbitmq "github.com/rabbitmq/amqp091-go"
	"github.com/thibleroy/trop-tiede/back/src/shared/broker"
	"go.uber.org/zap"
)

const DeviceCollectionName = "Device"

func testMessageHandlerC(delivery rabbitmq.Delivery) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("Received message",
		// Structured context as strongly typed Field values.
		zap.String("topic", delivery.RoutingKey),
		zap.String("message", string(delivery.Body)),
	)
}

func main() {
	//username: default_user_UrXVpvKrxFg-4KXGxAq
	//password: ZKlaA4FqTUZP2_T1oTWrMWZw0SeQfSRG
	conn, _ := broker.Connect("default_user_UrXVpvKrxFg-4KXGxAq", "ZKlaA4FqTUZP2_T1oTWrMWZw0SeQfSRG")
	broker.Consume(*conn, "testaa", testMessageHandlerC)
}
