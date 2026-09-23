package algorithm

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

var sortExamples = []struct {
	arr    []int
	answer []int
}{
	// 基础测试用例
	{[]int{3, 2, 1}, []int{1, 2, 3}},
	{[]int{5, 2, 3}, []int{2, 3, 5}},
	{[]int{1, 2, 3}, []int{1, 2, 3}},             // 已排序
	{[]int{3, 3, 1}, []int{1, 3, 3}},             // 包含重复元素
	{[]int{2, 2, 1, 3}, []int{1, 2, 2, 3}},       // 包含重复元素
	{[]int{4, 2, 4, 1, 3}, []int{1, 2, 3, 4,4}},       // 包含重复元素
	{[]int{3, 1, 4, 1, 5}, []int{1, 1, 3, 4, 5}}, // 包含重复元素

	// 边界情况
	{[]int{1}, []int{1}},       // 单元素
	{[]int{}, []int{}},         // 空数组
	{[]int{2, 1}, []int{1, 2}}, // 两个元素

	// 负数测试
	{[]int{-1, -3, 2}, []int{-3, -1, 2}},
	{[]int{0, -5, 5}, []int{-5, 0, 5}},

	// 重复元素
	{[]int{2, 2, 2}, []int{2, 2, 2}}, // 全相同
	{[]int{3, 1, 3, 1}, []int{1, 1, 3, 3}},

	// 较大数组
	{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	{[]int{5, 3, 8, 1, 9, 2}, []int{1, 2, 3, 5, 8, 9}},

	// 逆序数组
	{[]int{10, 9, 8, 7, 6}, []int{6, 7, 8, 9, 10}},

	// 包含零
	{[]int{0, 0, 1, 0}, []int{0, 0, 0, 1}},

	// 大数测试
	{[]int{100, 1, 1000, 10}, []int{1, 10, 100, 1000}},
}

func TestBubbleSort(t *testing.T) {
	standardTest(bubbleSort)
}

func TestQuickSort(t *testing.T) {
	standardTest(quickSort)
}

func standardTest(fn func(arr []int)) {
	for _, e := range sortExamples {
		answer := e.answer
		fn(e.arr)
		result := e.arr
		if !reflect.DeepEqual(answer, result) {
			fmt.Printf("answer:%v, result:%v", answer, result)
			os.Exit(1)
		}
	}
}
