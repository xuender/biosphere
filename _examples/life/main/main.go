package main

import (
	"github.com/xuender/biosphere"
	"github.com/xuender/biosphere/_examples/left"
)

func main() {
	b := biosphere.NewBiosphere(&left.Left{})
	b.EvalTimes = 1000
	b.TryTimes = 3
	b.Run()
	b.Best(3)
	b.Chart("chart")
}
