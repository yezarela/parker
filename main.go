package main

import (
	"github.com/yezarela/parker/park"
)

func main() {

	p := park.NewPark()

	p.InitSlots(40)
	p.In(park.Car{"A1", "Black"})
	p.Out(1)
}
