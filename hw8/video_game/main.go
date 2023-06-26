package video_game

import (
	"fmt"
	"sync"
	"time"
)

type Game struct {
	players            map[int]*Player
	questionPack       QuestionPack
	roundResultChannel chan []PlayerAnswer
	answersChannel     chan PlayerAnswer
}

func Run() {
	answerChannel := make(chan PlayerAnswer)
	game := Game{
		players:            generatePlayers(3, answerChannel),
		questionPack:       GenerateQuestionPack(),
		roundResultChannel: make(chan []PlayerAnswer),
		answersChannel:     answerChannel,
	}
	game.startGame()
}

func (g Game) startGame() {
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
	g.gameResults()
}

func (g Game) sendQuestions(wg *sync.WaitGroup) {
	for _, question := range g.questionPack.questions {
		fmt.Println("Question:", question)
		for _, player := range g.players {
			player.playerQuestionChannel <- question
		}

		roundResult := <-g.roundResultChannel

		g.roundResults(roundResult, question)
	}
	wg.Done()
}

func (g Game) timeIsOver(roundResult []PlayerAnswer) {

}

func (g Game) collectAnswers(ticker *time.Ticker) {
	var roundResult []PlayerAnswer

	for {
		select {
		case <-ticker.C:
			fmt.Println("Round over")
			g.roundResultChannel <- roundResult
			roundResult = nil
		case playerAnswer := <-g.answersChannel:
			roundResult = append(roundResult, playerAnswer)
		}
	}
}

func (g Game) roundResults(playersAnswers []PlayerAnswer, question Question) {
	for _, playerAnswer := range playersAnswers {
		if playerAnswer.answer == question.answerNumber {
			fmt.Println("Player", playerAnswer.playerId, "answered correctly")
			g.players[playerAnswer.playerId].AddPoint()
			fmt.Println("Player", playerAnswer.playerId, "has", g.players[playerAnswer.playerId].points, "points")
		} else {
			fmt.Println("Player", playerAnswer.playerId, "answered incorrectly")
		}
	}
}

func (g Game) gameResults() {
	fmt.Println("Game over")
	for _, player := range g.players {
		fmt.Println(player.name, "has", player.points, "points")
	}
}
