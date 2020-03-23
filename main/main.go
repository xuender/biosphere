package main

import (
	"github.com/xuender/biosphere"
	"github.com/xuender/biosphere/_examples/roby"
)

func main() {
	b := biosphere.NewBiosphere(&roby.Roby{})
	b.EvalTimes = 3000
	b.TryTimes = 40
	b.PrintTimes(500, 1000, 1500)
	b.Run()
	b.Best(1)
	b.Chart("/tmp/chart")
}
