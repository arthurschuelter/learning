package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672")
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %s", err)

	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to create a channel: %s", err)

	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello-world",
		true,
		false,
		false,
		false,
		nil,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for {
		body := "Hello World! Sent at: " + time.Now().Format("2006-01-02 15:04:05")
		err = ch.PublishWithContext(ctx,
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			},
		)
		if err != nil {
			log.Fatalf("Failed to publish message: %s", err)
		}

		log.Printf(" [x] Sent %s\n", body)
		time.Sleep(3 * time.Second)
	}
}
