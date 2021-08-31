package main

import (
	"fmt"

	"main.go/service"
)

func main() {
	var car service.VehicleInterface = &service.Car{
		Name:  "TOYOTA Land Cruiser",
		Speed: 60,
	}
	var bicycle service.VehicleInterface = &service.Bicycle{
		Name:  "Jamis Ventura Comp",
		Speed: 10,
	}
	var spaceX service.VehicleInterface = &service.SpaceX{
		Name:  "Falcon 9",
		Speed: 7.9,
	}

	fmt.Println(car.Ride() + " " + car.Transport("Tokyo"))
	fmt.Println(bicycle.Ride() + " " + bicycle.Transport("nearby park"))
	fmt.Println(spaceX.Ride() + " " + spaceX.Transport("space"))
}
