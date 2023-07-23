package main

import (
	"course/hw16/agent"
	"course/hw16_2"
	"fmt"
)

type FirstSub struct {
	id string
}

func (t *FirstSub) Consume(event any) {
	fmt.Printf("FirstSub: %v \n", event)
}

func (t *FirstSub) Id() string {
	return t.id
}

type SecondSub struct {
	id string
}

func (t *SecondSub) Consume(event any) {
	fmt.Printf("SecondSub get event: %v \n", event)
}

func (t *SecondSub) Id() string {
	return t.id
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

	first := FirstSub{id: "1"}
	a.AddSub(&first)

	second := SecondSub{id: "2"}
	a.AddSub(&second)

	a.Watch()
}
