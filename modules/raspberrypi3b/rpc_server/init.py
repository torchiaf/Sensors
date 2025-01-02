import subprocess
import ast
import pika
from config import module, connectionParams, print

connection = pika.BlockingConnection(connectionParams)

channel = connection.channel()

channel.queue_declare(queue=module.routingKey)

def on_request(ch, method, props, body):
    
    dict_str = body.decode("UTF-8")
    params = ast.literal_eval(dict_str)
    print(repr(params))

    res = subprocess.run(['./{}'.format(params["device"])], stdout=subprocess.PIPE, text=True)
    response = res.stdout

    ch.basic_publish(
        exchange='',
        routing_key=props.reply_to,
        properties=pika.BasicProperties(correlation_id = props.correlation_id),
        body=str(response)
    )
    ch.basic_ack(delivery_tag=method.delivery_tag)

channel.basic_qos(prefetch_count=1)
channel.basic_consume(queue=module.routingKey, on_message_callback=on_request)

print(" [x] Awaiting RPC requests")
channel.start_consuming()