import os
import pika

host = os.environ.get('RABBITMQ_CLUSTER_SERVICE_HOST')
port = os.environ.get('RABBITMQ_CLUSTER_SERVICE_PORT_AMQP')
username = os.environ.get('RABBITMQ_USERNAME')
password = os.environ.get('RABBITMQ_PASSWORD')

credentials = pika.PlainCredentials(username, password)
connectionParams = pika.ConnectionParameters(host,port,'/',credentials)