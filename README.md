# Sensors

## Sensors handles IoT devices mounted on Raspberry PI modules in a Kubernetes cluster.

### Install RabbitMq operator

    helm install rabbitmq-operator charts/rabbitmq-operator
    helm install rabbitmq-cluster charts/rabbitmq-cluster --values charts/settings.yaml

### Install Sensors

    helm install sensors charts/sensors --values charts/settings.yaml
