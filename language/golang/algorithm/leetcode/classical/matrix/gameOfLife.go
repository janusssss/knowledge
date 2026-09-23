/*
根据百度百科，生命游戏，简称为生命，是英国数学家约翰·何顿·康威在 1970 年发明的细胞自动机。
给定一个包含 m × n 个格子的面板，每一个格子都可以看成是一个细胞。
每个细胞都具有一个初始状态： 1 即为 活细胞 （live），或 0 即为 死细胞 （dead）。
每个细胞与其八个相邻位置（水平，垂直，对角线）的细胞都遵循以下四条生存定律：
 1. 如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
 2. 如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
 3. 如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
 4. 如果死细胞周围正好有三个活细胞，则该位置死细胞复活；

下一个状态是通过将上述规则同时应用于当前状态下的每个细胞所形成的，其中细胞的出生和死亡是同时发生的。
给你 m x n 网格面板 board 的当前状态，返回下一个状态。

给定当前 board 的状态，更新 board 到下一个状态。
注意 你不需要返回任何东西。
*/
package main

import (
	"fmt"
	"go-career/algorithm/leetcode"
)

type exampleGameOfLife struct {
	board  [][]int
	answer [][]int
}

func main() {
	s := []exampleGameOfLife{
		{leetcode.To2intSlice("[[1]]"), leetcode.To2intSlice("[[0]]")},
		{
			leetcode.To2intSlice("[" +
				"[0,1,0]," +
				"[0,0,1]," +
				"[1,1,1]," +
				"[0,0,0]]"),
			leetcode.To2intSlice("[[0,0,0],[1,0,1],[0,1,1],[0,1,0]]"),
		},
	}
	for _, v := range s {
		gameOfLife(v.board)
		fmt.Printf("%v\n%v\n", v.board, v.answer)
	}
}

/*
1. 初始化t,两个额外的行临时储存结果
2. 遍历计算累加值
3. t变量回填源矩阵
*/
func gameOfLife(board [][]int) {
	// 1.
	t := make([][]int, 2)
	for i, _ := range t {
		t[i] = make([]int, len(board[0]))
	}
	// 2.
	for i, r := range board {
		for j, _ := range r {
			if i > 1 {
				board[i-2][j] = t[i%2][j]
			}
			t[i%2][j] = compute(board, i, j)
		}
	}
	// 3.
	for n, i := len(board)-1, 0; n >= 0 && i < 2; i, n = i+1, n-1 {
		board[n] = t[n%2]
	}
}

// 当前位置经过生命游戏规格变换后的结果
func compute(b [][]int, r, v int) (res int) {
	r--
	v--
	c := 0
	for i := 1; i <= 9 && c < 4; i++ {
		if r >= 0 && r < len(b) && v >= 0 && v < len(b[0]) && i != 5 {
			c += b[r][v]
		}
		v++
		if i%3 == 0 {
			v = v - 3
			r++
		}
	}
	switch c {
	case 2:
		res = b[r-2][v+1]
	case 3:
		res = 1
	}
	return
}
