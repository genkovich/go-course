package video_game

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	id                    int
	name                  string
	points                int
	isActive              bool
	playerQuestionChannel chan Question
	answerChannel         chan PlayerAnswer
}

func generatePlayers(playersCount int, answerChannel chan PlayerAnswer) map[int]*Player {
	players := make(map[int]*Player, playersCount)
	for i := 1; i <= playersCount; i++ {
		players[i] = &Player{
			id:                    i,
			name:                  fmt.Sprintf("Player %d", i),
			points:                0,
			isActive:              true,
			playerQuestionChannel: make(chan Question),
			answerChannel:         answerChannel,
		}
	}
	return players
}

func (p *Player) AddPoint() {
	p.points++
}

func (p *Player) generateAnswer(question Question) {
	min := 1
	max := 4
	thinkingTime := time.Second * time.Duration(rand.Intn(12)+1)

	playerAnswer := PlayerAnswer{
		playerId: p.id,
		answer:   rand.Intn(max-min) + min,
		question: question,
	}

	select {
	case <-time.After(thinkingTime):
		p.answerChannel <- playerAnswer
	}

}

func (p *Player) showQuestion() {
	for {
		select {
		case question := <-p.playerQuestionChannel:
			fmt.Println("Question:", question, "Player:", p.id, "started")
			go p.generateAnswer(question)
		}
	}
}
