package life

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/biosphere"
)

func ExampleLife_move() {
	r := Life{}
	dna := make([]int, DNASize)
	for i := range dna {
		dna[i] = i
	}
	m := map[int]bool{}
	for a := 0; a < 4; a++ {
		for b := 0; b < 4; b++ {
			for c := 0; c < 4; c++ {
				for d := 0; d < 4; d++ {
					for e := 0; e < 4; e++ {
						m[r.movement(dna, [5]int{a, b, c, d, e})] = true
						m[r.movement(dna, [5]int{a, b, c, d, e})] = true
						// fmt.Println(r.movement(dna, [5]int{3, 3, 3, e, 3}, integer.MinInt))
					}
				}
			}
		}
	}
	fmt.Println("len", len(m))

	// Output:
	// len 1024
}
func TestLife_movement(t *testing.T) {
	r := Life{}
	dna := make([]int, DNASize)
	for i := range dna {
		dna[i] = i
	}
	m := map[int]bool{}
	for a := 0; a < 4; a++ {
		for b := 0; b < 4; b++ {
			for c := 0; c < 4; c++ {
				for d := 0; d < 4; d++ {
					for e := 0; e < 4; e++ {
						m[r.movement(dna, [5]int{a, b, c, d, e})] = true
						m[r.movement(dna, [5]int{a, b, c, d, e})] = true
					}
				}
			}
		}
	}
	assert.Equal(t, DNASize, len(m), "DNA链")
}

// func TestRoby_Work(t *testing.T) {
// 	r := NewRoby()
// 	r.Work()
// 	assert.Equal(t, 1, len(r.scores), "积分")
// 	r.Work()
// 	assert.Equal(t, 2, len(r.scores), "积分")
// }

func ExampleLife() {
	r := &Life{}
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
	r := Life{}
	dna := make([]int, DNASize)
	data := "451350252252250144351153022155351255154155150151331330054352253055054141053133055052054202053432044052531024051253153055410034113314251240440200233542312353433215154150312155350300531532015314135253514123215114121555320215321250410322201502130"
	for i, d := range data {
		dna[i] = int(d - 48)
	}
	assert.Greater(t, r.Score(dna), 100, "大于100分")
}
