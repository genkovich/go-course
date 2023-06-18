package gift

import "course/gifts/human"

type BikeHelmet struct{}

func (b BikeHelmet) GetTitle() string {
	return "Helmet"
}

func (b BikeHelmet) IsGoodGift(h human.Human) bool {
	return h.Hobby() == human.HobbyCycling
}
