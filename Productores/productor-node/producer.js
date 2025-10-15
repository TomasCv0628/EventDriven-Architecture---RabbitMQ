// producer.js
const amqp = require('amqplib');
const os = require('os');

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
      console.log("[✓] Conectado a RabbitMQ. Enviando mensajes...");

      let i = 0;
      while (true) {
        i++;
        const msg = { producer: "node", host: os.hostname(), seq: i };
        ch.sendToQueue(QUEUE, Buffer.from(JSON.stringify(msg)), { persistent: true });
        console.log(" [x] Enviado:", msg);
        await new Promise(r => setTimeout(r, 2000));
      }
    } catch (err) {
      console.log("[!] RabbitMQ no está listo o se perdió la conexión. Reintentando en 5 segundos...");
      await new Promise(r => setTimeout(r, 5000));
    }
  }
}

main();
