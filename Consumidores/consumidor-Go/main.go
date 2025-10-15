package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	rabbit := os.Getenv("RABBIT_HOST")
	if rabbit == "" {
		rabbit = "rabbitmq"
	}

	for {
		conn, err := amqp.Dial(fmt.Sprintf("amqp://user:password@%s:5672/", rabbit))
		if err != nil {
			log.Println("RabbitMQ no está listo. Reintentando en 5 segundos...")
			time.Sleep(5 * time.Second)
			continue
		}
		ch, _ := conn.Channel()
		msgs, _ := ch.Consume("eventos", "", false, false, false, false, nil)
		log.Println("[✓] Conectado a RabbitMQ. Esperando mensajes...")

		for d := range msgs {
			var msg map[string]interface{}
			json.Unmarshal(d.Body, &msg)
			log.Printf(" [x] Go consumer got %v", msg)
			d.Ack(false)
		}

		ch.Close()
		conn.Close()
		log.Println("Conexión perdida. Reintentando en 5 segundos...")
		time.Sleep(5 * time.Second)
	}
}
