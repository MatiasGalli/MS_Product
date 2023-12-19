package main

import (
	"fmt"
	"log"

	"github.com/MatiasGalli/MS_Product/config"
	"github.com/MatiasGalli/MS_Product/internal"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func getChannel() *amqp.Channel {
	channel := config.GetChannel()
	if channel == nil {
		log.Panic("Failed to get channel")
	}
	return channel
}

func declareQueue(channel *amqp.Channel) amqp.Queue {
	queue, err := channel.QueueDeclare(
		"product_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")
	return queue
}

func setQoS(channel *amqp.Channel) {
	err := channel.Qos(
		1,
		0,
		false,
	)
	failOnError(err, "Failed to set QoS")
}

func registerConsumer(channel *amqp.Channel, queue amqp.Queue) <-chan amqp.Delivery {
	msgs, err := channel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")
	return msgs
}

func main() {
	fmt.Println("Product MS starting...")

	godotenv.Load()
	fmt.Println("Loaded env variables...")

	config.SetupDatabase()
	fmt.Println("Database connection configured...")

	config.SetupRabbitMQ()
	fmt.Println("RabbitMQ Connection configured...")

	channel := getChannel()
	queue := declareQueue(channel)
	setQoS(channel)
	msgs := registerConsumer(channel, queue)

	var forever chan struct{}

	go func() {
		for d := range msgs {
			internal.Handler(d, channel)
		}
	}()
	fmt.Println("Product MS started...")
	log.Print(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
