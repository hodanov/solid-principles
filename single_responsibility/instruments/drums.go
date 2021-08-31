package instruments

import "fmt"

type Drums struct {}

func (d *Drums) PlayRhythm() string {
        return fmt.Sprintf("drums.PlayRhythm()")
}

