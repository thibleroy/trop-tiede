package main

import (
	"fmt"
	"net/http"

	"github.com/rabbitmq/amqp091-go"
	"github.com/thibleroy/trop-tiede/back/src/shared/broker"
)

var brokerConn *amqp091.Connection

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
	broker.Publish(*brokerConn, "testaa", "received index")
}

func check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Health check</h1>")
	broker.Publish(*brokerConn, "testaa", "received health_check")
}

func main() {
	brokerConn, _ = broker.Connect("default_user_UrXVpvKrxFg-4KXGxAq", "ZKlaA4FqTUZP2_T1oTWrMWZw0SeQfSRG")
	http.HandleFunc("/", index)
	http.HandleFunc("/health_check", check)
	fmt.Println("Server starting...")
	http.ListenAndServe(":3000", nil)
}
