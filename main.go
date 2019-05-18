package main

import (
	"github.com/yezarela/parker/cmd"

	"github.com/yezarela/parker/park"
)

func main() {

	p := park.NewPark()
	c := cmd.NewCommander(p)

	c.ReadFromStdIn()

}
