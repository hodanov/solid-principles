package service

import "fmt"

type Car struct {
        Name string
        Speed float64
//        Vehicle
}

func (c *Car) Ride() string{
	return fmt.Sprintf("%sに乗って", c.Name)
}

func (c *Car) Transport(dist string) string {
	return fmt.Sprintf("時速%.2f[km/h]で%sへ移動する", c.Speed, dist)
}
