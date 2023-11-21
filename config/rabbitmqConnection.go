package config

import (
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

var channel *amqp.Channel
var connection *amqp.Connection

func SetupRabbitMQ() {
	URL := os.Getenv("RabbitMQ_URL")
	if URL == "" {
		failOnError(nil, "Failed to get RabbitMQ URL")
	}

	fmt.Println(URL)

	connection, err := amqp.Dial(URL) //Establecer una conexion con el servidor de rabbitmq
	failOnError(err, "Failed to connect to RabbitMQ")

	channel, err = connection.Channel() //Sesion o instancia de comunicacion con el servidor de rabbitmq
	failOnError(err, "Failed to open a channel")

}

func GetChannel() *amqp.Channel {
	return channel
}

func CloseRabbitMQ() {
	channel.Close()
	connection.Close()
}
