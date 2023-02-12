package main

import (
	"encoding/json"

	rabbitmq "github.com/rabbitmq/amqp091-go"
	"github.com/thibleroy/trop-tiede/back/src/shared/broker"
	"github.com/thibleroy/trop-tiede/back/src/shared/env"
)

var brokerConn *rabbitmq.Connection

type Device struct {
	Id    string
	Label string
}

func main() {
	env := env.GetServerEnv()
	//username: default_user_UrXVpvKrxFg-4KXGxAq
	//password: ZKlaA4FqTUZP2_T1oTWrMWZw0SeQfSRG
	//BrokerUrl:      "host.docker.internal",
	client_options := broker.BrokerClientOptions{
		BrokerUrl:      env.RabbitMQBrokerUrl,
		BrokerPort:     env.RabbitMQBrokerPort,
		BrokerUsername: env.RabbitMQBrokerUsername,
		BrokerPassword: env.RabbitMQBrokerPassword,
	}
	brokerConn, _ = broker.Connect(client_options)
	device := Device{Id: "123", Label: "hello !"}
	deviceStr, _ := json.Marshal(device)
	broker.ServeRPC(*brokerConn, "device_rpc", string(deviceStr))
}
