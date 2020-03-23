package roby

import "github.com/xuender/oil/random"

// Stage 舞台
type Stage struct {
	data [][]int
	size int
}

// State 状态
func (s Stage) State(x, y int) (ret [5]int) {
	// 上
	if y == 0 {
		ret[0] = 2
	} else {
		ret[0] = s.data[y-1][x]
	}
	// 下
	if y == s.size-1 {
		ret[1] = 2
	} else {
		ret[1] = s.data[y+1][x]
	}
	// 左
	if x == 0 {
		ret[2] = 2
	} else {
		ret[2] = s.data[y][x-1]
	}
	// 右
	if x == s.size-1 {
		ret[3] = 2
	} else {
		ret[3] = s.data[y][x+1]
	}
	// 中
	ret[4] = s.data[y][x]
	return
}

// NewStage 新建舞台
func NewStage(size int) Stage {
	stage := Stage{
		size: size,
		data: make([][]int, size),
	}
	for i := range stage.data {
		stage.data[i] = make([]int, size)
	}
	for r, row := range stage.data {
		for c := range row {
			stage.data[r][c] = 0
		}
	}
	for _, n := range random.NewQueue(size*size, size*size/2) {
		stage.data[n/size][n%size] = 1
	}
	return stage
}
