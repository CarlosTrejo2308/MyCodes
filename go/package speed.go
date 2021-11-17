package speed

import "time"

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{speed: speed, batteryDrain: batteryDrain, battery: 100}
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	return Track{distance}
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	if car.batteryDrain > car.battery {
		return car
	}

	car.distance += car.speed
	car.battery -= car.batteryDrain
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	for car.batteryDrain < car.battery {
		car = Drive(car)
	}

	return car.distance >= track.distance
}

func times() {
	tm, _ := time.Parse(time.RFC3339, "2019-01-01T00:00:00Z")

}