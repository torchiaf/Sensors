import os
import builtins
import pika

# Redefine print to show the output when running in containers
def print(*args):
    builtins.print(*args, sep=' ', end='\n', file=None, flush=True)

# RabbitMQ configs
host = os.environ.get('RABBITMQ_CLUSTER_SERVICE_HOST')
port = os.environ.get('RABBITMQ_CLUSTER_SERVICE_PORT_AMQP')
username = os.environ.get('RABBITMQ_USERNAME')
password = os.environ.get('RABBITMQ_PASSWORD')

credentials = pika.PlainCredentials(username, password)
connectionParams = pika.ConnectionParameters(host,port,'/',credentials)