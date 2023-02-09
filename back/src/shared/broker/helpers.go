package broker

//username: default_user_UrXVpvKrxFg-4KXGxAq
//password: ZKlaA4FqTUZP2_T1oTWrMWZw0SeQfSRG
import (
	"log"

	broker "github.com/rabbitmq/amqp091-go"
)

type BrokerMessageHandler func(broker.Delivery)

type BrokerClientOptions struct {
	BrokerUrl      string
	BrokerPort     string
	BrokerUsername string
	BrokerPassword string
}

func formatBrokerURL(broker_client_options BrokerClientOptions) string {
	return "amqp://" + broker_client_options.BrokerUrl + ":" + broker_client_options.BrokerPort + "/"
}

func handleError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}

func Connect(client_options BrokerClientOptions) (*broker.Connection, error) {
	brokerConfig := broker.Config{
		SASL: []broker.Authentication{&broker.AMQPlainAuth{Username: client_options.BrokerUsername, Password: client_options.BrokerPassword}},
	}
	conn, err := broker.DialConfig(formatBrokerURL(client_options), brokerConfig)
	handleError(err, "Failed to connect")
	return conn, err
}

func Queue(connection *broker.Connection, topic string) (broker.Queue, *broker.Channel, error) {
	ch, err := connection.Channel()
	handleError(err, "Failed to connect")
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
