package broker

//username: default_user_UrXVpvKrxFg-4KXGxAq
//password: ZKlaA4FqTUZP2_T1oTWrMWZw0SeQfSRG
import (
	"log"

	broker "github.com/rabbitmq/amqp091-go"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}

func Connect(url string, username string, password string) (*broker.Connection, error) {
	brokerConfig := broker.Config{
		SASL: []broker.Authentication{&broker.AMQPlainAuth{Username: username, Password: password}},
	}
	conn, err := broker.DialConfig(url, brokerConfig)
	handleError(err, "Failed to connect")
	return conn, err
}

type BrokerMessageHandler func(broker.Delivery)

func Queue(connection *broker.Connection, topic string) (broker.Queue, *broker.Channel, error) {
	ch, err := connection.Channel()
	q, err := ch.QueueDeclare(
		topic, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	handleError(err, "Failed to declare a queue")
	return q, ch, err
}
