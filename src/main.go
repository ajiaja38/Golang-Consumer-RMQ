package main

import (
	"go-consumer-rmq/src/error"
	"go-consumer-rmq/src/service"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	error.FailOnError(err, "Failed Load Env")

	service.Consumer()
}
