package main

import (
	"github.com/xuender/biosphere"
	"github.com/xuender/biosphere/_examples/roby"
)

func main() {
	b := biosphere.NewBiosphere(&roby.Roby{})
	b.Run()
	b.Best(3)
}
