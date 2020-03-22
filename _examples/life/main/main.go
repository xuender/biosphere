package main

import (
	"github.com/xuender/biosphere"
	"github.com/xuender/biosphere/_examples/life"
)

func main() {
	b := biosphere.NewBiosphere(&life.Life{})
	b.EvalTimes = 5000
	b.TryTimes = 66
	b.VariationTimes = 5
	b.Run()
	b.Best(3)
	b.Chart("chart")
}
