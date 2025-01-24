package main

import (
	"fmt"
	"sort"
)

type at struct {
	tname string
	score int
}

func p3() {

	// 部分集合を列挙する例
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)

	elements := []at{
		{"A", a}, {"B", b}, {"C", c}, {"D", d}, {"E", e}} // 対象の要素
	N := len(elements) // 要素の数

	// 全ての部分集合を列挙
	var ans = make([]at, 32)

	for bit := 0; bit < (1 << N); bit++ {
		var subset string
		var tscore int
		for i := 0; i < N; i++ {
			// i番目の要素が部分集合に含まれるかを判定
			if bit&(1<<i) != 0 {
				subset = subset + elements[i].tname
				tscore += elements[i].score
				ans[bit] = at{subset, tscore}
			}
		}

	}

	sort.SliceStable(ans, func(i, j int) bool {

		// スコアで降順
		if ans[i].score != ans[j].score {
			return ans[i].score > ans[j].score
		}
		// スコアが同じなら名前で辞書順
		return ans[i].tname < ans[j].tname

	})

	for i := 0; i < 31; i++ {
		fmt.Println(ans[i].tname)
	}

}
