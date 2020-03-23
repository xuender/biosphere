package biosphere

import (
	"fmt"
	"strings"

	"github.com/xuender/oil/str"

	"github.com/xuender/oil/integer"
)

type obj struct {
	dna    []int
	scores []int
	score  int
}

func (o *obj) String() string {
	return fmt.Sprintf(
		"积分: %d, 年龄:%d\n%s\n",
		o.Score(),
		len(o.scores),
		strings.Join(str.Chunk(integer.String(o.dna...), 50), "\n"),
	)
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
