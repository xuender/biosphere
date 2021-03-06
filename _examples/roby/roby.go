package roby

import (
	"math/rand"

	"github.com/xuender/oil/integer"
)

// NUM 平均验证次数
const NUM = 10

// DNASize DNA 尺寸
const DNASize = 243

// MOVEMENT 6种动作
const MOVEMENT = 6

// Roby 机器人
type Roby struct{}

// Init 初始化
func (r *Roby) Init() []int {
	dna := make([]int, DNASize)
	for i := 0; i < DNASize; i++ {
		// DNA 全随机
		dna[i] = rand.Intn(MOVEMENT)
	}
	return dna
}

// Score 积分
func (r *Roby) Score(dna []int, times int) int {
	// 新建舞台
	// 舞台尺寸可以不断变化，从小到大，学习小舞台，到大舞台
	size := 4
	if times > 500 {
		size = 6
	}
	if times > 1000 {
		size = 8
	}
	if times > 1500 {
		size = 10
	}
	stage := NewStage(size)
	// 初始位置
	x := rand.Intn(size)
	y := rand.Intn(size)
	score := 0

	// 200个动作
	for i := 0; i < 200; i++ {
		state := stage.State(x, y)
		// 获取行动
		m := r.movement(dna, state)
		// 行动
		switch m {
		case 0, 1, 2, 3: // 上下左右
			if ok, nx, ny := r.move(m, state, x, y); ok {
				x = nx
				y = ny
			} else {
				score -= 5 // 撞墙扣5分
			}
		case 4: // 随机移动
			if ok, nx, ny := r.move(rand.Intn(4), state, x, y); ok {
				x = nx
				y = ny
			} else {
				score -= 5 // 撞墙扣5分
			}
		case 5: // 捡罐头
			if stage.data[y][x] == 1 {
				stage.data[y][x] = 0
				score += 10 // 捡到罐子10分
			} else {
				score-- // 没有罐子扣1分
			}
		}
	}
	return score
}

// Breed 繁殖
func (r *Roby) Breed(parent ...[]int) []int {
	dna := make([]int, DNASize)
	// 前后段落拼接
	// h := DNASize / len(parent)
	// for f, p := range parent {
	// 	for i := 0; i < h; i++ {
	// 		l := f*h + i
	// 		if l >= DNASize {
	// 			break
	// 		}
	// 		dna[l] = p[l]
	// 	}
	// }
	for i := range dna {
		dna[i] = parent[i%len(parent)][i]
	}
	return dna
}

// Variation 变异
func (r *Roby) Variation(dna []int) []int {
	dna[rand.Intn(DNASize)] = rand.Intn(MOVEMENT)
	return dna
}

// movement 动作
func (r *Roby) movement(dna []int, around [5]int) int {
	id := 0
	for i, a := range around {
		if a > 0 {
			id += integer.Exp(3, 4-i) * a
		}
	}
	return dna[id]
}

// move 移动检查
func (r *Roby) move(movement int, state [5]int, x, y int) (bool, int, int) {
	switch movement {
	case 0:
		// 上
		if state[0] == 2 {
			// 碰壁
			return false, x, y
		}
		return true, x, y - 1
	case 1:
		// 下
		if state[1] == 2 {
			// 碰壁
			return false, x, y
		}
		return true, x, y + 1
	case 2:
		// 左
		if state[2] == 2 {
			// 碰壁
			return false, x, y
		}
		return true, x - 1, y
	case 3:
		// 右
		if state[3] == 2 {
			// 碰壁
			return false, x, y
		}
		return true, x + 1, y
	default:
		return true, x, y
	}
}
