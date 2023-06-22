package vehicle

import (
	"course/hw6/road_trip/passenger"
	"fmt"
)

type Plane struct {
	title          string
	capacity       int
	passengers     map[string]passenger.Passenger
	speed          int
	maxSpeed       int
	minSpeed       int
	bookingList    map[string]string
	parachuteCount int
}

func (p *Plane) Move() {
	p.speed = int(float64(p.maxSpeed) * 0.6)
}

func (p *Plane) Stop() {
	if p.speed > p.minSpeed {
		fmt.Printf("You can`t stop the plane, please reduce the speed\n")
		return
	}
	p.speed = 0
}

func (p *Plane) ChangeSpeed(speed int) {
	if speed > p.maxSpeed {
		fmt.Printf("The plane can't move faster then %d\n", p.maxSpeed)
		return
	}

	if speed < p.minSpeed {
		fmt.Printf("The plane can't move slowly then %d\n", p.minSpeed)
		return
	}

	p.speed = speed
}

func (p *Plane) PickUp(passenger passenger.Passenger) {
	if _, ok := p.bookingList[passenger.Id]; !ok {
		fmt.Printf("%s is not booked on the plane\n", passenger.LastName)
		return
	}

	if _, ok := p.passengers[passenger.Id]; ok {
		fmt.Printf("%s is already on the plane\n", passenger.LastName)
		return
	}

	if len(p.passengers) >= p.capacity {
		fmt.Printf("The plane is full, %s can't board\n", passenger.LastName)
		return
	}

	p.passengers[passenger.Id] = passenger
	fmt.Printf("%s boarded the plane\n", passenger.LastName)

}

func (p *Plane) DropOff(passenger passenger.Passenger) {
	if p.speed > 0 && p.parachuteCount <= 0 {
		fmt.Printf("You can`t drop off the plane without parachute, please stop the plane\n")
		return
	}

	if _, ok := p.passengers[passenger.Id]; !ok {
		fmt.Printf("%s is not on the plane\n", passenger.LastName)
		return
	}

	delete(p.passengers, passenger.Id)

	if p.speed > 0 {
		p.parachuteCount--
		fmt.Printf("%s left the plane with parachute\n", passenger.LastName)
		return
	} else {
		fmt.Printf("%s left the plane\n", passenger.LastName)

	}
}

func (p *Plane) Title() string {
	return p.title
}

func NewPlane(title string, capacity int, maxSpeed int, minSpeed int) Plane {
	return Plane{
		title:      title,
		capacity:   capacity,
		passengers: make(map[string]passenger.Passenger),
		speed:      0,
		maxSpeed:   maxSpeed,
		minSpeed:   minSpeed,
	}
}
