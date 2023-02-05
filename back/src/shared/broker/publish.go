package broker

//username: default_user_UrXVpvKrxFg-4KXGxAq
//password: ZKlaA4FqTUZP2_T1oTWrMWZw0SeQfSRG
import (
	"context"
	"time"

	broker "github.com/rabbitmq/amqp091-go"
)

func Publish(connection broker.Connection, topic string, message string) error {
	q, ch, err := Queue(&connection, topic)
	handleError(err, "Failed to retrieve a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		broker.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}
