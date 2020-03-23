package roby

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStage(t *testing.T) {
	s := NewStage(10)
	assert.Equal(t, 10, len(s.data), "行列")
	count := 0
	for _, row := range s.data {
		for _, col := range row {
			count += col
		}
	}
	assert.Equal(t, 50, count, "罐子")
}

func ExampleNewStage() {
	s := NewStage(10)
	fmt.Println(len(s.data))

	// Output:
	// 10
}
