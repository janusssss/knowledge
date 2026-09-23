package main

import (
	"go-career/algorithm/leetcode"
	"sort"
)

type exampleMerge struct {
	intervals [][]int
	answer    [][]int
}

func main() {
	s := []exampleMerge{
		{leetcode.To2intSlice("[[1,4],[1,4]]"), leetcode.To2intSlice("[[1,4]]")},
		{leetcode.To2intSlice("[[1,3],[2,6],[8,10],[15,18]]"), leetcode.To2intSlice("[[1,6],[8,10],[15,18]]")},
		{leetcode.To2intSlice("[[1,4],[5,6]]"), leetcode.To2intSlice("[[1,4],[5,6]]")},
	}
	for _, v := range s {
		leetcode.Printf(merge(v.intervals), v.answer)
	}
}

func merge(intervals [][]int) [][]int {
	// 按每个区间的起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var merged [][]int
	for _, interval := range intervals {
		// 如果 merged 切片为空，或者当前区间的起始位置大于上一区间的结束位置，直接添加
		if len(merged) == 0 || merged[len(merged)-1][1] < interval[0] {
			merged = append(merged, interval)
		} else {
			// 否则的话，我们就可以与上一区间进行合并
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], interval[1])
		}
	}
	return merged
}
