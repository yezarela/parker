package main

import (
	"log"

	"github.com/yezarela/parker/park"
)

func main() {

	p := park.NewPark()

	p.InitSlots(40)
	a, err := p.In(park.Car{"A1", "Black"})
	if err != nil {
		log.Println(err)
	}

	log.Println(*a)
	out, err := p.Out(*a)
	if err != nil {
		log.Println(err)
	}

	log.Println(*out)
}
