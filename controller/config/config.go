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
		Release: models.Release{
			Name:      utils.IfNull(os.Getenv("APP_NAME"), "sensors"),
			Namespace: utils.IfNull(os.Getenv("APP_NAMESPACE"), "sensors"),
			Group:     utils.IfNull(os.Getenv("APP_GROUP"), "sensors.io"),
			Version:   utils.IfNull(os.Getenv("APP_VERSION"), "v1"),
		},
		RabbitMQ: models.RabbitMQ{
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
