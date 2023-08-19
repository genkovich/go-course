package consumer

import (
	"course/hw20/pkg/orange"
	"course/hw20/pkg/rabbitmq"
	"encoding/json"
	"fmt"
	"time"
)

func Consume() {
	connection := rabbitmq.NewChannel()
	channel := connection.GetChannel()

	msgs, err := channel.Consume(
		"oranges",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	counter := orange.NewSizeCounter()

	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		for {
			select {
			case msg := <-msgs:
				var o orange.Orange
				if err = json.Unmarshal(msg.Body, &o); err != nil {
					fmt.Println("Error unmarshalling message:", err)
					continue
				}
				fmt.Println("Received orange:", o)
				counter.ClassifyAndCount(o.Size)
			case <-ticker.C:
				fmt.Printf("Counts - Small: %d, Medium: %d, Large: %d\n", counter.Small, counter.Medium, counter.Large)
			}
		}
	}()

	fmt.Println("Waiting for messages...")
	select {}
}
