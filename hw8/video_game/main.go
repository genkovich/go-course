package video_game

import (
	"context"
	"fmt"
	"time"
)

type Game struct {
	players      map[int]*Player
	questionPack QuestionPack
}

var answersChannel chan PlayerAnswer
var gameChannel chan []PlayerAnswer

func Run() {
	answersChannel = make(chan PlayerAnswer)
	gameChannel = make(chan []PlayerAnswer)
	game := Game{
		players:      generatePlayers(3),
		questionPack: GenerateQuestionPack(),
	}
	game.startGame()
}

func (g Game) startGame() {
	for _, question := range g.questionPack.questions {
		ctxWithQuestion := context.WithValue(context.Background(), "question", question)
		ctxTimeout, cancel := context.WithTimeout(ctxWithQuestion, time.Second*10)
		go collectAnswers(ctxTimeout)
		defer cancel()

		for _, player := range g.players {
			ctxWithPlayer := context.WithValue(ctxTimeout, "playerId", player.id)
			go showQuestion(ctxWithPlayer)
		}

		<-ctxTimeout.Done()
		fmt.Printf("Time is over. Correct answer is %s\n", question.variants[question.answerNumber])

		playersAnswers := <-gameChannel

		g.roundResults(playersAnswers, question)
	}

	g.gameResults()
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
	fmt.Println("Round over")
}

func (g Game) gameResults() {
	fmt.Println("Game over")

	for _, player := range g.players {
		fmt.Println(player.name, "has", player.points, "points")
	}
}

func showQuestion(ctx context.Context) {
	question := ctx.Value("question")
	fmt.Println("Question:", question, "Player:", ctx.Value("playerId"), "started")
	go generateAnswer(ctx)
}

func collectAnswers(ctx context.Context) {
	var roundResult []PlayerAnswer

	for {
		select {
		case <-ctx.Done():
			gameChannel <- roundResult
			return
		case answer := <-answersChannel:
			fmt.Println("Player", answer.playerId, "on time")
			roundResult = append(roundResult, answer)
		}
	}

}
