package hw16_2

import "fmt"

type Player interface {
	Listen(subject any)
	GetID() string
	Notify(room *GameRoom, changes string)
}

type GamePlayer struct {
	Id string
}

func (gp *GamePlayer) GetID() string {
	return gp.Id
}

func (gp *GamePlayer) Listen(subject any) {
	fmt.Printf("ID: %s player notified with: %s \n", gp.Id, subject)
}

func (gp *GamePlayer) Notify(room *GameRoom, changes string) {
	room.NotifyPlayers(changes)
}
