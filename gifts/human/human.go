package human

import (
	"fmt"
	"math/rand"
)

type Human struct {
	name  string
	age   int
	hobby string
	gift  Gift
}

func (h Human) Name() string {
	if h.name == "" {
		return "Anonymous"
	}
	return h.name
}

func (h Human) Age() int {
	return h.age
}

func (h Human) Hobby() string {
	return h.hobby
}

func (h Human) UnpackGift() {
	if h.gift == nil {
		fmt.Printf("Nothing for %s\n", h.String())
	} else {
		fmt.Printf("%s received %s\n", h.String(), h.gift.GetTitle())
	}
}

const (
	HobbyNone    = ""
	HobbyCycling = "cycling"
	HobbyWine    = "wine"
)

func getRandomHobby(age int) string {
	switch rand.Intn(3) {
	case 1:
		return HobbyCycling
	case 2:
		if age > 18 {
			return HobbyWine
		}
		fallthrough
	default:
		return HobbyNone
	}
}

func (h Human) String() string {
	return fmt.Sprintf("%s: %v years old, hobby: %s", h.Name(), h.Age(), h.Hobby())
}

func NewRandom() *Human {

	randAge := rand.Intn(25)
	return &Human{
		age:   randAge,
		hobby: getRandomHobby(randAge),
	}
}

type Gift interface {
	GetTitle() string
	IsGoodGift(Human) bool
}

func FindGoodGifts(humans []*Human, gifts []Gift) {
	for _, h := range humans {
		for _, g := range gifts {
			if g.IsGoodGift(*h) {
				h.gift = g
				break
			}
		}
	}
}
