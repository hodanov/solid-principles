package service

import "fmt"

type Vehicle struct {
	Name  string
	Speed int
}

func (v *Vehicle) Ride() string {
	return fmt.Sprintf("%sに乗って", v.Name)
}

func (v *Vehicle) Transport(dist string) string {
	return fmt.Sprintf("時速%d[km/s]で%sへ移動する", v.Speed, dist)
}
