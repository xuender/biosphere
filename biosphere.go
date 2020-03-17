package biosphere

import (
	"fmt"
	"sort"

	"github.com/xuender/oil/random"
)

// Biosphere 生物圈
type Biosphere struct {
	bio Bio
	// GroupSize 族群大小
	GroupSize int
	// EvalTimes 进化次数
	EvalTimes int
	// 每代尝试次数
	TryTimes int
	// 变异次数
	VariationTimes int
	// 幸存数量
	Survival int
	group    []*obj // 族群
}

// Run 运行
func (b *Biosphere) Run() {
	// 初始化族群
	b.group = make([]*obj, b.GroupSize)
	for i := 0; i < b.GroupSize; i++ {
		b.group[i] = &obj{dna: b.bio.Init()}
	}
	// data := "455351051456452243351633534451154250254154221451354510453256220034336004402541356453453250053056215343005531151152203453243006450145534054355644033052004366341442126140213356430424341446226153426123453230616665611301614145544550552231432632141"
	// for i, d := range data {
	// 	b.group[0].dna[i] = int(d - 48)
	// }
	// 迭代
	for e := 0; e < b.EvalTimes; e++ {
		// 遍历族群
		for _, g := range b.group {
			// 个体尝试
			for t := 0; t < b.TryTimes; t++ {
				g.add(b.bio.Score(g.dna))
			}
		}
		sort.Slice(b.group, func(i, j int) bool {
			return b.group[i].Score() > b.group[j].Score()
		})
		fmt.Printf("迭代: %04d", e+1)
		b.best()
		// 繁殖
		s := make([]random.Scorer, len(b.group))
		for i, o := range b.group {
			s[i] = o
		}
		r := random.NewRoulette(s)
		for i := b.Survival; i < b.GroupSize; i++ {
			b.group[i].dna = b.bio.Breed(r.Take().(*obj).dna, r.Take().(*obj).dna)
			b.group[i].scores = []int{}
			b.group[i].score = 0
			// 变异
			for f := 0; f < b.VariationTimes; f++ {
				b.bio.Variation(b.group[i].dna)
			}
		}
	}
}

// best 显示最佳
func (b *Biosphere) best() {
	fmt.Printf(" 最高分: %03d, AGE: %02d ID: %s\n", b.group[0].Score(), len(b.group[0].scores)/b.TryTimes, b.group[0])
}

// NewBiosphere 新建生物圈
func NewBiosphere(bio Bio) *Biosphere {
	return &Biosphere{
		bio:            bio,
		GroupSize:      200,
		EvalTimes:      3000,
		TryTimes:       20,
		Survival:       20,
		VariationTimes: 1, // 变异不能太严重
	}
}
