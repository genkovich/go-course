package vehicle

import (
	"course/hw6/road_trip/passenger"
)

type vehicleDriving interface {
	Move()
	Stop()
	ChangeSpeed(speed int)
}

type vehicleBoarding interface {
	PickUp(passenger passenger.Passenger)
	DropOff(passenger passenger.Passenger)
}

type Vehicle interface {
	Title() string
	vehicleDriving
	vehicleBoarding
}
