package service

import "fmt"

type Bicycle struct {
        Name string
        Speed float64
}

func (b *Bicycle) Ride() string{
	return fmt.Sprintf("Ride in %s.", b.Name)
}

func (b *Bicycle) Transport(dist string) string {
	return fmt.Sprintf("Go to %s at %.2f[km/h]", dist, b.Speed)
}
