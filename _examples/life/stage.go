package life

import (
	"math/rand"

	"github.com/xuender/oil/random"
)

// Stage 舞台
type Stage [10][10]int

// State 状态
func (s Stage) State(x, y int) (ret [5]int) {
	// 上
	if y == 0 {
		ret[0] = -1
	} else {
		ret[0] = s[y-1][x]
	}
	// 下
	if y == 9 {
		ret[1] = -1
	} else {
		ret[1] = s[y+1][x]
	}
	// 左
	if x == 0 {
		ret[2] = -1
	} else {
		ret[2] = s[y][x-1]
	}
	// 右
	if x == 9 {
		ret[3] = -1
	} else {
		ret[3] = s[y][x+1]
	}
	// 中
	ret[4] = s[y][x]
	return
}

// NewStage 新建舞台
func NewStage() Stage {
	stage := Stage{}
	for r, row := range stage {
		for c := range row {
			stage[r][c] = 0
		}
	}
	for _, n := range random.NewQueue(100, 50) {
		stage[n/10][n%10] = rand.Intn(VITALITY+30) + 1
	}
	return stage
}
