package config

import (
	"os"

	"github.com/torchiaf/Sensors/controller/utils"
)

type RabbitMQConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

type Module struct {
	Name       string `yaml:"name"`
	NodeName   string `yaml:"nodeName"`
	Type       string `yaml:"type"`
	RoutingKey string `yaml:"routingKey"`
	Devices    []Device
}

type Device struct {
	Name   string `yaml:"name"`
	Type   string `yaml:"type"`
	Config []DeviceConfig
}

type DeviceConfig struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type config struct {
	IsDev    bool
	RabbitMQ RabbitMQConfig
	Modules  []Module
}

func isDevEnv() bool {
	env := os.Getenv("DEV_ENV")

	return len(env) > 0
}

func initConfig() config {

	modules := utils.ParseYamlFile[[]Module]("/sensors/modules.yaml")

	c := config{
		IsDev: isDevEnv(),
		RabbitMQ: RabbitMQConfig{
			Host:     os.Getenv("RABBITMQ_CLUSTER_SERVICE_HOST"),
			Port:     os.Getenv("RABBITMQ_CLUSTER_SERVICE_PORT_AMQP"),
			Username: os.Getenv("RABBITMQ_USERNAME"),
			Password: os.Getenv("RABBITMQ_PASSWORD"),
		},
		Modules: modules,
	}

	return c
}

var Config = initConfig()
