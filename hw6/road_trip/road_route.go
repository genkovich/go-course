package road_trip

import (
	"course/hw6/road_trip/passenger"
	"course/hw6/road_trip/vehicle"
	"fmt"
)

type RoadRoute struct {
	departure   string
	destination string
	duration    int
	vehicles    []VehicleDuration
}

type VehicleDuration struct {
	duration int
	vehicle.Vehicle
}

func (r *RoadRoute) AddVehicle(duration VehicleDuration) {
	r.vehicles = append(r.vehicles, duration)
}

func (r *RoadRoute) ShowAllVehicles() {
	fmt.Println("Vehicles on the route:")
	for _, v := range r.vehicles {
		fmt.Printf("%s\n", v.Title())
	}
}

func (r *RoadRoute) StartTrip(passenger passenger.Passenger) {
	allVehiclesDuration := 0
	for _, v := range r.vehicles {
		allVehiclesDuration += v.duration
	}

	if r.duration > allVehiclesDuration {
		fmt.Printf("Can't start trip, not enought transport\n")
		return
	}

	fmt.Printf("Starting trip from %s\n", r.departure)
	for _, v := range r.vehicles {
		v.PickUp(passenger)
		v.Move()
		fmt.Printf("%s covered %d by %s\n", passenger.LastName, v.duration, v.Title())
		v.Stop()
		v.DropOff(passenger)
	}
	fmt.Printf("The trip is over, welcome to %s\n", r.destination)

}

func NewRoute(departure string, destination string, duration int) RoadRoute {
	return RoadRoute{
		departure:   departure,
		destination: destination,
		duration:    duration,
		vehicles:    []VehicleDuration{},
	}
}
