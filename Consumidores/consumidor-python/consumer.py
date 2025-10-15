# consumer.py
import pika, json, os
RABBIT = os.getenv("RABBIT_HOST", "rabbitmq")
CREDENTIALS = pika.PlainCredentials('user','pass')
params = pika.ConnectionParameters(host=RABBIT, credentials=CREDENTIALS)

def callback(ch, method, properties, body):
    msg = json.loads(body)
    print(" [x] Python consumer got", msg)
    ch.basic_ack(delivery_tag = method.delivery_tag)

def main():
    connection = pika.BlockingConnection(params)
    channel = connection.channel()
    channel.queue_declare(queue='eventos', durable=True)
    channel.basic_qos(prefetch_count=1)
    channel.basic_consume(queue='eventos', on_message_callback=callback)
    print(' [*] Waiting for messages. To exit press CTRL+C')
    channel.start_consuming()

if __name__ == "__main__":
    main()
