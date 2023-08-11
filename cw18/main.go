package main

import (
	"encoding/json"
	"fmt"
	"github.com/mediocregopher/radix/v3"
	"time"
)

type Message struct {
	Id        int
	Body      string
	User      string
	Timestamp time.Time
}

type Chat struct {
	Connection *radix.Pool
	Messages   []Message
}

func (m Message) String() string {
	return fmt.Sprintf("[%s] %s: %s", m.Timestamp.Format("2006-01-02 15:4:05"), m.User, m.Body)
}

func main() {
	message := Message{Id: 1, Body: "Hello, World!", User: "test", Timestamp: time.Now()}
	secondMessage := Message{Id: 2, Body: "Goodbye, World!", User: "Second", Timestamp: time.Now()}

	p, err := radix.NewPool("tcp", "127.0.0.1:6379", 3)
	if err != nil {
		panic(err.Error())
	}

	chat := Chat{Connection: p}

	chat.sendMessage(message)
	chat.sendMessage(secondMessage)

	var messageString string

	err = p.Do(radix.Cmd(&messageString, "GET", fmt.Sprintf("message:%s:%d", message.User, message.Id)))
	if err != nil {
		panic(err.Error())
	}

	message = Message{}
	err = json.Unmarshal([]byte(messageString), &message)

	fmt.Println(message)

}

func (c *Chat) sendMessage(message Message) {
	messageJson, err := json.Marshal(message)

	err = c.Connection.Do(radix.Cmd(nil, "SET", fmt.Sprintf("message:%s:%d", message.User, message.Id), string(messageJson)))
	if err != nil {
		panic(err.Error())
	}
}
