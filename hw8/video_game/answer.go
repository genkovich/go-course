package video_game

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type PlayerAnswer struct {
	playerId int
	answer   int
}

func generateAnswer(ctx context.Context) {
	min := 1
	max := 4
	thinkingTime := time.Second * time.Duration(rand.Intn(12-2)+2)

	time.Sleep(thinkingTime)
	playerAnswer := PlayerAnswer{
		playerId: ctx.Value("playerId").(int),
		answer:   rand.Intn(max-min) + min,
	}

	select {
	case <-ctx.Done():
		return
	default:
		question := ctx.Value("question").(Question)
		fmt.Println("Player", playerAnswer.playerId, "answered", question.variants[playerAnswer.answer], "to question", question)
		answersChannel <- playerAnswer
	}

}
