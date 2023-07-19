package video_game

import (
	"fmt"
	"sync"
	"time"
)

type Game struct {
	players            map[int]*Player
	questionPack       QuestionPack
	roundResultChannel chan RoundResult
	answersChannel     chan PlayerAnswer
}

func (g *Game) startGame() {
	fmt.Println("Game Started")

	ticker := time.NewTicker(10 * time.Second)
	go g.collectAnswers(ticker)

	for _, player := range g.players {
		go player.showQuestion()
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go g.sendQuestions(&wg)

	wg.Wait()
	g.printGameResults()
}

func (g *Game) sendQuestions(wg *sync.WaitGroup) {
	for _, question := range g.questionPack.questions {
		fmt.Println("Question:", question)
		for _, player := range g.players {
			if player.isActive {
				player.playerQuestionChannel <- question
			}
		}

		roundResult := <-g.roundResultChannel

		roundResult.countRoundResults(g, question)
	}
	wg.Done()
}

func (g *Game) collectAnswers(ticker *time.Ticker) {
	roundResult := CreateRoundResult()

	for {
		select {
		case <-ticker.C:
			g.timeIsOver(roundResult)
			roundResult.NewRound()
		case playerAnswer := <-g.answersChannel:
			if g.players[playerAnswer.playerId].isActive {
				fmt.Println("Player", playerAnswer.playerId, "answered", playerAnswer.question.variants[playerAnswer.answer], "on question", playerAnswer.question)
				roundResult.AddAnswer(playerAnswer)
			}
		}
	}
}

func (g *Game) timeIsOver(roundResult RoundResult) {
	fmt.Println("Round over")
	for _, player := range g.players {
		if !roundResult.isPlayerAnswered(player.id) {
			fmt.Println("Player", player.id, "didn't answer and left the game")
			player.isActive = false
		}
	}

	g.roundResultChannel <- roundResult
}

func (g *Game) IsActivePlayer(playerId int) bool {
	return g.players[playerId].isActive
}

func (g *Game) AddPointToPlayer(playerId int) {
	g.players[playerId].AddPoint()
}

func (g *Game) printGameResults() {
	fmt.Println("Game over")
	for _, player := range g.players {
		fmt.Println(player.name, "has", player.points, "points")
	}
}
