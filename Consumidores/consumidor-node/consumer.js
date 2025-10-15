// consumer.js
const amqp = require('amqplib');

const RABBIT = process.env.RABBIT_HOST || 'rabbitmq';
const URL = `amqp://user:password@${RABBIT}`;
const QUEUE = 'eventos';

async function main() {
  while (true) {
    try {
      console.log("[*] Intentando conectar a RabbitMQ...");
      const conn = await amqp.connect(URL);
      const ch = await conn.createChannel();
      await ch.assertQueue(QUEUE, { durable: true });
      console.log("[✓] Conectado a RabbitMQ. Esperando mensajes...");

      ch.consume(QUEUE, msg => {
        if (msg !== null) {
          const content = JSON.parse(msg.content.toString());
          console.log(" [x] Node consumer got", content);
          ch.ack(msg);
        }
      });
    } catch (err) {
      console.log("[!] RabbitMQ no está listo o se perdió la conexión. Reintentando en 5 segundos...");
      await new Promise(r => setTimeout(r, 5000));
    }
  }
}

main();
