package life

import (
	"math/rand"

	"github.com/xuender/oil/integer"
)

// NUM 平均验证次数
const NUM = 10

// DNASize DNA 尺寸
const DNASize = 1024

// MOVEMENT 6种动作
const MOVEMENT = 6

// VITALITY 初始活力
const VITALITY = 200

// Life 生命
type Life struct{}

// Init 初始化
func (r *Life) Init() []int {
	dna := make([]int, DNASize)
	for i := 0; i < DNASize; i++ {
		// DNA 全随机
		dna[i] = rand.Intn(MOVEMENT)
	}
	return dna
}

// Score 积分
func (r *Life) Score(dna []int, times int) int {
	// 新建舞台
	stage := NewStage()
	// 初始位置
	x := rand.Intn(10)
	y := rand.Intn(10)
	score := 0

	// 200个动作
	for i := 0; i < VITALITY; i++ {
		state := stage.State(x, y, VITALITY-i)
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
		case 5: // 亲近
			if stage[y][x] > 0 {
				stage[y][x] = 0 // 训练过程除去对象
				if stage[y][x] > VITALITY-i {
					score += 15 // 有对象+10
					// 	i -= 10
					// } else {
					// 	i -= 5
				} else {
					score += 10 // 有对象+10

				}
			} else {
				score-- // 没有对象扣1分
			}
		}
	}
	return score
}

// Breed 繁殖
func (r *Life) Breed(parent ...[]int) []int {
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
func (r *Life) Variation(dna []int) []int {
	dna[rand.Intn(DNASize)] = rand.Intn(MOVEMENT)
	return dna
}

// movement 动作
func (r *Life) movement(dna []int, around [5]int) int {
	id := 0
	for i, a := range around {
		// if a > 0 {
		switch a {
		case 0:
			// 空 id += integer.Exp(4, i) * 0
		case 1:
			// 墙壁
			id += integer.Exp(4, i) * 1
		case 2:
			// 高活力
			id += integer.Exp(4, i) * 2
		case 3:
			// 低活力
			id += integer.Exp(4, i) * 3
		}
	}
	return dna[id]
}

// move 移动检查
func (r *Life) move(movement int, state [5]int, x, y int) (bool, int, int) {
	switch movement {
	case 0:
		// 上
		if state[0] == 1 {
			// 碰壁
			return false, x, y
		}
		return true, x, y - 1
	case 1:
		// 下
		if state[1] == 1 {
			// 碰壁
			return false, x, y
		}
		return true, x, y + 1
	case 2:
		// 左
		if state[2] == 1 {
			// 碰壁
			return false, x, y
		}
		return true, x - 1, y
	case 3:
		// 右
		if state[3] == 1 {
			// 碰壁
			return false, x, y
		}
		return true, x + 1, y
	default:
		return true, x, y
	}
}
