package instruments

import "fmt"

type Saxophone struct {}

func (s *Saxophone) PlayMelody() string {
        return fmt.Sprintf("saxophone.PlayMelody()")
}

