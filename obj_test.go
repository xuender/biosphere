package biosphere

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_obj_add(t *testing.T) {
	o := obj{}
	o.add(3)
	assert.Equal(t, 1, len(o.scores), "积分")
	assert.Equal(t, 3, o.Score(), "积分")
}

func Test_obj_Score(t *testing.T) {
	o := obj{}
	o.scores = append(o.scores, 3, 3, 2)
	assert.Equal(t, 3, o.Score(), "平均分")
}
