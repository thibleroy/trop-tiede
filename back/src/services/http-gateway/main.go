package main

import (
	"errors"
	"fmt"
	"net/http"

	rabbitmq "github.com/rabbitmq/amqp091-go"
	"github.com/thibleroy/trop-tiede/back/src/shared/broker"
	"github.com/thibleroy/trop-tiede/back/src/shared/env"
	"github.com/thibleroy/trop-tiede/back/src/shared/utils"
)

var brokerConn *rabbitmq.Connection

func deviceHandler(w http.ResponseWriter, r *http.Request) {
	corrId := utils.RandomString(32)
	res, err := broker.RPC(*brokerConn, "hey, message from http gateway for device", "device_rpc", "device_rpc_cb", corrId)
	utils.HandleError(err, "error RPC")
	fmt.Fprintln(w, res)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	corrId := utils.RandomString(32)
	res, err := broker.RPC(*brokerConn, "hey, message from http gateway for user", "user_rpc", "user_rpc_cb", corrId)
	utils.HandleError(err, "error RPC")
	fmt.Fprintln(w, res)
}

func main() {
	env := env.GetServerEnv()
	client_options := broker.BrokerClientOptions{
		BrokerUrl:      env.RabbitMQBrokerUrl,
		BrokerPort:     env.RabbitMQBrokerPort,
		BrokerUsername: env.RabbitMQBrokerUsername,
		BrokerPassword: env.RabbitMQBrokerPassword,
	}
	brokerConn, _ = broker.Connect(client_options)
	http.HandleFunc("/device", deviceHandler)
	http.HandleFunc("/user", userHandler)
	fmt.Println("Server starting...")
	err := http.ListenAndServe(":3000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
	}
}
