package rpc_client

import (
	"context"
	"fmt"

	"github.com/torchiaf/Sensors/rpc_client/config"
	"github.com/torchiaf/Sensors/rpc_client/models"
	"github.com/torchiaf/Sensors/rpc_client/rabbitmq"
)

func (client *RpcClient) Close() {
	client.rabbitmq.Close()
}

func (client *RpcClient) Read(module string, device string, args []string) (string, error) {
	queue := config.Config.Modules[module].RoutingKey

	return client.rabbitmq.Send(
		queue,
		models.Message{
			Device: device,
			Action: "read",
			Args:   args,
		},
	)
}

func (client *RpcClient) Write(module string, device string, args []string) (string, error) {
	queue := config.Config.Modules[module].RoutingKey

	return client.rabbitmq.Send(
		queue,
		models.Message{
			Device: device,
			Action: "write",
			Args:   args,
		},
	)
}

type RpcClient struct {
	rabbitmq *rabbitmq.Client
}

func New(ctx context.Context) *RpcClient {
	address := fmt.Sprintf("amqp://%s:%s@%s:%s/", config.Config.RabbitMQ.Username, config.Config.RabbitMQ.Password, config.Config.RabbitMQ.Host, config.Config.RabbitMQ.Port)

	rabbitmqClient := rabbitmq.New(ctx, address)

	return &RpcClient{
		rabbitmq: rabbitmqClient,
	}
}
