package biosphere

import (
	"fmt"
	"sort"

	"github.com/vbauerster/mpb/v5"
	"github.com/vbauerster/mpb/v5/decor"
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
	// 进度条
	p := mpb.New(mpb.WithWidth(64))
	bar := p.AddBar(
		// 总数
		int64(b.EvalTimes),
		// 进图条前缀
		mpb.PrependDecorators(
			decor.Name("迭代"),
			// 计数
			decor.CountersNoUnit(": %d / %d", decor.WCSyncWidth),
			decor.Any(func(s *decor.Statistics) string {
				return fmt.Sprintf(" 最高分: %03d", b.group[0].Score())
			}),
		),
		// 进度条后缀
		mpb.AppendDecorators(
			// 百分比
			decor.Percentage(),
			// 剩余时间
			decor.OnComplete(
				decor.AverageETA(decor.ET_STYLE_MMSS, decor.WC{W: 6}), "完毕",
			),
		),
	)

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
		// bar.IncrBy(e)
		bar.Increment()
	}
	p.Wait()
}

// Best 显示最佳
func (b *Biosphere) Best(size int) {
	if size > len(b.group) {
		size = len(b.group)
	}
	for i, o := range b.group[:size] {
		fmt.Printf("%d: %s\n", i, o)
	}
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
