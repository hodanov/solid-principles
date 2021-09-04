package service

import "fmt"

type SpaceX struct {
        Name string
        Speed float64
}

func (s *SpaceX) Ride() string{
	return fmt.Sprintf("Ride in %s.", s.Name)
}

func (s *SpaceX) Transport(dist string) string {
        var text string
        if dist == "space" {
                text = fmt.Sprintf("Go into the space at %.2f[km/s].", s.Speed)
        } else {
                text = "This rocket can only go into the space."
        }
	return text
}
