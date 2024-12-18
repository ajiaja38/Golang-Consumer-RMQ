package config

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func RmqConnection() (*amqp.Channel, error) {
	RMQ_URI := os.Getenv("RMQ_URI")

	conn, err := amqp.Dial(RMQ_URI)

	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()

	if err != nil {
		conn.Close()
		return nil, err
	}

	log.Print("Success Connected To RabbitMq")
	return ch, nil
}
