# Sensors

## Sensors handles IoT devices mounted on Raspberry PI modules in a Kubernetes cluster.

### Create a k3s cluster

https://docs.k3s.io/quick-start#install-script

- Create a single node cluster

    ```
    curl -sfL https://get.k3s.io | sh -
    ```

- Add raspberry agent nodes
    
    ```
    curl -sfL https://get.k3s.io | K3S_URL=https://myserver:6443 K3S_TOKEN=mynodetoken sh -
    ```

- Important! Add `sensors.role: worker` label to each raspberry nodes

### Create sensors namespace

    kubectl create namespace sensors

### Install RabbitMQ

    helm install rabbitmq-operator charts/rabbitmq-operator -n sensors
    helm install rabbitmq-cluster charts/rabbitmq-cluster --values charts/settings.yaml -n sensors

- Wait for RabbitMQ deployments to be ready

### Install Sensors

    helm install sensors charts/sensors --values charts/settings.yaml -n sensors
