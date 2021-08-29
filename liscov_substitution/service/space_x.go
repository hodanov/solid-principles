package service

import "fmt"

type SpaceX struct {
        Name string
        Speed float64
}

func (b *SpaceX) Ride() string{
	return fmt.Sprintf("%sに乗って", b.Name)
}

func (s *SpaceX) Transport(dist string) string {
	return fmt.Sprintf("秒速%.2f[km/s]で%sへ移動する", s.Speed, dist)
}
