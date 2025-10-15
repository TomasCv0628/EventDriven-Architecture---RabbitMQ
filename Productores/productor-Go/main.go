package main
import (
  "github.com/streadway/amqp"
  "log"
  "os"
  "time"
  "encoding/json"
)

func failOnError(err error, msg string) {
  if err != nil {
    log.Fatalf("%s: %s", msg, err)
  }
}

func main() {
  rabbit := os.Getenv("RABBIT_HOST")
  if rabbit=="" { rabbit = "rabbitmq" }
  conn, err := amqp.Dial("amqp://user:pass@" + rabbit + ":5672/")
  failOnError(err, "Failed to connect")
  defer conn.Close()
  ch, err := conn.Channel()
  failOnError(err, "Failed to open channel")
  defer ch.Close()
  _, err = ch.QueueDeclare("eventos", true, false, false, false, nil)
  failOnError(err, "Queue Declare failed")
  for i:=1; i<=20; i++ {
    body, _ := json.Marshal(map[string]interface{}{"producer":"go","seq":i})
    err = ch.Publish("", "eventos", false, false, amqp.Publishing{
      DeliveryMode: amqp.Persistent,
      ContentType: "application/json",
      Body: body,
    })
    log.Printf(" [x] Sent %s", body)
    time.Sleep(500 * time.Millisecond)
  }
}
