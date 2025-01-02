package models

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

type Config struct {
	IsDev    bool
	RabbitMQ RabbitMQConfig
	Modules  []Module
}

type Message struct {
	Device string `json:"device"`
	// Args   map[string]interface{} `json:"args"`
}
