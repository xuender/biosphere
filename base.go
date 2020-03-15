package biosphere

// base 基础
type base struct {
	init      func() []int
	score     func(dna []int) int
	breed     func(parent ...[]int) []int
	variation func(dna []int) []int
}

// Init 初始化
func (b *base) Init() []int {
	return b.init()
}

// Score 评分
func (b *base) Score(dna []int) int {
	return b.score(dna)
}

// Breed 繁殖
func (b *base) Breed(parent ...[]int) []int {
	return b.breed(parent...)
}

// Variation 变异
func (b *base) Variation(dna []int) []int {
	return b.variation(dna)
}
