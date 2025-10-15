package main
import (
  "github.com/streadway/amqp"
  "log"
  "os"
  "encoding/json"
)

func failOnError(err error, msg string) { if err!=nil { log.Fatalf("%s: %s", msg, err)} }

func main(){
  rabbit := os.Getenv("RABBIT_HOST"); if rabbit=="" { rabbit="rabbitmq" }
  conn, err := amqp.Dial("amqp://user:password@"+rabbit+":5672/")
  failOnError(err, "connect")
  ch, _ := conn.Channel()
  defer ch.Close()
  msgs, _ := ch.Consume("eventos", "", false, false, false, false, nil)
  forever := make(chan bool)
  go func(){
    for d := range msgs {
      var m map[string]interface{}
      json.Unmarshal(d.Body, &m)
      log.Printf(" [x] Go consumer got %v", m)
      d.Ack(false)
    }
  }()
  log.Printf(" [*] Waiting")
  <-forever
}
