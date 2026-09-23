/*
力扣->面试经典面试150题->15.三数之和
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，
同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。
难度：中等
*/
package main

import (
	"fmt"
	"sort"
)

type example struct {
	nums   []int
	result [][]int
}

func main() {
	examples := []example{
		//{[]int{34, 55, 79, 28, 46, 33, 2, 48, 31, -3, 84, 71, 52, -3, 93, 15, 21, -43, 57, -6, 86, 56, 94, 74, 83, -14, 28, -66, 46, -49, 62, -11, 43, 65, 77, 12, 47, 61, 26, 1, 13, 29, 55, -82, 76, 26, 15, -29, 36, -29, 10, -70, 69, 17, 49}, [][]int{}},
		{[]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}, [][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}, [][]int{{-4, 0, 4}, {-4, 1, 3}, {-3, -1, 4}, {-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, -1, 2}, {-1, 0, 1}}},
		{[]int{-2, 0, 0, 2, 2}, [][]int{{-2, 0, 2}}},
		{[]int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{-2, 0, 1, 1, 2}, [][]int{{-2, 0, 2}, {-2, 1, 1}}},
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 1, 1}, [][]int{}},
	}
	for _, ex := range examples {
		fmt.Printf("Algorithm:\n%v\nResult:\n%v\n------\n", threeSum(ex.nums), ex.result)
	}
}

/*
1. nums排序（小到大）
2. l, r := 0, len(nums)-1，循环l->正数
3. m := l + 1，嵌套循环m->r
*/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	l, r := 0, len(nums)-1
	for nums[l] <= 0 {
		t := 0 - (nums[l] + nums[r])
		for m := l + 1; m < r; m++ {
			if nums[m] > t {
				r--
				for nums[r] == nums[r+1] && nums[r] >= 0 {
					r--
				}
				t = 0 - (nums[l] + nums[r])
				m--
				continue
			}

			if nums[m] == t {
				result = append(result, []int{nums[l], nums[m], nums[r]})
				r--
				for nums[r] == nums[r+1] && r > l && nums[r] >= 0 {
					r--
				}
				t = 0 - (nums[l] + nums[r])
			}
		}
		l++
		for nums[l] == nums[l-1] && l < r && nums[l] <= 0 {
			l++
		}
		if r <= l {
			break
		}
		r = len(nums) - 1
		t = 0 - (nums[l] + nums[r])
	}
	return result
}
