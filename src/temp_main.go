/*package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)


ticker := time.NewTicker(2 * time.Second)
quit := make(chan struct{})


func main() {
	go server()
	go client()
}

func server() {
	conn, ch, q := getQueue()
	defer conn.Close()
	defer ch.Close()

	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello RabbitMQ"),
	}

	ch.Publish("", q.Name, false, false, msg)
}

func client() {
	conn, ch, q := getQueue()
	defer conn.Close()
	defer ch.Close()

	msgs, err := ch.Consume(q.Name, // queue,
		"",    //consumer string,
		true,  //autoAck bool,
		false, //exclusive bool,
		false, //noLocal bool,
		false, //noWait bool,
		nil)   //args Table)

	failOnError(err, "Failed to Register a consumer")

	for msg := range msgs {
		log.Println("received message with message: %s", msg.Body)
	}
}

func getQueue() (*amqp.Connection, *amqp.Channel, *amqp.Queue) {
	conn, err := amqp.Dial("amqp://guest@localhost:5672")
	failOnError(err, "Failes to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to poen a Channel")

	q, err := ch.QueueDeclare("hello",
		false, //durable bool,
		false, //autoDelete bool,
		false, //exclusive bool,
		false, //noWait bool,
		nil)   //args amqp.Table
	failOnError(err, "Failed to declare queue")

	return conn, ch, &q

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

*/