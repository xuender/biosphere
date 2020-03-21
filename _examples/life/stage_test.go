package life

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStage(t *testing.T) {
	s := NewStage()
	assert.Equal(t, 10, len(s), "行列")
	count := 0
	for _, row := range s {
		for _, col := range row {
			if col > 0 {
				count++
			}
		}
	}
	assert.Equal(t, 50, count, "对象")
}

func ExampleNewStage() {
	s := NewStage()
	fmt.Println(len(s))

	// Output:
	// 10
}
