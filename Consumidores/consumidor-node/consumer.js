const amqp = require('amqplib');
const RABBIT = process.env.RABBIT_HOST || 'rabbitmq';
const URL = `amqp://user:pass@${RABBIT}`;

(async () => {
  const conn = await amqp.connect(URL);
  const ch = await conn.createChannel();
  await ch.assertQueue('eventos', { durable: true });
  ch.prefetch(1);
  console.log(' [*] Node consumer waiting');
  ch.consume('eventos', (msg) => {
    if (msg !== null) {
      console.log(" [x] Node consumer got", msg.content.toString());
      ch.ack(msg);
    }
  });
})();
