package road_trip

import (
	passenger2 "course/hw6/road_trip/passenger"
	"course/hw6/road_trip/vehicle"
)

func StartTrip() {
	passenger := passenger2.Passenger{
		Id:        "123",
		FirstName: "John",
		LastName:  "Doe",
	}

	tesla := vehicle.NewCar("Tesla", 5, 200, 5)
	boeing := vehicle.NewPlane("Boeing", 200, 1000, 100)
	uz := vehicle.NewTrain("UZ", 1000, 100, 10)

	route := NewRoute("Kyiv", "Lviv", 1000)
	route.AddVehicle(VehicleDuration{300, &tesla})
	route.AddVehicle(VehicleDuration{500, &boeing})
	route.AddVehicle(VehicleDuration{200, &uz})

	route.ShowAllVehicles()
	route.StartTrip(passenger)
}
