package gift

import "course/gifts/human"

type Card struct{}

func (b Card) GetTitle() string {
	return "GiftCard"
}

func (b Card) IsGoodGift(h human.Human) bool {
	return true
}
