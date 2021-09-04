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

func transport(v service.VehicleInterface, d string) string {
	return fmt.Sprintf("%s %s", v.Ride(), v.Transport(d))
}

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

	fmt.Println(transport(car, "Tokyo"))
	fmt.Println(transport(bicycle, "nearby park"))
	fmt.Println(transport(rocket, "space"))
}
