package instruments

import "fmt"

type Piano struct {}

func (p *Piano) PlayChords() string {
        return fmt.Sprintf("piano.PlayChords()")
}

