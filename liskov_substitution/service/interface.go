package service

type VehicleInterface interface {
        Ride() string
        Transport(string) string
}

var CarImpl VehicleInterface = &Car{}
var BicycleImpl VehicleInterface = &Bicycle{}
var SpaceXImpl VehicleInterface = &SpaceX{}
