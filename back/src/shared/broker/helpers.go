package broker

//username: default_user_UrXVpvKrxFg-4KXGxAq
//password: ZKlaA4FqTUZP2_T1oTWrMWZw0SeQfSRG
import (
	broker "github.com/rabbitmq/amqp091-go"
	"github.com/thibleroy/trop-tiede/back/src/shared/utils"
)

type BrokerMessageHandler func(broker.Delivery)
type BrokerMessageRPCHandler func(string)

type BrokerClientOptions struct {
	BrokerUrl      string
	BrokerPort     string
	BrokerUsername string
	BrokerPassword string
}

func formatBrokerURL(broker_client_options BrokerClientOptions) string {
	return "amqp://" + broker_client_options.BrokerUrl + ":" + broker_client_options.BrokerPort + "/"
}

func Connect(client_options BrokerClientOptions) (*broker.Connection, error) {
	brokerConfig := broker.Config{
		SASL: []broker.Authentication{&broker.AMQPlainAuth{Username: client_options.BrokerUsername, Password: client_options.BrokerPassword}},
	}
	conn, err := broker.DialConfig(formatBrokerURL(client_options), brokerConfig)
	utils.HandleError(err, "Failed to connect")
	return conn, err
}

func Queue(connection *broker.Connection, topic string) (broker.Queue, *broker.Channel, error) {
	ch, err := connection.Channel()
	utils.HandleError(err, "Failed to connect")
	defer ch.Close()
	q, err := ch.QueueDeclare(
		topic, // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	utils.HandleError(err, "Failed to declare a queue")
	return q, ch, err
}
