package video_game

import "fmt"

type RoundResult struct {
	answers []PlayerAnswer
}

func CreateRoundResult() RoundResult {
	return RoundResult{}
}

func (r *RoundResult) NewRound() {
	r.answers = []PlayerAnswer{}
}

func (r *RoundResult) isPlayerAnswered(playerId int) bool {
	for _, playerAnswer := range r.answers {
		if playerAnswer.playerId == playerId {
			return true
		}
	}
	return false
}

func (r *RoundResult) AddAnswer(playerAnswer PlayerAnswer) {
	r.answers = append(r.answers, playerAnswer)
}

func (r *RoundResult) countRoundResults(g *Game, question Question) {
	for _, playerAnswer := range r.answers {
		if !g.IsActivePlayer(playerAnswer.playerId) {
			continue
		}
		if playerAnswer.answer == question.answerNumber {
			fmt.Println("Player", playerAnswer.playerId, "answered correctly")
			g.AddPointToPlayer(playerAnswer.playerId)
			fmt.Println("Player", playerAnswer.playerId, "has", g.players[playerAnswer.playerId].points, "points")
		} else {
			fmt.Println("Player", playerAnswer.playerId, "answered incorrectly")
		}
	}
}
