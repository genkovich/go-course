package video_game

func Run() {
	answerChannel := make(chan PlayerAnswer)
	game := Game{
		players:            generatePlayers(3, answerChannel),
		questionPack:       GenerateQuestionPack(),
		roundResultChannel: make(chan RoundResult),
		answersChannel:     answerChannel,
	}
	game.startGame()
}
