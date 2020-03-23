package roby

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/biosphere"
)

func TestRoby_movement(t *testing.T) {
	r := Roby{}
	dna := make([]int, 243)
	for i := range dna {
		dna[i] = i
	}
	m := map[int]bool{}
	for a := 0; a < 3; a++ {
		for b := 0; b < 3; b++ {
			for c := 0; c < 3; c++ {
				for d := 0; d < 3; d++ {
					for e := 0; e < 3; e++ {
						m[r.movement(dna, [5]int{a, b, c, d, e})] = true
					}
				}
			}
		}
	}
	assert.Equal(t, 243, len(m), "DNA链")
}

// func TestRoby_Work(t *testing.T) {
// 	r := NewRoby()
// 	r.Work()
// 	assert.Equal(t, 1, len(r.scores), "积分")
// 	r.Work()
// 	assert.Equal(t, 2, len(r.scores), "积分")
// }

func ExampleRoby() {
	r := &Roby{}
	b := biosphere.NewBiosphere(r)
	b.GroupSize = 10
	b.TryTimes = 3
	b.EvalTimes = 5
	b.Survival = 3
	b.VariationTimes = 20
	b.Run()

	fmt.Println("test")
	// Output:
	// 350
}

func TestRoby_Score(t *testing.T) {
	r := Roby{}
	dna := make([]int, DNASize)
	data := "451350252252250144351153022155351255154155150151331330054352253055054141053133055052054202053432044052531024051253153055410034113314251240440200233542312353433215154150312155350300531532015314135253514123215114121555320215321250410322201502130"
	for i, d := range data {
		dna[i] = int(d - 48)
	}
	assert.Greater(t, r.Score(dna, 100), 100, "大于100分")
}
