package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

type Channel struct {
	channel *amqp.Channel
}

func NewChannel() *Channel {
	fmt.Println("RabbitMQ in Golang: Getting started tutorial")

	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to RabbitMQ instance")

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	return &Channel{channel}
}

func (c *Channel) GetChannel() *amqp.Channel {
	return c.channel
}
