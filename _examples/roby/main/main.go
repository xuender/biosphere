package main

import (
	"github.com/xuender/biosphere"
	"github.com/xuender/biosphere/_examples/roby"
)

func main() {
	b := biosphere.NewBiosphere(&roby.Roby{})
	b.EvalTimes = 3000
	b.TryTimes = 3
	b.Run()
	b.Best(3)
	b.Chart("chart")
}
