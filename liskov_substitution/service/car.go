package service

import "fmt"

type Car struct {
        Name string
        Speed float64
}

var CarImpl VehicleInterface = &Car{}

func (c *Car) Ride() string{
	return fmt.Sprintf("Ride in %s.", c.Name)
}

func (c *Car) Transport(dist string) string {
	return fmt.Sprintf("Go to %s at %.2f[km/h]", dist, c.Speed)
}
