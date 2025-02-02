package rabbitmq

import (
	"context"
	"log"
	"math/rand"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/torchiaf/Sensors/rpc_client/models"
	"github.com/torchiaf/Sensors/rpc_client/utils"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())

	return min + rand.Intn(max-min)
}

type Client struct {
	ctx           context.Context
	connection    *amqp.Connection
	channel       *amqp.Channel
	responseQueue amqp.Queue
	msgs          <-chan amqp.Delivery
}

func (client *Client) Send(routingKey string, message models.Message) (res string, err error) {
	corrId := randomString(32)

	log.Printf("Msg: %s:", utils.ToString(message))

	err = client.channel.PublishWithContext(client.ctx,
		"",         // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       client.responseQueue.Name,
			Body:          []byte(utils.ToString(message)),
		})
	failOnError(err, "Failed to publish a message")

	for d := range client.msgs {
		if corrId == d.CorrelationId {
			res = string(d.Body)
			failOnError(err, "Error msgs")
			break
		}
	}

	return res, nil
}

func (client *Client) Close() {
	client.channel.Close()
	client.connection.Close()
}

func New(ctx context.Context, address string) *Client {
	conn, err := amqp.Dial(address)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	return &Client{
		ctx:           ctx,
		connection:    conn,
		channel:       ch,
		responseQueue: q,
		msgs:          msgs,
	}
}
