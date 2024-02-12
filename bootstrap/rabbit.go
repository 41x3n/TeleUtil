package bootstrap

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
}

func CreateRabbitMQConnection(env *Env) *amqp.Connection {
	url := env.RabbitMQURL

	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")

	log.Println("Connection to RabbitMQ established.")
	return conn
}

func CloseRabbitMQConnection(conn *amqp.Connection) {
	conn.Close()
}

func CreateChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	return ch
}

func CloseRabbitMQChannel(ch *amqp.Channel) {
	ch.Close()
}

func NewRabbitMQ(env *Env) *RabbitMQ {
	conn := CreateRabbitMQConnection(env)
	ch := CreateChannel(conn)

	return &RabbitMQ{
		Conn: conn,
		Ch:   ch,
	}

}
