package vehicle

import (
	"course/hw6/road_trip/passenger"
	"fmt"
)

type Car struct {
	title      string
	capacity   int
	passengers map[string]passenger.Passenger
	speed      int
	maxSpeed   int
	minSpeed   int
}

func (c *Car) Move() {
	c.speed = int(float64(c.maxSpeed) * 0.7)
}

func (c *Car) Stop() {
	c.speed = 0
}

func (c *Car) ChangeSpeed(speed int) {
	if speed > c.maxSpeed {
		fmt.Printf("The car can't move faster then %d\n", c.maxSpeed)
		return
	}

	if speed < c.minSpeed {
		fmt.Printf("The car can't move slowly then %d\n", c.minSpeed)
		return
	}

	c.speed = speed
}

func (c *Car) PickUp(passenger passenger.Passenger) {
	if _, ok := c.passengers[passenger.Id]; ok {
		fmt.Printf("%s is already in the car\n", passenger.LastName)
		return
	}

	if len(c.passengers) >= c.capacity {
		fmt.Printf("The car is full, %s can't board\n", passenger.LastName)
		return
	}

	c.passengers[passenger.Id] = passenger
	fmt.Printf("%s boarded the car\n", passenger.LastName)
}

func (c *Car) DropOff(passenger passenger.Passenger) {
	if _, ok := c.passengers[passenger.Id]; !ok {
		fmt.Printf("%s is not in the car\n", passenger.LastName)
		return
	}

	delete(c.passengers, passenger.Id)
	fmt.Printf("%s left the car\n", passenger.LastName)
}

func (c *Car) Title() string {
	return c.title
}

func NewCar(title string, capacity int, maxSpeed int, minSpeed int) Car {
	return Car{
		title:      title,
		capacity:   capacity,
		maxSpeed:   maxSpeed,
		minSpeed:   minSpeed,
		passengers: make(map[string]passenger.Passenger),
		speed:      0,
	}
}
