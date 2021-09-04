package service

import "fmt"

type Vehicle struct {
	Name  string
	Speed int
}

func (v *Vehicle) Ride() string {
	return fmt.Sprintf("Ride in %s", v.Name)
}

func (v *Vehicle) Transport(dist string) string {
	return fmt.Sprintf("move to %s at %d[km/s]", dist, v.Speed)
}
