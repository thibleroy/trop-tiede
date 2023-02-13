package broker

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/thibleroy/trop-tiede/back/src/shared/utils"
)

func RPC(conn amqp091.Connection, message string, rpc_queue_name string, rpc_cb_queue_name string, correlation_id string) (res interface{}, err error) {
	ch, err := conn.Channel()
	utils.HandleError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		rpc_cb_queue_name, // name
		false,             // durable
		false,             // delete when unused
		true,              // exclusive
		false,             // noWait
		nil,               // arguments
	)
	utils.HandleError(err, "Failed to declare a queue")

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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",             // exchange
		rpc_queue_name, // routing key
		false,          // mandatory
		false,          // immediate
		amqp091.Publishing{
			ContentType:   "text/plain",
			CorrelationId: correlation_id,
			ReplyTo:       q.Name,
			Body:          []byte(message),
		})
	utils.HandleError(err, "Failed to publish a message")
	fmt.Println("Sent RPC request", message)

	for d := range msgs {
		if correlation_id == d.CorrelationId {
			res = string(d.Body)
			fmt.Println("Received RPC response", res)
			utils.HandleError(err, "Failed to convert body to integer")
			break
		}
	}
	return
}

type RPCRequestHandler func(amqp091.Delivery, context.Context) error

func PublishRPCResponse(conn amqp091.Connection, d amqp091.Delivery, message string, ctx context.Context) {
	ch, err := conn.Channel()
	utils.HandleError(err, "Failed to open a channel")
	defer ch.Close()
	ch.PublishWithContext(ctx,
		"",        // exchange
		d.ReplyTo, // routing key
		false,     // mandatory
		false,     // immediate
		amqp091.Publishing{
			ContentType:   "text/plain",
			CorrelationId: d.CorrelationId,
			Body:          []byte(message),
		})
	utils.HandleError(err, "Failed to publish a message")
	fmt.Println("Sent RPC response", message)
	d.Ack(false)
}

func ServeRPC(conn amqp091.Connection, rpc_queue_name string, ctx context.Context, handler RPCRequestHandler) {
	ch, err := conn.Channel()
	utils.HandleError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		rpc_queue_name, // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	utils.HandleError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	utils.HandleError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.HandleError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			fmt.Println("Received RPC request, body is", string(d.Body))
			err := handler(d, ctx)
			utils.HandleError(err, "Error handling RPC request")
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}
