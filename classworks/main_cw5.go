package main

import (
	"course/gifts/gift"
	"course/gifts/human"
	"fmt"
)

func main() {

	var humans []*human.Human
	for i := 0; i < 5; i++ {
		humans = append(humans, human.NewRandom())
		fmt.Println(humans[i])
		humans[i].UnpackGift()
	}

	var gifts = []human.Gift{
		gift.BikeHelmet{},
		gift.NewWineBottle(18),
		gift.Card{},
	}

	human.FindGoodGifts(humans, gifts)

	for _, h := range humans {
		h.UnpackGift()
	}

}
