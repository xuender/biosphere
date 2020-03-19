package main

import (
	"github.com/xuender/biosphere"
	"github.com/xuender/biosphere/_examples/roby"
)

func main() {
	b := biosphere.NewBiosphere(&roby.Roby{})
	b.EvalTimes = 10000
	b.TryTimes = 4
	b.Run()
	b.Best(3)
}
