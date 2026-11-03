package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to create a channel: %s", err)

	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"hello-world",
		"consumer",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to register consumer: %s", err)
	}

	log.Println("Waiting for messages...")
	for msg := range msgs {
		log.Printf("Received: %s", msg.Body)
	}

}
