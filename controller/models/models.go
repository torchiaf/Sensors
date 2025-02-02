package models

type RabbitMQ struct {
	Host     string
	Port     string
	Username string
	Password string
}

type Release struct {
	Name      string
	Namespace string
	Group     string
	Version   string
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
	Release  Release
	RabbitMQ RabbitMQ
	Modules  []Module
}

type Message struct {
	Device string   `json:"device"`
	Action string   `json:"action"`
	Args   []string `json:"args"`
}

type Circuit struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	} `yaml:"metadata"`
	Spec struct {
		Id string `yaml:"id"`
	} `yaml:"spec"`
}
