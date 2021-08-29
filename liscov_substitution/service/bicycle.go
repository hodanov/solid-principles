package service

import "fmt"

type Bicycle struct {
        Name string
        Speed float64
        //Vehicle
}

func (b *Bicycle) Ride() string{
	return fmt.Sprintf("%sに乗って", b.Name)
}

func (b *Bicycle) Transport(dist string) string {
	return fmt.Sprintf("時速%.2f[km/h]で%sへ移動する", b.Speed, dist)
}
