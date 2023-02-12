package broker

//username: default_user_UrXVpvKrxFg-4KXGxAq
//password: ZKlaA4FqTUZP2_T1oTWrMWZw0SeQfSRG
import (
	"log"

	broker "github.com/rabbitmq/amqp091-go"
	"github.com/thibleroy/trop-tiede/back/src/shared/utils"
)

func Consume(connection broker.Connection, topic string, handler BrokerMessageHandler) {
	q, ch, err := Queue(&connection, topic)
	utils.HandleError(err, "Failed to register a queue")
	defer ch.Close()
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.HandleError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			handler(d)
		}
	}()

	log.Printf("Topic is " + topic + ". Waiting for messages.")

}
