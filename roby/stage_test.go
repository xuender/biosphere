package roby

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
			count += col
		}
	}
	assert.Equal(t, 50, count, "罐子")
}

func ExampleNewStage() {
	s := NewStage()
	fmt.Println(len(s))

	// Output:
	// 10
}
