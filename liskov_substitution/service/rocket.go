package service

import "fmt"

type Rocket struct {
        Name string
        Speed float64
}

var RocketImpl VehicleInterface = &Rocket{}

func (s *Rocket) Ride() string{
	return fmt.Sprintf("Ride in %s.", s.Name)
}

func (s *Rocket) Transport(dist string) string {
        var text string
        if dist == "space" {
                text = fmt.Sprintf("Go into the space at %.2f[km/s].", s.Speed)
        } else {
                text = "This rocket can only go into the space."
        }
	return text
}
