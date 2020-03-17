package main

import (
	"fmt"
	"math/rand"

	"github.com/xuender/biosphere"
	"github.com/xuender/biosphere/roby"
)

func main() {
	b := biosphere.NewBiosphere(&roby.Roby{})
	b.Run()
	fmt.Println("test")
}

func init() {
	rand.Seed(0)
}
