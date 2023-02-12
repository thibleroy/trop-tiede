package broker

//username: default_user_UrXVpvKrxFg-4KXGxAq
//password: ZKlaA4FqTUZP2_T1oTWrMWZw0SeQfSRG
import (
	"context"
	"time"

	broker "github.com/rabbitmq/amqp091-go"
	"github.com/thibleroy/trop-tiede/back/src/shared/utils"
	"go.uber.org/zap"
)

func Publish(connection broker.Connection, topic string, message string) error {
	q, ch, err := Queue(&connection, topic)
	utils.HandleError(err, "Failed to retrieve a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("Published message",
		zap.String("topic", topic),
		zap.String("message", message),
	)
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
