package main

import (
	"fmt"

	"main.go/service"
)

var (
	car     = service.CarImpl
	bicycle = service.BicycleImpl
	rocket  = service.RocketImpl
)

func main() {
	car = &service.Car{
		Name:  "TOYOTA Land Cruiser",
		Speed: 60,
	}
	bicycle = &service.Bicycle{
		Name:  "Jamis Ventura Comp",
		Speed: 10,
	}
	rocket = &service.Rocket{
		Name:  "SpaceX Falcon 9",
		Speed: 7.9,
	}

	fmt.Println(car.Ride() + " " + car.Transport("Tokyo"))
	fmt.Println(bicycle.Ride() + " " + bicycle.Transport("nearby park"))
	fmt.Println(rocket.Ride() + " " + rocket.Transport("space"))
}
