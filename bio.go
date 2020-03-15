package biosphere

// Bio 生命
type Bio interface {
	// Init 初始化
	Init() []int
	// 评分
	Score(dna []int) int
	// 繁殖
	Breed(parent ...[]int) []int
	// 变异
	Variation(dna []int) []int
}
