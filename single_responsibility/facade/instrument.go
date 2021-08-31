package facade

import (
	"fmt"

	"main.go/instruments"
)

type InstrumentFacade struct {
        Bass *instruments.Bass
        Drums *instruments.Drums
        Piano *instruments.Piano
        Saxophone *instruments.Saxophone
}

func NewInstrumentFacade() *InstrumentFacade{
        return &InstrumentFacade{
                &instruments.Bass{},
                &instruments.Drums{},
                &instruments.Piano{},
                &instruments.Saxophone{},
        }
}

func (i *InstrumentFacade) PlayMusic() {
        fmt.Println(i.Bass.PlayBass())
        fmt.Println(i.Drums.PlayRhythm())
        fmt.Println(i.Piano.PlayChords())
        fmt.Println(i.Saxophone.PlayMelody())
}
