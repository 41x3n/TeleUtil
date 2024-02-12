package processor

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ProcessImage(delivery amqp.Delivery) error {
	// Process the image here...
	log.Printf("Processing image: %s", delivery.Body)

	// If the task is successful, acknowledge the message
	if err := delivery.Ack(false); err != nil {
		log.Printf("Failed to acknowledge message: %s", err)
		return err
	}

	log.Printf("Acknowledged message")
	return nil
}
