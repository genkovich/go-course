package vehicle

import (
	"course/hw6/road_trip/passenger"
	"fmt"
)

type Train struct {
	title      string
	capacity   int
	passengers map[string]passenger.Passenger
	speed      int
	maxSpeed   int
	minSpeed   int
}

func (t *Train) Move() {
	t.speed = int(float64(t.maxSpeed) * 0.8)
}

func (t *Train) Stop() {
	t.speed = 0
}

func (t *Train) ChangeSpeed(speed int) {
	if speed > t.maxSpeed {
		fmt.Printf("The train can't move faster then %d\n", t.maxSpeed)
		return
	}

	if speed < t.minSpeed {
		fmt.Printf("The train can't move slowly then %d\n", t.minSpeed)
		return
	}

	t.speed = speed
}

func (t *Train) PickUp(passenger passenger.Passenger) {
	if _, ok := t.passengers[passenger.Id]; ok {
		fmt.Printf("%s is already on the train\n", passenger.LastName)
		return
	}

	if len(t.passengers) >= t.capacity {
		fmt.Printf("The train is full, %s can't board\n", passenger.LastName)
		return
	}

	t.passengers[passenger.Id] = passenger
	fmt.Printf("%s boarded the train\n", passenger.LastName)
}

func (t *Train) DropOff(passenger passenger.Passenger) {
	if _, ok := t.passengers[passenger.Id]; !ok {
		fmt.Printf("%s is not on the train\n", passenger.LastName)
		return
	}

	delete(t.passengers, passenger.Id)
	fmt.Printf("%s left the train\n", passenger.LastName)
}

func (t *Train) Title() string {
	return t.title
}

func NewTrain(title string, capacity int, maxSpeed int, minSpeed int) Train {
	return Train{
		title:      title,
		capacity:   capacity,
		speed:      0,
		maxSpeed:   maxSpeed,
		minSpeed:   minSpeed,
		passengers: make(map[string]passenger.Passenger),
	}
}
