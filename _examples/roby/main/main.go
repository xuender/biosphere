package main

import (
	"io/ioutil"

	"github.com/xuender/biosphere"
	"github.com/xuender/biosphere/_examples/roby"
)

func main() {
	b := biosphere.NewBiosphere(&roby.Roby{})
	ioutil.WriteFile("chart.svg", []byte(chart.Polyline(b.Run())), 0644)
	b.Best(3)
}
