package main

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	right := points[0][1]
	ans := 1
	for _, v := range points {
		if v[0] > right {
			right = v[1]
			ans++
		}
	}
	return ans
}
