package biosphere

import (
	"fmt"

	"github.com/xuender/oil/integer"
)

type obj struct {
	dna    []int
	scores []int
	score  int
}

func (o *obj) String() string {
	return fmt.Sprintf(" %d", o.dna[:24])
}

func (o *obj) add(score int) {
	if o.scores == nil {
		o.scores = []int{}
	}
	o.scores = append(o.scores, score)
	o.score = 0
}

func (o *obj) Score() int {
	if o.score > 0 {
		return o.score
	}
	sum := 0
	for _, s := range o.scores {
		sum += s
	}
	o.score = integer.Div(sum, len(o.scores))
	return o.score
}
