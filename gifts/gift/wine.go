package gift

import "course/gifts/human"

type WineBottle struct {
	minAge int
}

func NewWineBottle(minAge int) WineBottle {
	return WineBottle{
		minAge: minAge,
	}
}

func (wb WineBottle) GetTitle() string {
	return "Wine bottle"
}

func (wb WineBottle) IsGoodGift(h human.Human) bool {
	if h.Age() < wb.minAge {
		return false
	}

	return h.Hobby() == human.HobbyWine
}
