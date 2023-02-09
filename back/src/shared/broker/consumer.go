package broker

//username: default_user_UrXVpvKrxFg-4KXGxAq
//password: ZKlaA4FqTUZP2_T1oTWrMWZw0SeQfSRG
import (
	"log"
	"net/http"

	broker "github.com/rabbitmq/amqp091-go"
)

func Consume(connection broker.Connection, topic string, handler BrokerMessageHandler) {
	q, ch, err := Queue(&connection, topic)
	handleError(err, "Failed to register a queue")
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	handleError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			handler(d)
		}
	}()

	log.Printf("Topic is " + topic + ". Waiting for messages.")

}

func ConsumeRPC(connection broker.Connection, topic string, handler BrokerRPCMessageHandler, correlation_id string, w http.ResponseWriter) {
	q, ch, err := Queue(&connection, topic)
	handleError(err, "Failed to register a queue")
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	handleError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			if d.CorrelationId == correlation_id {
				handler(d, w)
			}
		}
	}()

	log.Printf("Topic is " + topic + ". Waiting for messages.")
}
