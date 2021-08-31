package main

import "main.go/facade"

func main() {
	instrument := facade.NewInstrumentFacade()
	instrument.PlayMusic()
}
