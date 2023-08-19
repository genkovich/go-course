package producer

import (
	"course/hw20/pkg/orange"
	"course/hw20/pkg/rabbitmq"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

func test() {

}

func Produce() {
	connection := rabbitmq.NewChannel()
	channel := connection.GetChannel()

	queue, err := channel.QueueDeclare(
		"oranges",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	jsonBody, err := json.Marshal(orange.NewOrange())

	if err != nil {
		panic(err)
	}

	err = channel.Publish(
		"",
		"oranges",
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/octet-stream",
			Body:         jsonBody,
			Priority:     5,
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Queue status:", queue)
	fmt.Println("Successfully published message")
}
