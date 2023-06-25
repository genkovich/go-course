package video_game

import "fmt"

type Player struct {
	id     int
	name   string
	points int
}

func generatePlayers(playersCount int) map[int]*Player {
	players := make(map[int]*Player, playersCount)
	for i := 1; i <= playersCount; i++ {
		players[i] = &Player{
			id:     i,
			name:   fmt.Sprintf("Player %d", i),
			points: 0,
		}
	}
	return players
}

func (p *Player) AddPoint() {
	p.points++
}
