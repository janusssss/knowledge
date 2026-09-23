/*
给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用原地算法。

输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
输出：[[1,0,1],[0,0,0],[1,0,1]]
*/
package main

import (
	"fmt"
	"go-career/algorithm/leetcode"
)

type exampleSetZeroes struct {
	matrix [][]int
	answer [][]int
}

func main() {
	s := []exampleSetZeroes{
		{leetcode.To2intSlice("[[1,1,1],[1,0,1],[1,1,1]]"), leetcode.To2intSlice("[[1,0,1],[0,0,0],[1,0,1]]")},
	}
	for _, v := range s {
		setZeroes(v.matrix)
		fmt.Println(v.matrix, "\n", v.answer)
	}
}

/*
 1. 遍历矩阵matrix，把需要置零的行和列的下标值存入响应的map的key中
 2. 遍历map把响应的行和列置零
    a. 把每列的下标值存入map中
    a. 先遍历行的map，把行置零，以及该列下标删除
    c. 遍历列的map，嵌套遍历需要删除的map，逐一置零
*/
func setZeroes(matrix [][]int) {
	// rm: 需要删除行的下标值
	// vm: 需要删除列的下标值
	rm, vm := make(map[int]struct{}), make(map[int]struct{})
	// 1.
	for i, r := range matrix {
		for j, v := range r {
			if v == 0 {
				rm[i] = struct{}{}
				vm[j] = struct{}{}
			}
		}
	}
	// m: 没列需要删除的下标值
	m := make(map[int]struct{})
	for i := range len(matrix) {
		m[i] = struct{}{}
	}
	// 2.a
	for k := range rm {
		//2.b
		matrix[k] = make([]int, len(matrix[0]))
		delete(m, k)
	}
	//2.c
	for k := range vm {
		for i := range m {
			matrix[i][k] = 0
		}
	}
}
