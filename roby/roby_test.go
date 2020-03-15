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
	data := "455351051456452243351633534451154250254154221451354510453256220034336004402541356453453250053056215343005531151152203453243006450145534054355644033052004366341442126140213356430424341446226153426123453230616665611301614145544550552231432632141"
	for i, d := range data {
		dna[i] = int(d - 48)
	}
	assert.Greater(t, r.Score(dna), 100, "大于100分")
}
