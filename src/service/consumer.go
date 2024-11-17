package service

import (
	"encoding/json"
	"go-consumer-rmq/src/config"
	"go-consumer-rmq/src/error"
	"go-consumer-rmq/src/model"
	"log"
)

func Consumer() {
	ch, err := config.RmqConnection()

	if err != nil {
		error.FailOnError(err, "Failed to Connect To RMQ")
		return
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"monitoring",
		true,
		false,
		false,
		false,
		nil,
	)

	error.FailOnError(err, "Failed to declare Queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	error.FailOnError(err, "Failed to register consumer")

	var forever chan struct{}
	var user model.UserDao

	go func() {
		for d := range msgs {
			err := json.Unmarshal(d.Body, &user)

			error.FailOnError(err, "Failed to Decode Struct")

			log.Print(user)
		}
	}()

	log.Printf("[*] Waiting for messages, To exit press CTRL+C")

	<-forever
}
