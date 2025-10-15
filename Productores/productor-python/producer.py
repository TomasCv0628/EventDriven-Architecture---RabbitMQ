# producer.py
import pika, json, time, os, socket

RABBIT = os.getenv("RABBIT_HOST", "rabbitmq")
CREDENTIALS = pika.PlainCredentials('user','password')
params = pika.ConnectionParameters(host=RABBIT, credentials=CREDENTIALS)

def main():
    while True:
        try:
            connection = pika.BlockingConnection(params)
            channel = connection.channel()
            channel.queue_declare(queue='eventos', durable=True)

            hostname = socket.gethostname()
            i = 0
            print("[*] Conectado a RabbitMQ. Enviando mensajes...")
            while True:
                i += 1
                msg = {"producer": "python", "host": hostname, "seq": i}
                channel.basic_publish(
                    exchange='',
                    routing_key='eventos',
                    body=json.dumps(msg),
                    properties=pika.BasicProperties(delivery_mode=2)
                )
                print(" [x] Enviado:", msg)
                time.sleep(2)
        except pika.exceptions.AMQPConnectionError:
            print("RabbitMQ no está listo. Reintentando en 5 segundos...")
            time.sleep(5)
        except Exception as e:
            print("Error inesperado:", e)
            time.sleep(5)

if __name__ == "__main__":
    main()
