package main

import (
	"fmt"

	"main.go/service"
)

var (
	car     = service.CarImpl
	bicycle = service.BicycleImpl
	spaceX  = service.SpaceXImpl
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
	spaceX = &service.SpaceX{
		Name:  "Falcon 9",
		Speed: 7.9,
	}

	fmt.Println(car.Ride() + " " + car.Transport("Tokyo"))
	fmt.Println(bicycle.Ride() + " " + bicycle.Transport("nearby park"))
	fmt.Println(spaceX.Ride() + " " + spaceX.Transport("space"))
}
