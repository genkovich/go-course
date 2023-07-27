package main

import (
	"course/hw16/agent"
	"course/hw16_2"
	"fmt"
)

type FirstSub struct {
	id        string
	eventChan chan any
	stop      chan struct{}
}

func (f *FirstSub) Consume() {
	for {
		select {
		case event := <-f.eventChan:
			fmt.Printf("FirstSub: %v \n", event)
		case <-f.stop:
			f.stop <- struct{}{}
			return
		}
	}
}

func (f *FirstSub) Publish(event any) {
	f.eventChan <- event
}

func (f *FirstSub) Id() string {
	return f.id
}

func (f *FirstSub) Stop() {
	f.stop <- struct{}{}
	<-f.stop
}

type SecondSub struct {
	id        string
	eventChan chan any
	stop      chan struct{}
}

func (s *SecondSub) Consume() {
	for {
		select {
		case event := <-s.eventChan:
			fmt.Printf("SecondSub get event: %v \n", event)
		case <-s.stop:
			s.stop <- struct{}{}
			return
		}
	}
}

func (s *SecondSub) Publish(event any) {
	s.eventChan <- event
}

func (s *SecondSub) Id() string {
	return s.id
}

func (s *SecondSub) Stop() {
	s.stop <- struct{}{}
	<-s.stop
}

func main() {
	game := hw16_2.NewGame()
	player1 := hw16_2.GamePlayer{Id: "1"}
	player2 := hw16_2.GamePlayer{Id: "2"}
	game.Add(&player1)
	game.Add(&player2)
	player1.Notify(game, "some changes")
	player2.Notify(game, "another changes")
	player3 := hw16_2.GamePlayer{Id: "3"}
	game.Add(&player3)
	player3.Notify(game, "third changes")

	////////////
	a := agent.NewAgent()

	first := FirstSub{
		id:        "1",
		eventChan: make(chan any),
	}

	a.AddSub(&first)

	second := SecondSub{
		id:        "2",
		eventChan: make(chan any),
	}

	a.AddSub(&second)

	a.Watch()
}
