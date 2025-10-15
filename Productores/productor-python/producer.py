# producer.py
import pika, json, time, os, socket
RABBIT = os.getenv("RABBIT_HOST", "rabbitmq")
CREDENTIALS = pika.PlainCredentials('user','password')
params = pika.ConnectionParameters(host=RABBIT, credentials=CREDENTIALS)

def main():
    connection = pika.BlockingConnection(params)
    channel = connection.channel()
    channel.queue_declare(queue='eventos', durable=True)

    hostname = socket.gethostname()
    for i in range(1, 21):
        msg = {"producer":"python", "host": hostname, "seq": i}
        channel.basic_publish(
            exchange='',
            routing_key='eventos',
            body=json.dumps(msg),
            properties=pika.BasicProperties(delivery_mode=2)
        )
        print(" [x] Sent", msg)
        time.sleep(0.5)
    connection.close()

if __name__ == "__main__":
    main()
