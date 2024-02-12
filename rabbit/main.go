package rabbit

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/41x3n/TeleUtil/bootstrap"
	"github.com/41x3n/TeleUtil/processor"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	q             amqp.Queue
	queueDeclared bool
)

const FileQueue = "file_queue"

func GetQueue(app *bootstrap.Application) (*amqp.Queue, error) {
	if !queueDeclared {
		ch := app.RabbitMQ.Ch

		var err error
		q, err = ch.QueueDeclare(
			FileQueue,
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to declare a queue: %v", err)
		}
		queueDeclared = true
	}

	return &q, nil
}

func PublishMessage(app *bootstrap.Application, message string) error {
	ch := app.RabbitMQ.Ch

	q, err := GetQueue(app)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(
		ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)

	if err != nil {
		return fmt.Errorf("failed to publish a message: %v", err)
	}

	return nil
}

func ConsumeMessages(app *bootstrap.Application) {
	ch := app.RabbitMQ.Ch

	q, err := GetQueue(app)
	if err != nil {
		log.Fatalf("failed to get queue: %v", err)
		return
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("failed to register a consumer: %v", err)
	}

	sem := make(chan struct{}, 10) // Create a semaphore channel with a capacity of 10 and uses struct as the type of the channel because we don't need to send any data, just a signal and struct is the smallest data type in Go

	go func() {
		for d := range msgs {
			sem <- struct{}{} // Send to the semaphore channel
			go func(delivery amqp.Delivery) {
				err := processor.ProcessImage(delivery)
				if err != nil {
					log.Printf("Failed to process image: %s", err)
				}
				<-sem // Receive from the semaphore channel
			}(d)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
}
