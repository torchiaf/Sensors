import os
import pika

host = os.environ.get('RABBITMQ_CLUSTER_SERVICE_HOST')
port = os.environ.get('RABBITMQ_CLUSTER_SERVICE_PORT')
username = os.environ.get('RMQ_USERNAME')
password = os.environ.get('RMQ_PASSWORD')

credentials = pika.PlainCredentials(username, password)
connectionParams = pika.ConnectionParameters(host,port,'/',credentials)