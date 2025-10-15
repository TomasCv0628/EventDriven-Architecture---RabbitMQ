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
		ch.QueueDeclare("eventos", true, false, false, false, nil)
		log.Println("[✓] Conectado a RabbitMQ. Enviando mensajes...")

		i := 0
		for {
			i++
			body, _ := json.Marshal(map[string]interface{}{
				"producer": "go",
				"seq":      i,
			})
			err = ch.Publish("", "eventos", false, false, amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "application/json",
				Body:         body,
			})
			if err != nil {
				log.Println("Error al enviar mensaje:", err)
				break
			}
			log.Printf(" [x] Enviado %s", string(body))
			time.Sleep(2 * time.Second)
		}

		ch.Close()
		conn.Close()
		time.Sleep(5 * time.Second)
	}
}
