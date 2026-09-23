/*
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
*/

package main

import "fmt"

type spiralOrderExample struct {
	matrix [][]int
	answer []int
}

func main() {
	es := []spiralOrderExample{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
	}
	for _, e := range es {
		fmt.Print(spiralOrder(e.matrix), "\n", e.answer, "\n--------------------------------------\n")
	}
}

/*
没有技巧，纯模拟
*/
func spiralOrder(matrix [][]int) []int {
	n := len(matrix[0]) * len(matrix)
	x, y := 0, 0
	ans := make([]int, n)
	r, d, l, u := len(matrix[0])-1, len(matrix)-1, 0, 1
	c := 0
	for i := 0; i < n; i++ {
		ans[i] = matrix[y][x]
		switch c {
		case 0:
			if x == r {
				y++
				c = 1
				r--
				continue
			}
			x++
		case 1:
			if y == d {
				x--
				c = 2
				d--
				continue
			}
			y++
		case 2:
			if x == l {
				y--
				l++
				c = 3
				continue
			}
			x--
		case 3:
			if y == u {
				x++
				c = 0
				u++
				continue
			}
			y--
		}
	}
	return ans
}
