package hw16_2

type Game interface {
	Add(p Player)
	Remove(o Player)
	NotifyPlayers(subject any)
}

type GameRoom struct {
	players map[string]Player
}

func (g *GameRoom) Add(p Player) {
	g.players[p.GetID()] = p
}

func (g *GameRoom) Remove(p Player) {
	delete(g.players, p.GetID())
}

func (g *GameRoom) NotifyPlayers(subject any) {
	for _, p := range g.players {
		p.Listen(subject)
	}
}

func NewGame() *GameRoom {
	return &GameRoom{
		players: make(map[string]Player),
	}
}
