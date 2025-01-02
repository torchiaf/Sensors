package config

import (
	"os"

	"github.com/torchiaf/Sensors/controller/models"
	"github.com/torchiaf/Sensors/controller/utils"
)

func isDevEnv() bool {
	env := os.Getenv("DEV_ENV")

	return len(env) > 0
}

func initConfig() models.Config {

	modules := utils.ParseYamlFile[[]models.Module]("/sensors/modules.yaml")

	c := models.Config{
		IsDev: isDevEnv(),
		RabbitMQ: models.RabbitMQConfig{
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
