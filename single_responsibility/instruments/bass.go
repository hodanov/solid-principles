package instruments

import "fmt"

type Bass struct{}

func (b *Bass) PlayBass() string {
	return fmt.Sprintf("bass.PlayBass()")
}
