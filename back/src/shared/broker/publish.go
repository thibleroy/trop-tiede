package broker

//username: default_user_UrXVpvKrxFg-4KXGxAq
//password: ZKlaA4FqTUZP2_T1oTWrMWZw0SeQfSRG
import (
	"context"
	"time"

	broker "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

func Publish(connection broker.Connection, topic string, message string) error {
	q, ch, err := Queue(&connection, topic)
	handleError(err, "Failed to retrieve a queue")
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

func PublishRPCResponse(connection broker.Connection, topic string, message string, delivery broker.Delivery) error {
	_, ch, err := Queue(&connection, topic)
	handleError(err, "Failed to retrieve a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("Published message",
		zap.String("topic", topic),
		zap.String("message", message),
	)
	return ch.PublishWithContext(ctx,
		"",               // exchange
		delivery.ReplyTo, // routing key
		false,            // mandatory
		false,            // immediate
		broker.Publishing{
			ContentType:   "text/plain",
			Body:          []byte(message),
			CorrelationId: delivery.CorrelationId,
		})
}

func PublishRPCRequest(connection broker.Connection, topic string, message string, correlation_id string) error {
	q, ch, err := Queue(&connection, topic)
	handleError(err, "Failed to retrieve a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("Published message",
		zap.String("topic", topic),
		zap.String("message", message),
	)
	err = ch.PublishWithContext(ctx,
		"",    // exchange
		topic, // routing key
		false, // mandatory
		false, // immediate
		broker.Publishing{
			ContentType:   "text/plain",
			Body:          []byte(message),
			CorrelationId: correlation_id,
			ReplyTo:       q.Name,
		})
	handleError(err, "Failed to publish")
	return err
}
