package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"

	rabbitmq "github.com/rabbitmq/amqp091-go"
	"github.com/thibleroy/trop-tiede/back/src/shared/broker"
)

var brokerConn *rabbitmq.Connection

func deviceHandler(w http.ResponseWriter, r *http.Request) {
	corrId := randomString(32)
	broker.PublishRPCRequest(*brokerConn, "device_rpc", "received ping from http-gateway", corrId)
	res := broker.ConsumeRPC(*brokerConn, "device_rpc", corrId)

	fmt.Fprintln(w, res)
}
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func main() {
	//username: default_user_UrXVpvKrxFg-4KXGxAq
	//password: ZKlaA4FqTUZP2_T1oTWrMWZw0SeQfSRG
	//BrokerUrl:      "host.docker.internal",
	broker_options := broker.BrokerClientOptions{
		BrokerUrl:      "localhost",
		BrokerPort:     "12345",
		BrokerUsername: "default_user_UrXVpvKrxFg-4KXGxAq",
		BrokerPassword: "ZKlaA4FqTUZP2_T1oTWrMWZw0SeQfSRG",
	}
	brokerConn, _ = broker.Connect(broker_options)
	http.HandleFunc("/device", deviceHandler)
	fmt.Println("Server starting...")
	err := http.ListenAndServe(":3000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
	}
}
