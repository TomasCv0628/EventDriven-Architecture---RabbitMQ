// producer.js
const amqp = require('amqplib');
const RABBIT = process.env.RABBIT_HOST || 'rabbitmq';
const URL = `amqp://user:pass@${RABBIT}`;

(async () => {
  const conn = await amqp.connect(URL);
  const ch = await conn.createChannel();
  await ch.assertQueue('eventos', { durable: true });

  for (let i=1; i<=20; i++){
    const msg = { producer: 'node', seq: i, host: require('os').hostname() };
    ch.sendToQueue('eventos', Buffer.from(JSON.stringify(msg)), { persistent:true });
    console.log(' [x] Sent', msg);
    await new Promise(r => setTimeout(r, 500));
  }
  await ch.close();
  await conn.close();
})();
