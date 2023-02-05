package broker

//username: default_user_UrXVpvKrxFg-4KXGxAq
//password: ZKlaA4FqTUZP2_T1oTWrMWZw0SeQfSRG
import (
	"log"

	broker "github.com/rabbitmq/amqp091-go"
)

func consume(connection broker.Connection, topic string, handler brokerMessageHandler) {
	q, ch, err := queue(&connection, topic)
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
	var forever chan struct{}

	go func() {
		for d := range msgs {
			handler(d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
