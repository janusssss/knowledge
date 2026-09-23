/*
给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

示例 1：
输入：matrix = [
[1,2,3],
[4,5,6],
[7,8,9]]
输出：[
[7,4,1],
[8,5,2],
[9,6,3]]
*/
package main

import (
	"fmt"
	"go-career/algorithm/leetcode"
)

func main() {
	s := []exampleRotate{
		{leetcode.To2intSlice("[[1,2,3,4,5],[6,7,8,9,10],[11,12,13,14,15],[16,17,18,19,20],[21,22,23,24,25]]"),
			leetcode.To2intSlice("[[21,16,11,6,1],[22,17,12,7,2],[23,18,13,8,3],[24,19,14,9,4],[25,20,15,10,5]]")},
		{leetcode.To2intSlice("[[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]"),
			leetcode.To2intSlice("[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]")},
	}
	for _, o := range s {
		rotate(o.matrix)
		fmt.Println(o.matrix, "\n", o.answer)
	}
}

type exampleRotate struct {
	matrix [][]int
	answer [][]int
}

/*
 1. 从外层正方形开始旋转，逐层递进，直到2*2或1*1的正方形的最小正方形，单层旋转分3次完成
    a. 顶部x轴旋转至右侧y轴
    b. 右侧y轴旋转至底部x轴
    c. 底部x轴旋转至左侧y轴
    注意: 旋转过程中，顶部x轴成为了临时存储其他轴向数据的地方，而且首尾是对调的
*/
func rotate(matrix [][]int) {
	// 1
	for x := 0; x < len(matrix)/2; x++ {
		// h:正方形左侧和顶部边界点
		// t:正方形右侧和底部边界点
		h, t := x, len(matrix)-x-1
		// a
		for i := h; i < t; i++ {
			matrix[h][i], matrix[i][t] = matrix[i][t], matrix[h][i]
		}
		matrix[h][h], matrix[t][t] = matrix[t][t], matrix[h][h]

		// b
		for i := h + 1; i < t; i++ {
			matrix[h][i], matrix[t][t-(i-h)] = matrix[t][t-(i-h)], matrix[h][i]
		}
		matrix[h][h], matrix[t][h] = matrix[t][h], matrix[h][h]

		// c
		for i := h + 1; i < t; i++ {
			matrix[h][i], matrix[t-(i-h)][h] = matrix[t-(i-h)][h], matrix[h][i]
		}
	}
}
